package shop

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	"github.com/mughieams/evermos-assessment/app/common/util"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
)

// GetWarehouses is a method to get all warehouses of a shop
func (s *shopImpl) GetWarehouses(ctx context.Context, shopID int64) ([]warehouse.Warehouse, *errors.Error) {
	warehouses, err := s.repo.GetWarehousesByShopID(ctx, shopID)
	if err != nil {
		return nil, errors.InternalServer(pkgerrors.WithStack(err))
	}

	return util.SlimMap(warehouses, func(data dbgen.GetWarehousesByShopIDRow) warehouse.Warehouse {
		return warehouse.Warehouse{
			ID:     data.ID,
			Name:   data.Name,
			Status: warehouse.Status(data.Status.Bool),
			ShopID: data.ShopID,
		}
	}), nil
}
