package user

import (
	"context"

	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	"github.com/mughieams/evermos-assessment/app/common/util"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// User defines the User usecase interface
type User struct {
	ID          int64
	Email        string
	Phone string
}

// GetUsers is a method to get Users
func (p *userImpl) GetUsers(ctx context.Context) ([]User, *errors.Error) {
	users, err := p.repo.GetUsers(ctx)
	if err != nil {
		return nil, errors.InternalServer(pkgerrors.WithStack(err))
	}

	return util.SlimMap(users, func(data dbgen.GetUsersRow) User {
		return User{
			ID:          data.ID,
			Email:        data.Email,
			Phone: data.Phone,
		}
	}), nil
}
