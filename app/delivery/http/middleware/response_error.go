package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

// UnknownError is a default error for unknown error
var UnknownError = status.New(codes.Unknown, "Internal server error").Err()

// ResponseError is a middleware to handle error response
func ResponseError(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	res, err := handler(ctx, req)
	if err != nil {
		var svcerr *errors.Error
		if ok := errors.As(err, &svcerr); !ok {
			return res, UnknownError
		}
		return res, status.New(svcerr.GRPCCode(), svcerr.Message()).Err()
	}
	return res, nil
}
