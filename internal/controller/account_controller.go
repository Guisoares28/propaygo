package controller

import (
	"net/http"
	"propay/internal/controller/dto"
	"propay/internal/service"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *service.AccountService
}

func NewAccountController(accountService *service.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

func (ac AccountController) RegisterAccount(c *gin.Context) {
	var account *dto.AccountRequestDto

	id, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error: ": "Erro ao recuperar o ID",
		})
		return
	}

	userID := id.(uint)

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	accountEntity, err := ac.accountService.Create(
		userID,
		account.Title,
		account.Amount,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error ": "Erro ao criar a conta",
		})
		return
	}

	c.JSON(http.StatusCreated, accountEntity)
}
