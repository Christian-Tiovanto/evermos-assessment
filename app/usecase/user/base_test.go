package user_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"

	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/usecase/user"
	mock_user "github.com/mughieams/evermos-assessment/test/mock/usecase/user"
)

type UserTestSuite struct {
	suite.Suite

	ctx  context.Context
	ctrl *gomock.Controller

	cfg config.JWT

	repo    *mock_user.MockRepository
	mockSvc *mock_user.MockUsecase
	svc     user.Usecase
}

var (
	_ suite.TestingSuite      = (*UserTestSuite)(nil)
	_ suite.SetupAllSuite     = (*UserTestSuite)(nil)
	_ suite.SetupTestSuite    = (*UserTestSuite)(nil)
	_ suite.TearDownTestSuite = (*UserTestSuite)(nil)
)

func (s *UserTestSuite) SetupSuite() {
	s.ctx = context.Background()
	s.cfg = config.JWT{
		Secret: "saltandpaper",
	}
}

func (s *UserTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.repo = mock_user.NewMockRepository(s.ctrl)

	s.mockSvc = mock_user.NewMockUsecase(s.ctrl)
	s.svc = user.NewUser(s.cfg, s.repo)
}

func (s *UserTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func Test_UserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
