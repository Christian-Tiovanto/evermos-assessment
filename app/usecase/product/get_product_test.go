package product_test

import (
	stderrors "errors"

	"github.com/jackc/pgx/v5"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/product"
)

func (s *ProductTestSuite) Test_GetProductByID() {
	s.Run("fail get product by id", func() {
		s.repo.EXPECT().GetProductByID(s.ctx, int64(1)).Return(dbgen.GetProductByIDRow{}, stderrors.New("database down"))

		res, err := s.svc.GetProductByID(s.ctx, int64(1))
		s.Equal(errors.InternalServerError, err.Type())
		s.Empty(product.Product{}, res)
	})

	s.Run("fail get product by id not found", func() {
		s.repo.EXPECT().GetProductByID(s.ctx, int64(1)).Return(dbgen.GetProductByIDRow{}, pgx.ErrNoRows)

		res, err := s.svc.GetProductByID(s.ctx, int64(1))
		s.Equal(errors.BadRequestError, err.Type())
		s.Empty(product.Product{}, res)
	})

	s.Run("success get product by id", func() {
		productData := dbgen.GetProductByIDRow{
			ID:          1,
			Name:        "Product 1",
			Description: "Description 1",
			Price:       10000,
		}
		s.repo.EXPECT().GetProductByID(s.ctx, int64(1)).Return(productData, nil)

		res, err := s.svc.GetProductByID(s.ctx, int64(1))
		s.Nil(err)
		s.Equal(product.Product{
			ID:          1,
			Name:        "Product 1",
			Description: "Description 1",
			Price:       10000,
		}, res)
	})
}
