package handler

import (
	"context"

	usecase "github.com/mughieams/evermos-assessment/app/usecase/warehouse"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

// WarehouseServer is the config method caller
type WarehouseServer struct {
	proto.UnimplementedWarehouseServiceServer
	usecase usecase.Usecase
}

var _ proto.WarehouseServiceServer = (*WarehouseServer)(nil)

// NewWarehouseServer is handler to wrap Server
func NewWarehouseServer(usecase usecase.Usecase) *WarehouseServer {
	return &WarehouseServer{usecase: usecase}
}

// UpdateProductStock is a method to update product stock
func (server *WarehouseServer) UpdateProductStock(ctx context.Context, req *proto.UpdateProductStockRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.UpdateProductStock(ctx, usecase.UpdateProductStockParams{
		WarehouseID: req.WarehouseId,
		ProductID:   req.ProductId,
		Quantity:    req.Quantity,
	}); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Product stock updated"}, nil
}
