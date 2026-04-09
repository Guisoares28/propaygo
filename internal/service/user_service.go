package service

import (
	"errors"
	"propay/internal/model"
	"propay/internal/repository"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) CreateUser(name, email, password string) (*model.User, error) {
	if email == "" || password == "" {
		return nil, errors.New("Email ou Senha inválidos")
	}

	_, err := s.repo.FindByEmail(email)

	if err == nil {
		return nil, errors.New("Já existe um usuário cadastrado")
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err = s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}
