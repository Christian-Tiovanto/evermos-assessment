// Package http provides HTTP server
package http

import (
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"

	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/rabbitmq"
	"github.com/mughieams/evermos-assessment/app/delivery/http/handler"
	"github.com/mughieams/evermos-assessment/app/delivery/http/middleware"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/order"
	product "github.com/mughieams/evermos-assessment/app/usecase/product"
	"github.com/mughieams/evermos-assessment/app/usecase/shop"
	"github.com/mughieams/evermos-assessment/app/usecase/user"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

// RegisterGrpcHandler register gRPC interceptors and handlers
func RegisterGrpcHandler(grpcEntry *rkgrpc.GrpcEntry, cfg config.Config, dbpool postgresql.PgxPoolIface, subs rabbitmq.Subscription) error {
	// order is matter, place thoughtfully
	grpcEntry.AddUnaryInterceptors(
		middleware.ResponseError,
		middleware.Auth(cfg.JWT),
	)

	registerUser(cfg, grpcEntry, dbpool)
	registerShop(cfg, grpcEntry, dbpool)
	registerProduct(cfg, grpcEntry, dbpool)
	registerWarehouse(cfg, grpcEntry, dbpool)
	registerOrder(cfg, grpcEntry, dbpool, subs)
	return nil
}

func registerUser(cfg config.Config, grpcEntry *rkgrpc.GrpcEntry, dbpool postgresql.PgxPoolIface) {
	repo := dbgen.New(dbpool)
	svc := user.NewUser(cfg.JWT, repo)
	grpcEntry.AddRegFuncGrpc(func(server *grpc.Server) {
		proto.RegisterUserServiceServer(server, handler.NewUserServer(svc))
	})
	grpcEntry.AddRegFuncGw(proto.RegisterUserServiceHandlerFromEndpoint)
}

func registerShop(_ config.Config, grpcEntry *rkgrpc.GrpcEntry, dbpool postgresql.PgxPoolIface) {
	repo := dbgen.New(dbpool)
	svc := shop.NewShop(repo)
	grpcEntry.AddRegFuncGrpc(func(server *grpc.Server) {
		proto.RegisterShopServiceServer(server, handler.NewShopServer(svc))
	})
	grpcEntry.AddRegFuncGw(proto.RegisterShopServiceHandlerFromEndpoint)
}

func registerProduct(_ config.Config, grpcEntry *rkgrpc.GrpcEntry, dbpool postgresql.PgxPoolIface) {
	repo := dbgen.New(dbpool)
	svc := product.NewProduct(repo)
	grpcEntry.AddRegFuncGrpc(func(server *grpc.Server) {
		proto.RegisterProductServiceServer(server, handler.NewProductServer(svc))
	})
	grpcEntry.AddRegFuncGw(proto.RegisterProductServiceHandlerFromEndpoint)
}

func registerWarehouse(_ config.Config, grpcEntry *rkgrpc.GrpcEntry, dbpool postgresql.PgxPoolIface) {
	repo := dbgen.New(dbpool)
	svc := warehouse.NewWarehouse(dbpool, repo)
	grpcEntry.AddRegFuncGrpc(func(server *grpc.Server) {
		proto.RegisterWarehouseServiceServer(server, handler.NewWarehouseServer(svc))
	})
	grpcEntry.AddRegFuncGw(proto.RegisterWarehouseServiceHandlerFromEndpoint)
}

func registerOrder(_ config.Config, grpcEntry *rkgrpc.GrpcEntry, dbpool postgresql.PgxPoolIface, subs rabbitmq.Subscription) {
	repo := dbgen.New(dbpool)
	productSvc := product.NewProduct(repo)
	svc := order.NewOrder(repo, productSvc, subs)
	grpcEntry.AddRegFuncGrpc(func(server *grpc.Server) {
		proto.RegisterOrderServiceServer(server, handler.NewOrderServer(svc))
	})
	grpcEntry.AddRegFuncGw(proto.RegisterOrderServiceHandlerFromEndpoint)
}
