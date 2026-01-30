package service

import (
	"context"
	"errors"
	"time"

	"ecommerce/internal/config"
	"ecommerce/internal/domain/entity"
	"ecommerce/internal/domain/repository"
	"ecommerce/pkg/utils"
)

type AuthService interface {
	Register(ctx context.Context, username, email, password, firstName, lastName string) error
	Login(ctx context.Context, email, password string) (string, error)
	RequestPasswordReset(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, token, newPassword string) error
}

type authService struct {
	userRepo repository.UserRepository
	cfg      config.Config
}

func NewAuthService(userRepo repository.UserRepository, cfg config.Config) AuthService {
	return &authService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *authService) Register(ctx context.Context, username, email, password, firstName, lastName string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &entity.User{
		Username:  username,
		Email:     email,
		Password:  hashedPassword,
		FirstName: firstName,
		LastName:  lastName,
	}

	return s.userRepo.Create(ctx, user)
}

func (s *authService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	expiresIn, _ := time.ParseDuration(s.cfg.JWT.ExpiresIn)
	if expiresIn == 0 {
		expiresIn = time.Hour * 24
	}

	token, err := utils.GenerateToken(user.ID, user.Role, s.cfg.JWT.Secret, expiresIn)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *authService) RequestPasswordReset(ctx context.Context, email string) error {
	_, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}
	return nil
}

func (s *authService) ResetPassword(ctx context.Context, token, newPassword string) error {
	return nil
}
