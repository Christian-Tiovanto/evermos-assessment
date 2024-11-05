package handler

import (
	"context"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	usecase "github.com/mughieams/evermos-assessment/app/usecase/order"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

// OrderServer is the config method caller
type OrderServer struct {
	proto.UnimplementedOrderServiceServer
	usecase usecase.Usecase
}

var _ proto.OrderServiceServer = (*OrderServer)(nil)

// NewOrderServer is handler to wrap Server
func NewOrderServer(usecase usecase.Usecase) *OrderServer {
	return &OrderServer{usecase: usecase}
}

// CheckoutOrder is a method to checkout order
func (server *OrderServer) CheckoutOrder(ctx context.Context, req *proto.CheckoutOrderRequest) (*proto.MessageResponse, error) {
	if req.Quantity <= 0 {
		return nil, errors.BadRequest("quantity must be greater than 0")
	}

	if err := server.usecase.CheckoutOrder(ctx, usecase.Order{
		ShopID:      req.ShopId,
		WarehouseID: req.WarehouseId,
		ProductID:   req.ProductId,
		Quantity:    req.Quantity,
	}); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Order checked out"}, nil
}

// ConfirmOrder is a method to confirm order
func (server *OrderServer) ConfirmOrder(ctx context.Context, req *proto.OrderIDRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.UpdateStatus(ctx, req.OrderId, usecase.Status_Paid); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Order confirmed"}, nil
}

// CancelOrder is a method to cancel order
func (server *OrderServer) CancelOrder(ctx context.Context, req *proto.OrderIDRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.CancelOrder(ctx, req.OrderId); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Order cancelled"}, nil
}
