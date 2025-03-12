// Package user provides usecase implementation for user
package user

import (
	"context"

	pkgerrors "github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"

	"github.com/mughieams/evermos-assessment/app/common/config"
	"github.com/mughieams/evermos-assessment/app/common/errors"
	dbgen "github.com/mughieams/evermos-assessment/app/repository/postgresql/db"
)

// Usecase defines the user usecase interface
type Usecase interface {
	GetUsers(ctx context.Context) ([]User, *errors.Error)
	GeneratePassword(password []byte) (string, *errors.Error)
	ComparePassword(hashedPassword, password []byte) *errors.Error
	Login(ctx context.Context, arg LoginParams) (string, *errors.Error)
	Register(ctx context.Context, arg RegisterParams) *errors.Error
	UpdateUser(ctx context.Context, arg UpdateUserParams) *errors.Error
}

// Repository defines repository interface for user usecase
type Repository interface {
	GetUsers(ctx context.Context) ([]dbgen.GetUsersRow, error)
	CreateUser(ctx context.Context, arg dbgen.CreateUserParams) error
	GetUserByEmailOrPhone(ctx context.Context, arg dbgen.GetUserByEmailOrPhoneParams) (dbgen.Users, error)
	UpdateUserById(ctx context.Context, arg dbgen.UpdateUserByIdParams) error
}

type userImpl struct {
	cfg  config.JWT
	repo Repository
}

var _ Usecase = (*userImpl)(nil)

// NewUser instantiate a new user usecase
func NewUser(cfg config.JWT, repo Repository) Usecase {
	return &userImpl{
		cfg:  cfg,
		repo: repo,
	}
}

func (u *userImpl) GeneratePassword(password []byte) (string, *errors.Error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.InternalServer(pkgerrors.WithStack(err))
	}

	return string(hashedPassword), nil
}

func (u *userImpl) ComparePassword(hashedPassword, password []byte) *errors.Error {
	if err := bcrypt.CompareHashAndPassword(hashedPassword, password); err != nil {
		return errors.Authentication("invalid email/phone or password")
	}
	return nil
}
