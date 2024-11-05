// Package order provides usecase implementation for order
package order

import (
	"context"

	"github.com/mughieams/evermos-assessment/app/common/dependencies/rabbitmq"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/product"
)

// Status defines the order status
type Status string

const (
	// Status_Pending status pending
	Status_Pending Status = "PENDING"
	// Status_Timeout status timeout
	Status_Timeout Status = "TIMEOUT"
	// Status_Paid status paid
	Status_Paid Status = "PAID"
	// Status_Cancelled status cancelled
	Status_Cancelled Status = "CANCELLED"
)

// DefaultStatus default status for order
const DefaultStatus = Status_Pending

func (s Status) String() string {
	return string(s)
}

// Order defines the order model
type Order struct {
	ID          int64
	UserID      int64
	ShopID      int64
	Status      Status
	ProductID   int64
	WarehouseID int64
	Quantity    int32
	TotalPrice  float64
}

// Usecase defines the order usecase interface
type Usecase interface {
	CheckoutOrder(ctx context.Context, arg Order) *errors.Error
	CancelOrder(ctx context.Context, id int64) *errors.Error
	UpdateStatus(ctx context.Context, id int64, status Status) *errors.Error
}

// Repository defines repository interface for order usecase
type Repository interface {
	CreateOrder(ctx context.Context, arg dbgen.CreateOrderParams) (int64, error)
	UpdateOrderStatus(ctx context.Context, arg dbgen.UpdateOrderStatusParams) (dbgen.UpdateOrderStatusRow, error)
}

type orderImpl struct {
	repo    Repository
	product product.Usecase
	subs    rabbitmq.Subscription
}

var _ Usecase = (*orderImpl)(nil)

// NewOrder instantiate a new order usecase
func NewOrder(repo Repository, product product.Usecase, subs rabbitmq.Subscription) Usecase {
	return &orderImpl{
		repo:    repo,
		product: product,
		subs:    subs,
	}
}
