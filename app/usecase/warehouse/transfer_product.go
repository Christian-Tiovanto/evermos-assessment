package warehouse

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// TransferProductParams is a parameter to transfer product
type TransferProductParams struct {
	FromWarehouseID int64
	ToWarehouseID   int64
	ProductID       int64
	Quantity        int32
}

func (s *warehouseImpl) TransferProduct(ctx context.Context, arg TransferProductParams) *errors.Error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	repoWithTx := s.repo.WithTx(tx)
	rowsAffected, err := repoWithTx.DecreaseStock(ctx, dbgen.DecreaseStockParams{
		WarehouseID: arg.FromWarehouseID,
		ProductID:   arg.ProductID,
		Quantity:    arg.Quantity,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	if rowsAffected == 0 {
		return errors.NotFound("product does not have sufficient stock")
	}

	rowsAffected, err = repoWithTx.IncreaseStock(ctx, dbgen.IncreaseStockParams{
		WarehouseID: arg.ToWarehouseID,
		ProductID:   arg.ProductID,
		Quantity:    arg.Quantity,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	if rowsAffected == 0 {
		if err = repoWithTx.AddProductToStock(ctx, dbgen.AddProductToStockParams{
			WarehouseID: arg.ToWarehouseID,
			ProductID:   arg.ProductID,
			Quantity:    arg.Quantity,
		}); err != nil {
			return errors.InternalServer(pkgerrors.WithStack(err))
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
