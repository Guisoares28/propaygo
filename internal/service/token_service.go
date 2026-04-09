package service

import (
	"errors"
	"propay/internal/model"
	"propay/internal/repository"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
	repo         *repository.UserRepository
	chaveSecreta string
}

/*
Está função cria o objeto TokenService, porém ela precisa que a váriavel de ambiente JWT_SECRET
esteja criada dentro do arquivo .env, caso não esteja criada ele retornará um erro.
*/
func NewTokenService(r *repository.UserRepository) (*TokenService, error) {

	chave, error := os.LookupEnv("JWT_SECRET")

	if !error {
		return nil, errors.New("Não foi possível criar o tokenService, crie váriavel JWT_SECRET vazia ou inexistente")
	}

	return &TokenService{
		repo:         r,
		chaveSecreta: chave,
	}, nil
}

/*
Está função gera um token armazenando o id do usuário e o tempo de expiração
*/
func (ts *TokenService) CreateToken(user *model.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	//criação do token com algoritmo HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//assinando token
	tokenString, err := token.SignedString([]byte(ts.chaveSecreta))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (ts *TokenService) DecodedToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(tk *jwt.Token) (interface{}, error) {
		//verifica o método de assinatura
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Método de assinatura inesperado")
		}
		return []byte(ts.chaveSecreta), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, errors.New("Token expirado, faça login novamente")
		}

		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, exist := claims["user_id"]

		if !exist {
			return 0, errors.New("ID do usuário não foi encontrado no token")
		}

		userIDFloat, ok := userID.(float64)

		if !ok {
			return 0, errors.New("formato de user_id inválido")
		}

		return uint(userIDFloat), nil
	}

	return 0, errors.New("Token inválido")
}
