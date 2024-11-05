package order

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

func (s *orderImpl) UpdateStatus(ctx context.Context, id int64, status Status) *errors.Error {
	order, err := s.repo.UpdateOrderStatus(ctx, dbgen.UpdateOrderStatusParams{
		ID:     id,
		Status: status.String(),
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	if order.Status != status.String() {
		return errors.BadRequest("order already PAID or CANCELLED")
	}

	return nil
}
