package service

import (
	"errors"
	"propay/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
	tk   *TokenService
}

func NewAuthService(r *repository.UserRepository, tk *TokenService) *AuthService {
	return &AuthService{
		repo: r,
		tk:   tk,
	}
}

func (auth *AuthService) Login(email, password string) (string, error) {
	user, err := auth.repo.FindByEmail(email)

	if err != nil {
		return "", errors.New("Email ou senha inválidos")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("Email ou senha inválidos")
	}

	token, err := auth.tk.CreateToken(user)

	if err != nil {
		return "", err
	}

	return token, nil
}
