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

// ActivateWarehouse is a method to activate warehouse
func (server *WarehouseServer) ActivateWarehouse(ctx context.Context, req *proto.WarehouseIDRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.UpdateStatus(ctx, req.WarehouseId, usecase.Status_Active); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Warehouse activated"}, nil
}

// DeactivateWarehouse is a method to deactivate warehouse
func (server *WarehouseServer) DeactivateWarehouse(ctx context.Context, req *proto.WarehouseIDRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.UpdateStatus(ctx, req.WarehouseId, usecase.Status_Deactive); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Warehouse deactivated"}, nil
}

// TransferProduct is a method to transfer product
func (server *WarehouseServer) TransferProduct(ctx context.Context, req *proto.TransferProductRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.TransferProduct(ctx, usecase.TransferProductParams{
		FromWarehouseID: req.FromWarehouseId,
		ToWarehouseID:   req.ToWarehouseId,
		ProductID:       req.ProductId,
		Quantity:        req.Quantity,
	}); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Product transferred"}, nil
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
