package shop_test

import (
	stderrors "errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/product"
)

func (s *ShopTestSuite) Test_GetProducts() {
	s.Run("fail get products", func() {
		shopID := int64(1)

		s.repo.EXPECT().GetProductsByShopID(s.ctx, shopID).Return(nil, stderrors.New("database down"))

		products, err := s.svc.GetProducts(s.ctx, shopID)
		s.Nil(products)
		s.Equal(errors.InternalServerError, err.Type())
	})

	s.Run("success get products", func() {
		productsDB := []dbgen.GetProductsByShopIDRow{
			{
				ID:          1,
				Name:        "Product 1",
				Description: "Description 1",
				Price:       10000,
				Quantity:    10,
				ShopID:      1,
				WarehouseID: 1,
			},
		}
		shopID := int64(1)

		s.repo.EXPECT().GetProductsByShopID(s.ctx, shopID).Return(productsDB, nil)

		products, err := s.svc.GetProducts(s.ctx, shopID)
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
		}, products)
	})
}
