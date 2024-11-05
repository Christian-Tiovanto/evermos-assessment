package shop_test

import (
	stderrors "errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/warehouse"
)

func (s *ShopTestSuite) Test_CreateWarehouse() {
	s.Run("fail create warehouse", func() {
		arg := warehouse.Warehouse{
			Name:   "Warehouse 1",
			ShopID: 1,
		}

		s.repo.EXPECT().CreateWarehouse(s.ctx, dbgen.CreateWarehouseParams{
			Name:   arg.Name,
			ShopID: arg.ShopID,
		}).Return(stderrors.New("database down"))

		err := s.svc.CreateWarehouse(s.ctx, arg)
		s.Equal(errors.InternalServerError, err.Type())
	})

	s.Run("success create warehouse", func() {
		arg := warehouse.Warehouse{
			Name:   "Warehouse 1",
			ShopID: 1,
		}
		s.repo.EXPECT().CreateWarehouse(s.ctx, dbgen.CreateWarehouseParams{
			Name:   arg.Name,
			ShopID: arg.ShopID,
		}).Return(nil)

		err := s.svc.CreateWarehouse(s.ctx, arg)
		s.Nil(err)
	})
}
