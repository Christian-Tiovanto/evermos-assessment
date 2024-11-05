package product_test

import (
	stderrors "errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/product"
)

func (s *ProductTestSuite) Test_GetProducts() {
	s.Run("fail get products", func() {
		s.repo.EXPECT().GetProducts(s.ctx).Return(nil, stderrors.New("database down"))

		products, err := s.svc.GetProducts(s.ctx)
		s.Equal(errors.InternalServerError, err.Type())
		s.Equal([]product.Product(nil), products)
	})

	s.Run("success get products", func() {
		productsData := []dbgen.GetProductsRow{
			{
				ID:          1,
				Name:        "Product 1",
				Description: "Description 1",
				Price:       10000,
				Quantity:    10,
				ShopID:      1,
				WarehouseID: 1,
			},
			{
				ID:          2,
				Name:        "Product 2",
				Description: "Description 2",
				Price:       20000,
				Quantity:    20,
				ShopID:      2,
				WarehouseID: 2,
			},
		}
		s.repo.EXPECT().GetProducts(s.ctx).Return(productsData, nil)

		products, err := s.svc.GetProducts(s.ctx)
		s.Nil(err)
		s.Equal([]product.Product{
			{
				ID:          1,
				Name:        "Product 1",
				Description: "Description 1",
				Price:       10000,
				Quantity:    10,
				ShopID:      1,
				WarehouseID: 1,
			},
			{
				ID:          2,
				Name:        "Product 2",
				Description: "Description 2",
				Price:       20000,
				Quantity:    20,
				ShopID:      2,
				WarehouseID: 2,
			},
		}, products)
	})
}
