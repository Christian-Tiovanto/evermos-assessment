// Package product provides usecase implementation for product
package product

import (
	"context"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// Usecase defines the product usecase interface
type Usecase interface {
	GetProducts(ctx context.Context) ([]Product, *errors.Error)
	GetProductByID(ctx context.Context, id int64) (Product, *errors.Error)
}

// Repository defines repository interface for product usecase
type Repository interface {
	GetProducts(ctx context.Context) ([]dbgen.GetProductsRow, error)
	GetProductByID(ctx context.Context, id int64) (dbgen.GetProductByIDRow, error)
}

type productImpl struct {
	repo Repository
}

var _ Usecase = (*productImpl)(nil)

// NewProduct instantiate a new product usecase
func NewProduct(repo Repository) Usecase {
	return &productImpl{
		repo: repo,
	}
}
