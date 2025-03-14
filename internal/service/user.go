package service

import (
	"context"
	"errors"
	"github.com/huangyul/go-webook/internal/domain"
	"github.com/huangyul/go-webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserNotFound          = repository.ErrUserNotFound
	ErrUserAlreadyExists     = repository.ErrUserEmailAlreadyExists
	ErrUserEmailIllegally    = errors.New("user email illegal")
	ErrUserPasswordIllegally = errors.New("user password illegal")
	ErrUserPasswordNotMatch  = errors.New("user password not match")
)

type UserService interface {
	RegisterByEmail(ctx context.Context, email string, password string) error
	LoginByEmail(ctx context.Context, email string, password string) (*domain.User, error)
}

var _ UserService = (*userService)(nil)

type userService struct {
	repo repository.UserRepository
}

func (u *userService) RegisterByEmail(ctx context.Context, email string, password string) error {
	passwordStr, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return u.repo.Insert(ctx, &domain.User{
		Email:    email,
		Password: string(passwordStr),
	})
}

func (u *userService) LoginByEmail(ctx context.Context, email string, password string) (*domain.User, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, ErrUserPasswordNotMatch
	}
	return user, nil
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}
