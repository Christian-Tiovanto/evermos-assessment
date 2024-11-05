package shop_test

import (
	stderrors "errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/shop"
)

func (s *ShopTestSuite) Test_CreateShop() {
	s.Run("fail create shop", func() {
		arg := shop.CreateParams{
			Name:    "Shop 1",
			Address: "Address 1",
		}

		s.repo.EXPECT().CreateShop(s.ctx, dbgen.CreateShopParams{
			Name:    arg.Name,
			Address: arg.Address,
		}).Return(stderrors.New("database down"))

		err := s.svc.CreateShop(s.ctx, arg)
		s.Equal(errors.InternalServerError, err.Type())
	})

	s.Run("success create shop", func() {
		arg := shop.CreateParams{
			Name:    "Shop 1",
			Address: "Address 1",
		}
		s.repo.EXPECT().CreateShop(s.ctx, dbgen.CreateShopParams{
			Name:    arg.Name,
			Address: arg.Address,
		}).Return(nil)

		err := s.svc.CreateShop(s.ctx, arg)
		s.Nil(err)
	})
}
