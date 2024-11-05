package warehouse

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

func (s *warehouseImpl) ReserveProductStock(ctx context.Context, arg UpdateProductStockParams) *errors.Error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	repoWithTx := s.repo.WithTx(tx)
	err = repoWithTx.ReserveStock(ctx, dbgen.ReserveStockParams{
		ProductID:   arg.ProductID,
		WarehouseID: arg.WarehouseID,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	rowsAffected, err := repoWithTx.DecreaseStock(ctx, dbgen.DecreaseStockParams{
		Quantity:    arg.Quantity,
		ProductID:   arg.ProductID,
		WarehouseID: arg.WarehouseID,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	if rowsAffected == 0 {
		return errors.NotFound("product does not have sufficient stock")
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
