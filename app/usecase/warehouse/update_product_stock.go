package warehouse

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// UpdateProductStockParams is a parameter to update product stock
type UpdateProductStockParams struct {
	WarehouseID int64
	ProductID   int64
	Quantity    int32
}

func (s *warehouseImpl) UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) *errors.Error {
	err := s.repo.UpdateProductStock(ctx, dbgen.UpdateProductStockParams{
		WarehouseID: arg.WarehouseID,
		ProductID:   arg.ProductID,
		Quantity:    arg.Quantity,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
