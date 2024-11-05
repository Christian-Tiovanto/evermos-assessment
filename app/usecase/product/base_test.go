package product_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/mughieams/evermos-assessment/app/usecase/product"
	mock_product "github.com/mughieams/evermos-assessment/test/mock/usecase/product"
)

type ProductTestSuite struct {
	suite.Suite

	ctx  context.Context
	ctrl *gomock.Controller

	repo *mock_product.MockRepository
	svc  product.Usecase
}

var (
	_ suite.TestingSuite      = (*ProductTestSuite)(nil)
	_ suite.SetupAllSuite     = (*ProductTestSuite)(nil)
	_ suite.SetupTestSuite    = (*ProductTestSuite)(nil)
	_ suite.TearDownTestSuite = (*ProductTestSuite)(nil)
)

func (s *ProductTestSuite) SetupSuite() {
	s.ctx = context.Background()
}

func (s *ProductTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_product.NewMockRepository(s.ctrl)

	s.svc = product.NewProduct(s.repo)
}

func (s *ProductTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func Test_Product(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}
