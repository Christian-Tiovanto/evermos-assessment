package warehouse

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

func (s *warehouseImpl) ReleaseProductStock(ctx context.Context, arg UpdateProductStockParams) *errors.Error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	repoWithTx := s.repo.WithTx(tx)
	rowsAffected, err := repoWithTx.IncreaseStock(ctx, dbgen.IncreaseStockParams{
		Quantity:    arg.Quantity,
		ProductID:   arg.ProductID,
		WarehouseID: arg.WarehouseID,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	if rowsAffected == 0 {
		return errors.BadRequest("product not found")
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
