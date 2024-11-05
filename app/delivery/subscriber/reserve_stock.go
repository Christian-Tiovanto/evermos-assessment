package subscriber

import (
	"context"
	"encoding/json"
	"fmt"

	"go.uber.org/zap"

	database "github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/rabbitmq"
	"github.com/mughieams/evermos-assessment/app/common/log"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
	proto "github.com/mughieams/evermos-assessment/protobuf/api"
)

// ReserveStockSubscriber is a struct to handle release stock subscriber
type ReserveStockSubscriber struct {
	subs      rabbitmq.Subscription
	warehouse warehouse.Usecase
}

// NewReserveStockSubscriber is a constructor to create new ReserveStockSubscriber
func NewReserveStockSubscriber(subs rabbitmq.Subscription, pool database.PgxPoolIface) (*ReserveStockSubscriber, error) {
	if err := subs.QueueDeclare(ReserveStockQueue); err != nil {
		return nil, err
	}

	repo := dbgen.New(pool)
	warehouse := warehouse.NewWarehouse(pool, repo)
	return &ReserveStockSubscriber{
		subs:      subs,
		warehouse: warehouse,
	}, nil
}

// Listen is a method to listen the message
func (s *ReserveStockSubscriber) Listen() error {
	return s.subs.Consume(ReserveStockQueue, s.Execute)
}

// Execute is a method to execute the message
func (s *ReserveStockSubscriber) Execute(msg []byte) {
	logger := log.FromContext(context.Background()).With(
		zap.String("service", "ReserveStockSubscriber.Execute"),
	)

	var payload proto.UpdateProductStockRequest
	if err := json.Unmarshal(msg, &payload); err != nil {
		logger.Error("failed to unmarshal message", zap.Error(err))
		logger.Error(
			fmt.Sprintf("%s: %s", UnmarshallError, err.Error()),
			zap.String("error", err.Error()),
			zap.Stack("stacktrace"),
		)
		return
	}

	if errSvc := s.warehouse.ReserveProductStock(context.Background(), warehouse.UpdateProductStockParams{
		WarehouseID: payload.WarehouseId,
		ProductID:   payload.ProductId,
		Quantity:    payload.Quantity,
	}); errSvc != nil {
		logger.Error("failed to release stock", zap.Error(errSvc))
	}
}
