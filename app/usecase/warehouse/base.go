// Package warehouse provides usecase implementation for warehouse
package warehouse

import (
	"context"

	"github.com/jackc/pgx/v5"

	database "github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// Status defines the warehouse status
type Status bool

const (
	// Status_Active status active
	Status_Active Status = true
	// Status_Deactive status deactive
	Status_Deactive Status = false
)

// DefaultStatus default status for warehouse
const DefaultStatus = Status_Active

// Bool returns the boolean value of the status
func (s Status) Bool() bool {
	return bool(s)
}

// Warehouse defines the warehouse model
type Warehouse struct {
	ID     int64
	Name   string
	Status Status
	ShopID int64
}

// Usecase defines the warehouse usecase interface
type Usecase interface {
	UpdateStatus(ctx context.Context, id int64, status Status) *errors.Error
	TransferProduct(ctx context.Context, arg TransferProductParams) *errors.Error
	UpdateProductStock(ctx context.Context, arg UpdateProductStockParams) *errors.Error
	ReserveProductStock(ctx context.Context, arg UpdateProductStockParams) *errors.Error
	ReleaseProductStock(ctx context.Context, arg UpdateProductStockParams) *errors.Error
}

// Repository defines repository interface for warehouse usecase
type Repository interface {
	WithTx(tx pgx.Tx) *dbgen.Queries
	UpdateWarehouseStatus(ctx context.Context, arg dbgen.UpdateWarehouseStatusParams) error
	IncreaseStock(ctx context.Context, arg dbgen.IncreaseStockParams) (int64, error)
	DecreaseStock(ctx context.Context, arg dbgen.DecreaseStockParams) (int64, error)
	UpdateProductStock(ctx context.Context, arg dbgen.UpdateProductStockParams) error
	ReserveStock(ctx context.Context, arg dbgen.ReserveStockParams) error
}

type warehouseImpl struct {
	pool database.PgxPoolIface
	repo Repository
}

var _ Usecase = (*warehouseImpl)(nil)

// NewWarehouse instantiate a new warehouse usecase
func NewWarehouse(pool database.PgxPoolIface, repo Repository) Usecase {
	return &warehouseImpl{
		pool: pool,
		repo: repo,
	}
}
