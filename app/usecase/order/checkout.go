package order

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/actor"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	"github.com/mughieams/evermos-assessment/app/delivery/subscriber"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

func (s *orderImpl) CheckoutOrder(ctx context.Context, arg Order) *errors.Error {
	actor := actor.FromContext(ctx)
	if actor.IsEmpty() {
		return errors.Forbidden("actor is empty")
	}

	product, errSvc := s.product.GetProductByID(ctx, arg.ProductID)
	if errSvc != nil {
		return errSvc
	}
	if product.Quantity < arg.Quantity {
		return errors.BadRequest("stock is not enough")
	}

	arg.TotalPrice = product.Price * float64(arg.Quantity)
	var totalPrice pgtype.Numeric
	if err := totalPrice.Scan(fmt.Sprintf("%f", arg.TotalPrice)); err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	var err error
	arg.ID, err = s.repo.CreateOrder(ctx, dbgen.CreateOrderParams{
		UserID:      actor.ID,
		ShopID:      arg.ShopID,
		WarehouseID: arg.WarehouseID,
		ProductID:   arg.ProductID,
		Quantity:    arg.Quantity,
		TotalPrice:  totalPrice,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	errSubs := s.reserveOrder(arg)
	if errSubs != nil {
		return errors.InternalServer(pkgerrors.WithStack(errSubs))
	}

	go s.waitingForPayment(ctx, arg)

	return nil
}

// DefaultWaitingTime default waiting time before release stock
const DefaultWaitingTime = 10 * time.Minute

func (s *orderImpl) waitingForPayment(ctx context.Context, order Order) {
	select {
	case <-ctx.Done():
		return
	case <-time.After(DefaultWaitingTime):
		s.releaseOrder(order)
	}
}

func (s *orderImpl) releaseOrder(order Order) {
	payload, _ := json.Marshal(proto.UpdateProductStockRequest{
		WarehouseId: order.WarehouseID,
		ProductId:   order.ProductID,
		Quantity:    order.Quantity,
	})
	_ = s.subs.Publish(subscriber.ReleaseStockQueue, payload)
	_ = s.UpdateStatus(context.Background(), order.ID, Status_Timeout)
}

func (s *orderImpl) reserveOrder(order Order) *errors.Error {
	payload, err := json.Marshal(proto.UpdateProductStockRequest{
		WarehouseId: order.WarehouseID,
		ProductId:   order.ProductID,
		Quantity:    order.Quantity,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}
	_ = s.subs.Publish(subscriber.ReserveStockQueue, payload)

	return nil
}
