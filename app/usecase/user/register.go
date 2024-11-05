package user

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// RegisterParams defines the user register params
type RegisterParams struct {
	Name     string
	Email    string
	Phone    string
	Password string
}

// Register is a method to register user
func (u *userImpl) Register(ctx context.Context, arg RegisterParams) *errors.Error {
	hashedPassword, err := u.GeneratePassword([]byte(arg.Password))
	if err != nil {
		return err
	}

	errDB := u.repo.CreateUser(ctx, dbgen.CreateUserParams{
		Email:    arg.Email,
		Phone:    arg.Phone,
		Password: hashedPassword,
	})
	if errDB != nil {
		if postgresql.PgErrCode(errDB) == postgresql.ErrUniqueViolationCode {
			return errors.BadRequest("email/phone already registered")
		}
		return errors.InternalServer(pkgerrors.WithStack(errDB))
	}

	return nil
}
