package user

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/dependencies/postgresql"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// UpdateUserParams defines the parameters for updating a user
type UpdateUserParams struct {
	ID       int64
	Password string
	Email    string
	Phone    string
}

// UpdateUser is a method to update user details
func (u *userImpl) UpdateUser(ctx context.Context, arg UpdateUserParams) *errors.Error {
	// Update user data in the database
	hashedPassword, err := u.GeneratePassword([]byte(arg.Password))
	if err != nil {
		return err
	}

	errDB := u.repo.UpdateUserById(ctx, dbgen.UpdateUserByIdParams{
		ID:       arg.ID,
		Password: hashedPassword,
		Email:    arg.Email,
		Phone:    arg.Phone,
	})

	if errDB != nil {
		if postgresql.PgErrCode(errDB) == postgresql.ErrUniqueViolationCode {
			return errors.BadRequest("email/phone already in use")
		}
		return errors.InternalServer(pkgerrors.WithStack(errDB))
	}

	return nil
}
