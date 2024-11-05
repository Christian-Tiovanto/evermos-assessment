package product

import (
	"context"

	"github.com/jackc/pgx/v5"
	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
)

// GetProducts is a method to get products
func (p *productImpl) GetProductByID(ctx context.Context, id int64) (Product, *errors.Error) {
	data, err := p.repo.GetProductByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return Product{}, errors.BadRequest("product not found")
		}
		return Product{}, errors.InternalServer(pkgerrors.WithStack(err))
	}

	return Product{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Price:       data.Price,
	}, nil
}
