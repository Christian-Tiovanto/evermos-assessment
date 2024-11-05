package subscriber

import (
	"context"
	"encoding/json"

	database "github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/dependencies/rabbitmq"
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
	var payload proto.UpdateProductStockRequest
	if err := json.Unmarshal(msg, &payload); err != nil {
		return
	}

	_ = s.warehouse.ReserveProductStock(context.Background(), warehouse.UpdateProductStockParams{
		WarehouseID: payload.WarehouseId,
		ProductID:   payload.ProductId,
		Quantity:    payload.Quantity,
	})
}
