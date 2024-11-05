package user

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	pkgerrors "github.com/pkg/errors"

	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

const (
	defaultExpireTime = 24 * time.Hour
)

// LoginParams defines the user login params
type LoginParams struct {
	Email    string
	Phone    string
	Password string
}

// Login is a method to login user
func (u *userImpl) Login(ctx context.Context, arg LoginParams) (string, *errors.Error) {
	user, err := u.repo.GetUserByEmailOrPhone(ctx, dbgen.GetUserByEmailOrPhoneParams{
		Email: arg.Email,
		Phone: arg.Phone,
	})
	if err != nil {
		return "", errors.InternalServer(pkgerrors.WithStack(err))
	}

	if err := u.ComparePassword([]byte(user.Password), []byte(arg.Password)); err != nil {
		return "", err
	}

	token, err := u.generateToken(user)
	if err != nil {
		return "", errors.InternalServer(pkgerrors.WithStack(err))
	}

	return token, nil
}

func (u *userImpl) generateToken(user dbgen.Users) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"phone": user.Phone,
		"exp":   time.Now().Add(defaultExpireTime).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.cfg.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
