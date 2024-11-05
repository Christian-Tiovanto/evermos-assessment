package shop

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	"github.com/mughieams/evermos-assessment/app/common/util"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/product"
)

// GetProducts is a method to get all products of a shop
func (s *shopImpl) GetProducts(ctx context.Context, shopID int64) ([]product.Product, *errors.Error) {
	products, err := s.repo.GetProductsByShopID(ctx, shopID)
	if err != nil {
		return nil, errors.InternalServer(pkgerrors.WithStack(err))
	}

	return util.SlimMap(products, func(data dbgen.GetProductsByShopIDRow) product.Product {
		return product.Product{
			ID:          data.ID,
			Name:        data.Name,
			Description: data.Description,
			Price:       data.Price,
			Quantity:    data.Quantity,
			ShopID:      data.ShopID,
			WarehouseID: data.WarehouseID,
		}
	}), nil
}
