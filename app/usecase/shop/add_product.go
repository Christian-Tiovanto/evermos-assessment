package shop

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// AddProductParams defines the parameters for AddProduct method
type AddProductParams struct {
	ShopID      int64
	WarehouseID int64
	Name        string
	Description string
	Price       float32
	Quantity    int32
}

// AddProduct is a method to add product
func (s *shopImpl) AddProduct(ctx context.Context, arg AddProductParams) *errors.Error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	repoWithTx := s.repo.WithTx(tx)

	var price pgtype.Numeric
	if err = price.Scan(arg.Price); err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	productID, err := repoWithTx.AddProduct(ctx, dbgen.AddProductParams{
		Name:        arg.Name,
		Description: arg.Description,
		Price:       price,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	if err := repoWithTx.AddProductToStock(ctx, dbgen.AddProductToStockParams{
		WarehouseID: arg.WarehouseID,
		ProductID:   productID,
		Quantity:    arg.Quantity,
	}); err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
