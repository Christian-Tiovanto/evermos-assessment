package handler

import (
	"context"

	usecase "github.com/mughieams/evermos-assessment/app/usecase/user"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

// UserServer is the config method caller
type UserServer struct {
	proto.UnimplementedUserServiceServer
	usecase usecase.Usecase
}

var _ proto.UserServiceServer = (*UserServer)(nil)

// NewUserServer is handler to wrap Server
func NewUserServer(usecase usecase.Usecase) *UserServer {
	return &UserServer{usecase: usecase}
}

// Register is a method to register user
func (server *UserServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.Register(ctx, usecase.RegisterParams{
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	}); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "User registered"}, nil
}

// Login is a method to login user
func (server *UserServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	token, err := server.usecase.Login(ctx, usecase.LoginParams{
		Email:    req.Email,
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &proto.LoginResponse{Token: token}, nil
}
