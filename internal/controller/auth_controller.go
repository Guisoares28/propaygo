package controller

import (
	"net/http"
	"propay/internal/controller/dto"
	"propay/internal/service"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var login *dto.LoginRequest

	err := c.ShouldBindJSON(&login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "dados inválidos",
		})
		return
	}

	token, err := ac.authService.Login(login.Email, login.Password)

	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
