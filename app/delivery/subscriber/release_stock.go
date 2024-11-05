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

// ReleaseStockSubscriber is a struct to handle release stock subscriber
type ReleaseStockSubscriber struct {
	subs      rabbitmq.Subscription
	warehouse warehouse.Usecase
}

// NewReleaseStockSubscriber is a constructor to create new ReleaseStockSubscriber
func NewReleaseStockSubscriber(subs rabbitmq.Subscription, pool database.PgxPoolIface) (*ReleaseStockSubscriber, error) {
	if err := subs.QueueDeclare(ReleaseStockQueue); err != nil {
		return nil, err
	}

	repo := dbgen.New(pool)
	warehouse := warehouse.NewWarehouse(pool, repo)
	return &ReleaseStockSubscriber{
		subs:      subs,
		warehouse: warehouse,
	}, nil
}

// Listen is a method to listen the message
func (s *ReleaseStockSubscriber) Listen() error {
	return s.subs.Consume(ReleaseStockQueue, s.Execute)
}

// Execute is a method to execute the message
func (s *ReleaseStockSubscriber) Execute(msg []byte) {
	var payload proto.UpdateProductStockRequest
	if err := json.Unmarshal(msg, &payload); err != nil {
		return
	}

	_ = s.warehouse.ReleaseProductStock(context.Background(), warehouse.UpdateProductStockParams{
		WarehouseID: payload.WarehouseId,
		ProductID:   payload.ProductId,
		Quantity:    payload.Quantity,
	})
}
