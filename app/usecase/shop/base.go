// Package shop provides usecase implementation for shop
package shop

import (
	"context"

	"github.com/jackc/pgx/v5"

	database "github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/product"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
)

// Usecase defines the shop usecase interface
type Usecase interface {
	CreateShop(ctx context.Context, arg CreateParams) *errors.Error
	CreateWarehouse(ctx context.Context, arg warehouse.Warehouse) *errors.Error
	GetWarehouses(ctx context.Context, shopID int64) ([]warehouse.Warehouse, *errors.Error)
	AddProduct(ctx context.Context, arg AddProductParams) *errors.Error
	GetProducts(ctx context.Context, shopID int64) ([]product.Product, *errors.Error)
}

// Repository defines repository interface for shop usecase
type Repository interface {
	WithTx(tx pgx.Tx) *dbgen.Queries
	CreateShop(ctx context.Context, arg dbgen.CreateShopParams) error
	CreateWarehouse(ctx context.Context, arg dbgen.CreateWarehouseParams) error
	GetWarehousesByShopID(ctx context.Context, shopID int64) ([]dbgen.GetWarehousesByShopIDRow, error)
	GetProductsByShopID(ctx context.Context, shopID int64) ([]dbgen.GetProductsByShopIDRow, error)
}

type shopImpl struct {
	pool database.PgxPoolIface
	repo Repository
}

var _ Usecase = (*shopImpl)(nil)

// NewShop instantiate a new shop usecase
func NewShop(repo Repository) Usecase {
	return &shopImpl{
		repo: repo,
	}
}
