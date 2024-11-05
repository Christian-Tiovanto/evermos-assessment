package product

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	"github.com/mughieams/evermos-assessment/app/common/util"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// Product defines the product usecase interface
type Product struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Quantity    int32
	ShopID      int64
	WarehouseID int64
}

// GetProducts is a method to get products
func (p *productImpl) GetProducts(ctx context.Context) ([]Product, *errors.Error) {
	products, err := p.repo.GetProducts(ctx)
	if err != nil {
		return nil, errors.InternalServer(pkgerrors.WithStack(err))
	}

	return util.SlimMap(products, func(data dbgen.GetProductsRow) Product {
		return Product{
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
