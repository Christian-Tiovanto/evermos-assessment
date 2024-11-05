package shop_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/mughieams/evermos-assessment/app/usecase/shop"
	mock_shop "github.com/mughieams/evermos-assessment/test/mock/usecase/shop"
)

type ShopTestSuite struct {
	suite.Suite

	ctx  context.Context
	ctrl *gomock.Controller

	repo *mock_shop.MockRepository
	svc  shop.Usecase
}

var (
	_ suite.TestingSuite      = (*ShopTestSuite)(nil)
	_ suite.SetupAllSuite     = (*ShopTestSuite)(nil)
	_ suite.SetupTestSuite    = (*ShopTestSuite)(nil)
	_ suite.TearDownTestSuite = (*ShopTestSuite)(nil)
)

func (s *ShopTestSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *ShopTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_shop.NewMockRepository(s.ctrl)

	s.svc = shop.NewShop(s.repo)
}

func (s *ShopTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func Test_Shop(t *testing.T) {
	suite.Run(t, new(ShopTestSuite))
}
