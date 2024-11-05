package warehouse

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

func (s *warehouseImpl) UpdateStatus(ctx context.Context, id int64, status Status) *errors.Error {
	err := s.repo.UpdateWarehouseStatus(ctx, dbgen.UpdateWarehouseStatusParams{
		ID:     id,
		Status: pgtype.Bool{Bool: status.Bool(), Valid: true},
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
