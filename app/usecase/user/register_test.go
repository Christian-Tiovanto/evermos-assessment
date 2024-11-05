package user_test

import (
	stderrors "errors"

	"github.com/jackc/pgx/v5/pgconn"
	"go.uber.org/mock/gomock"

	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	"github.com/mughieams/evermos-assessment/app/usecase/user"
)

func (s *UserTestSuite) Test_Register() {
	s.Run("fail_create_user", func() {
		hashedPassword, err := s.svc.GeneratePassword([]byte("password"))
		s.Nil(err)

		s.repo.EXPECT().CreateUser(s.ctx, gomock.Any()).Return(stderrors.New("database down"))

		err = s.svc.Register(s.ctx, user.RegisterParams{
			Name:     "Sample",
			Email:    "sample@mail.com",
			Phone:    "08123456789",
			Password: hashedPassword,
		})
		s.Equal(errors.InternalServerError, err.Type())
	})

	s.Run("fail_email_phone_registered", func() {
		hashedPassword, err := s.svc.GeneratePassword([]byte("password"))
		s.Nil(err)

		s.repo.EXPECT().CreateUser(s.ctx, gomock.Any()).Return(&pgconn.PgError{Code: postgresql.ErrUniqueViolationCode})

		err = s.svc.Register(s.ctx, user.RegisterParams{
			Name:     "Sample",
			Email:    "sample@mail.com",
			Phone:    "08123456789",
			Password: hashedPassword,
		})
		s.Equal(errors.BadRequestError, err.Type())
	})

	s.Run("success", func() {
		hashedPassword, err := s.svc.GeneratePassword([]byte("password"))
		s.Nil(err)

		s.repo.EXPECT().CreateUser(s.ctx, gomock.Any()).Return(nil)

		err = s.svc.Register(s.ctx, user.RegisterParams{
			Name:     "Sample",
			Email:    "sample@mail.com",
			Phone:    "08123456789",
			Password: hashedPassword,
		})
		s.Nil(err)
	})
}
