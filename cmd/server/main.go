// Package main initialize dependencies and start HTTP server
package main

import (
	"context"
	"fmt"
	"os"
	"time"
	_ "time/tzdata"

	pkgerrors "github.com/pkg/errors"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"

	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/rabbitmq"
	"github.com/mughieams/evermos-assessment/app/common/log"
	delivery "github.com/mughieams/evermos-assessment/app/delivery/http"
	"github.com/mughieams/evermos-assessment/app/delivery/subscriber"
)

var logger = log.GetLogger()

func main() {
	if err := run(); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
}

func run() error {
	ctx := context.Background()

	logger.Info("Loading config...")
	cfg, err := config.Load(".env")
	if err != nil {
		return pkgerrors.Wrap(err, "load config error")
	}

	logger.Info("Connecting to postgresql...")
	dbpool, err := postgresql.NewPool(cfg.PostgreSQL)
	if err != nil {
		return pkgerrors.Wrap(err, "connect to postgresql error")
	}

	logger.Info("Connecting to rabbitmq...")
	subs, err := rabbitmq.NewSubscription(&cfg.RabbitMQ)
	if err != nil {
		return pkgerrors.Wrap(err, "connect to rabbitmq error")
	}
	defer func() {
		_ = subs.Close()
	}()

	logger.Info("Starting subscription...")
	releaseStock, err := subscriber.NewReleaseStockSubscriber(subs, dbpool)
	if err != nil {
		return pkgerrors.Wrap(err, "initialize release stock subscriber error")
	}
	go runSubscriber(releaseStock)
	reserveStock, err := subscriber.NewReserveStockSubscriber(subs, dbpool)
	if err != nil {
		return pkgerrors.Wrap(err, "initialize reserve stock subscriber error")
	}
	go runSubscriber(reserveStock)

	logger.Info("Starting http server...")
	err = startServer(ctx, cfg, dbpool, subs)
	if err != nil {
		return pkgerrors.Wrap(err, "initialize server error")
	}

	return nil
}

const (
	bootGrpcName = "commerce-server"

	serverShutdownTimeout = 60 * time.Second
	gracefulDelayTime     = 2 * time.Second
)

func startServer(ctx context.Context, cfg config.Config, dbpool postgresql.PgxPoolIface, subs rabbitmq.Subscription) error {
	boot, err := os.ReadFile(cfg.Server.BootPath)
	if err != nil {
		return pkgerrors.Wrap(err, "open boot file error")
	}

	// Bootstrap basic entries from boot config.
	rkentry.BootstrapBuiltInEntryFromYAML(boot)
	rkentry.BootstrapPluginEntryFromYAML(boot)
	// rkentry.LoggerEntryStdout.Logger = logger.Logger

	grpcEntry := rkgrpc.RegisterGrpcEntryYAML(boot)[bootGrpcName].(*rkgrpc.GrpcEntry)

	if err := delivery.RegisterGrpcHandler(grpcEntry, cfg, dbpool, subs); err != nil {
		return pkgerrors.Wrap(err, "register grpc handler error")
	}

	// Bootstrap grpc entry
	grpcEntry.Bootstrap(ctx)

	if err := delivery.RegisterRestHandler(grpcEntry.GwMux, cfg, dbpool); err != nil {
		return pkgerrors.Wrap(err, "register rest handler error")
	}

	gracefulShutdown(grpcEntry)
	return nil
}

func gracefulShutdown(grpcEntry *rkgrpc.GrpcEntry) {
	rkentry.GlobalAppCtx.WaitForShutdownSig()
	logger.Info("Termination signal received, waiting for graceful shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), serverShutdownTimeout)
	defer cancel()

	grpcEntry.Interrupt(ctx)

	time.Sleep(gracefulDelayTime)
	logger.Info("Server gracefully stopped")
}

type subs interface {
	Listen() error
}

func runSubscriber(s subs) {
	err := s.Listen()
	if err != nil {
		// trigger panic so that the subscriber pods will restart
		panic(err)
	}
}
