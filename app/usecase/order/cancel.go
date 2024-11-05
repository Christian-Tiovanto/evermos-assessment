package order

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

func (s *orderImpl) CancelOrder(ctx context.Context, id int64) *errors.Error {
	order, err := s.repo.UpdateOrderStatus(ctx, dbgen.UpdateOrderStatusParams{
		ID:     id,
		Status: Status_Cancelled.String(),
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	if order.Status != Status_Cancelled.String() {
		return errors.BadRequest("order already PAID or CANCELLED")
	}

	go func() {
		s.releaseOrder(Order{
			ProductID:   order.ProductID,
			WarehouseID: order.WarehouseID,
			Quantity:    order.Quantity,
		})
	}()

	return nil
}
