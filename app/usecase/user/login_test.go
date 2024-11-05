package user_test

import (
	stderrors "errors"

	"go.uber.org/mock/gomock"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
	"github.com/mughieams/evermos-assessment/app/usecase/user"
)

func (s *UserTestSuite) Test_Login() {
	s.Run("fail_get_user", func() {
		s.repo.EXPECT().GetUserByEmailOrPhone(s.ctx, gomock.Any()).Return(dbgen.Users{}, stderrors.New("database down"))

		token, err := s.svc.Login(s.ctx, user.LoginParams{
			Email:    "sample@mail.com",
			Password: "password",
		})
		s.Empty(token)
		s.Equal(errors.InternalServerError, err.Type())
	})

	s.Run("fail_password_mismatch", func() {
		s.repo.EXPECT().GetUserByEmailOrPhone(s.ctx, gomock.Any()).Return(dbgen.Users{
			Password: "hashedpassword",
		}, nil)

		token, err := s.svc.Login(s.ctx, user.LoginParams{
			Email:    "sample@mail.com",
			Password: "password",
		})
		s.Empty(token)
		s.Equal(errors.AuthenticationError, err.Type())
	})

	s.Run("success", func() {
		hashPassword, err := s.svc.GeneratePassword([]byte("password"))
		s.Nil(err)

		s.repo.EXPECT().GetUserByEmailOrPhone(s.ctx, gomock.Any()).Return(dbgen.Users{
			ID:       1,
			Email:    "sample@mail.com",
			Password: hashPassword,
		}, nil)

		token, err := s.svc.Login(s.ctx, user.LoginParams{
			Email:    "sample@mail.com",
			Password: "password",
		})
		s.NotEmpty(token)
		s.Nil(err)
	})
}
