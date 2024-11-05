// Package middleware provides gRPC and Rest API middlewares for our service.
package middleware

import (
	"context"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/mughieams/evermos-assessment/app/common/actor"
	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/common/errors"
)

var skipTokenAuthMethods = map[string]struct{}{
	"/mughieams.evermosassessment.v1.UserService/Register": {},
	"/mughieams.evermosassessment.v1.UserService/Login":    {},
}

// Auth is a middleware to authenticates user by JWT token
func Auth(cfg config.JWT) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// check which routes need to skip authorization
		_, existsInSkipTokenAuth := skipTokenAuthMethods[info.FullMethod]
		if existsInSkipTokenAuth {
			return handler(ctx, req)
		}

		user, err := authenticate(ctx, cfg)
		if err != nil {
			return nil, err
		}
		ctx = actor.SetToContext(ctx, user)
		return handler(ctx, req)
	}
}

func authenticate(ctx context.Context, cfg config.JWT) (actor.Actor, *errors.Error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return actor.Actor{}, errors.Authentication("invalid request error")
	}

	authHeader := md.Get("authorization")
	if len(authHeader) == 0 {
		return actor.Actor{}, errors.Authentication("invalid request error")
	}

	tokens := strings.Split(authHeader[0], " ")
	if len(tokens) != 2 || tokens[0] != "Bearer" {
		return actor.Actor{}, errors.Authentication("invalid request error")
	}

	claims := jwt.MapClaims{}
	token, errToken := jwt.ParseWithClaims(tokens[1], claims, func(_ *jwt.Token) (interface{}, error) {
		return []byte(cfg.Secret), nil
	})
	if errToken != nil || !token.Valid {
		return actor.Actor{}, errors.Authentication("invalid token error")
	}

	return actor.Actor{
		ID:    int64(claims["id"].(float64)),
		Email: claims["email"].(string),
		Phone: claims["phone"].(string),
	}, nil
}
