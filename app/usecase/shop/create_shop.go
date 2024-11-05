package shop

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// CreateParams defines the shop create params
type CreateParams struct {
	Name    string
	Address string
}

// CreateShop is a method to create shop
func (s *shopImpl) CreateShop(ctx context.Context, arg CreateParams) *errors.Error {
	err := s.repo.CreateShop(ctx, dbgen.CreateShopParams{
		Name:    arg.Name,
		Address: arg.Address,
	})
	if err != nil {
		return errors.InternalServer(pkgerrors.WithStack(err))
	}

	return nil
}
