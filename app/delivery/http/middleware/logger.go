package middleware

import (
	"context"
	stderrors "errors"
	"fmt"
	"net/http"
	"time"

	rkgrpcmid "github.com/rookie-ninja/rk-grpc/v2/middleware"
	rkgrpcctx "github.com/rookie-ninja/rk-grpc/v2/middleware/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	"github.com/mughieams/evermos-assessment/app/common/log"
)

const (
	incomingResponseTag = "incoming_response"
)

// Logger is a middleware to log incoming request and response
func Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	reqInfo := buildRequestLog(ctx, info)
	ctx = log.AppendContext(ctx, zap.Object("requestInfo", reqInfo))

	start := time.Now()
	res, err := handler(ctx, req)
	stop := time.Now()

	d := stop.Sub(start)
	logger := initLogger(ctx).With(
		zap.String("tag", incomingResponseTag),
		zap.Object("requestInfo", reqInfo),
		zap.Int64("latency", d.Nanoseconds()),
		zap.String("latencyHuman", d.String()),
	)

	if err != nil {
		return res, logIncomingResponseError(logger, err)
	}

	logIncomingResponseOK(logger)
	return res, nil
}

func initLogger(ctx context.Context) log.Logger {
	l := log.NewLogger()
	l.Logger = l.Logger.WithOptions(zap.AddStacktrace(zap.PanicLevel))
	return l.WithContext(ctx)
}

func logIncomingResponseOK(logger log.Logger) {
	logger.Info("OK", zap.Int("status", http.StatusOK))
}

func logIncomingResponseError(logger log.Logger, err error) error {
	var svcerr *errors.Error
	if ok := stderrors.As(err, &svcerr); !ok {
		logger.With(zap.String("stacktrace", fmt.Sprintf("%+v", err))).Error(fmt.Sprintf("Unknown error: %s", err.Error()))
		return err
	}
	logger = logger.With(zap.Object("error", errLog{svcerr}))

	loggerFn := logger.Warning
	if svcerr.HTTPStatus() == http.StatusInternalServerError {
		loggerFn = logger.Error
	}
	loggerFn(svcerr.Message(), zap.Int("status", svcerr.HTTPStatus()))
	return err
}

func buildRequestLog(ctx context.Context, info *grpc.UnaryServerInfo) requestInfo {
	remoteIP, remotePort, _ := rkgrpcmid.GetRemoteAddressSet(ctx)
	grpcService, grpcMethod := rkgrpcmid.GetGrpcInfo(info.FullMethod)
	gwMethod, gwPath, gwScheme, gwUserAgent := rkgrpcmid.GetGwInfo(rkgrpcctx.GetIncomingHeaders(ctx))
	return requestInfo{
		remoteAddr:  remoteIP + ":" + remotePort,
		grpcMethod:  grpcMethod,
		grpcService: grpcService,
		gwMethod:    gwMethod,
		gwPath:      gwPath,
		gwScheme:    gwScheme,
		gwUserAgent: gwUserAgent,
	}
}

type requestInfo struct {
	remoteAddr  string
	grpcMethod  string
	grpcService string
	gwMethod    string
	gwPath      string
	gwScheme    string
	gwUserAgent string
}

func (h requestInfo) MarshalLogObject(obj zapcore.ObjectEncoder) error {
	obj.AddString("remoteAddr", h.remoteAddr)
	obj.AddString("grpcMethod", h.grpcMethod)
	obj.AddString("grpcService", h.grpcService)
	obj.AddString("gwMethod", h.gwMethod)
	obj.AddString("gwPath", h.gwPath)
	obj.AddString("gwScheme", h.gwScheme)
	obj.AddString("gwUserAgent", h.gwUserAgent)
	return nil
}

type errLog struct {
	*errors.Error
}

func (e errLog) MarshalLogObject(obj zapcore.ObjectEncoder) error {
	obj.AddString("code", e.Code())
	obj.AddString("message", e.Message())
	obj.AddInt("status", e.HTTPStatus())
	obj.AddString("grpcCode", e.GRPCCode().String())
	obj.AddString("detail", e.String())
	return nil
}
