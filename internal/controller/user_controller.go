package controller

import (
	"net/http"
	"propay/internal/model"
	"propay/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	us *service.UserService
}

func NewUserController(service *service.UserService) *UserController {
	return &UserController{us: service}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	var user *model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao cadastrar usuário",
		})
		return
	}

	user, err := uc.us.CreateUser(
		user.Name,
		user.Email,
		user.Password,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao cadastrar usuário",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
