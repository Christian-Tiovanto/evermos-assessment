package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/mughieams/evermos-assessment/app/common/util"
	usecase "github.com/mughieams/evermos-assessment/app/usecase/product"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

// ProductServer is the config method caller
type ProductServer struct {
	proto.UnimplementedProductServiceServer
	usecase usecase.Usecase
}

var _ proto.ProductServiceServer = (*ProductServer)(nil)

// NewProductServer is handler to wrap Server
func NewProductServer(usecase usecase.Usecase) *ProductServer {
	return &ProductServer{usecase: usecase}
}

// GetProducts is a method to get all products
func (server *ProductServer) GetProducts(ctx context.Context, _ *emptypb.Empty) (*proto.GetProductsResponse, error) {
	products, err := server.usecase.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	return &proto.GetProductsResponse{
		Products: util.SlimMap(products, func(data usecase.Product) *proto.Product {
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
