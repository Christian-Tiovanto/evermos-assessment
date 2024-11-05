package shop_test

import (
	stderrors "errors"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
)

func (s *ShopTestSuite) Test_GetWarehouses() {
	s.Run("fail get warehouse", func() {
		shopID := int64(1)
		s.repo.EXPECT().GetWarehousesByShopID(s.ctx, shopID).Return(nil, stderrors.New("database down"))

		_, err := s.svc.GetWarehouses(s.ctx, shopID)
		s.Equal(errors.InternalServerError, err.Type())
	})

	s.Run("success get warehouse", func() {
		shopID := int64(1)
		warehousesDB := []dbgen.GetWarehousesByShopIDRow{
			{
				ID:     1,
				Name:   "Warehouse 1",
				ShopID: 1,
				Status: pgtype.Bool{Bool: true, Valid: true},
			},
		}
		s.repo.EXPECT().GetWarehousesByShopID(s.ctx, shopID).Return(warehousesDB, nil)

		res, err := s.svc.GetWarehouses(s.ctx, shopID)
		s.Nil(err)
		s.Equal([]warehouse.Warehouse{
			{
				ID:     1,
				Name:   "Warehouse 1",
				ShopID: 1,
				Status: true,
			},
		}, res)
	})
}
