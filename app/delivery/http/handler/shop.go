package handler

import (
	"context"

	"github.com/mughieams/evermos-assessment/app/common/util"
	"github.com/mughieams/evermos-assessment/app/usecase/product"
	usecase "github.com/mughieams/evermos-assessment/app/usecase/shop"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

// ShopServer is the config method caller
type ShopServer struct {
	proto.UnimplementedShopServiceServer
	usecase usecase.Usecase
}

var _ proto.ShopServiceServer = (*ShopServer)(nil)

// NewShopServer is handler to wrap Server
func NewShopServer(usecase usecase.Usecase) *ShopServer {
	return &ShopServer{usecase: usecase}
}

// CreateShop is a method to create shop
func (server *ShopServer) CreateShop(ctx context.Context, req *proto.CreateShopRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.CreateShop(ctx, usecase.CreateParams{
		Name:    req.Name,
		Address: req.Address,
	}); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Shop created"}, nil
}

// CreateWarehouse is a method to create a new warehouse for a shop
func (server *ShopServer) CreateWarehouse(ctx context.Context, req *proto.CreateWarehouseRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.CreateWarehouse(ctx, warehouse.Warehouse{
		ShopID: req.ShopId,
		Name:   req.Name,
	}); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Warehouse created"}, nil
}

// GetWarehouses is a method to get all warehouses in shop
func (server *ShopServer) GetWarehouses(ctx context.Context, req *proto.ShopIDRequest) (*proto.GetWarehousesResponse, error) {
	warehouses, err := server.usecase.GetWarehouses(ctx, req.ShopId)
	if err != nil {
		return nil, err
	}

	return &proto.GetWarehousesResponse{
		Warehouses: util.SlimMap(warehouses, func(data warehouse.Warehouse) *proto.Warehouse {
			return &proto.Warehouse{
				Id:     data.ID,
				Name:   data.Name,
				Status: data.Status.Bool(),
				ShopId: data.ShopID,
			}
		}),
	}, nil
}

// AddProduct is a method to add product to shop
func (server *ShopServer) AddProduct(ctx context.Context, req *proto.AddProductRequest) (*proto.MessageResponse, error) {
	if err := server.usecase.AddProduct(ctx, usecase.AddProductParams{
		ShopID:      req.ShopId,
		WarehouseID: req.WarehouseId,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Quantity:    req.Quantity,
	}); err != nil {
		return nil, err
	}

	return &proto.MessageResponse{Message: "Product added to shop"}, nil
}

// GetProducts is a method to get all products in shop
func (server *ShopServer) GetProducts(ctx context.Context, req *proto.ShopIDRequest) (*proto.GetProductsResponse, error) {
	products, err := server.usecase.GetProducts(ctx, req.ShopId)
	if err != nil {
		return nil, err
	}

	return &proto.GetProductsResponse{
		Products: util.SlimMap(products, func(data product.Product) *proto.Product {
			return &proto.Product{
				Id:          data.ID,
				Name:        data.Name,
				Description: data.Description,
				Price:       float32(data.Price),
				Quantity:    data.Quantity,
				ShopId:      data.ShopID,
				WarehouseId: data.WarehouseID,
			}
		}),
	}, nil
}
