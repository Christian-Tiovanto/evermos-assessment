package shop

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
)

// CreateWarehouse is a method to create a new warehouse for a shop
func (s *shopImpl) CreateWarehouse(ctx context.Context, arg warehouse.Warehouse) *errors.Error {
	err := s.repo.CreateWarehouse(ctx, dbgen.CreateWarehouseParams{
		ShopID: arg.ShopID,
		Name:   arg.Name,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
