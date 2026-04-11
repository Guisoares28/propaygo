package controller

import (
	"net/http"
	"propay/internal/controller/dto"
	"propay/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	accountService *service.AccountService
}

func NewAccountController(accountService *service.AccountService) *AccountController {
	return &AccountController{accountService: accountService}
}

func (ac AccountController) RegisterAccount(c *gin.Context) {
	var account dto.AccountRequestDto

	id, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error: ": "Erro ao recuperar o ID",
		})
		return
	}

	userID := id.(uint)

	if err := c.ShouldBindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao converter json",
		})
		return
	}

	accountEntity, err := ac.accountService.Create(
		userID,
		account.Title,
		account.Amount,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error ": "Erro ao criar a conta",
		})
		return
	}

	c.JSON(http.StatusCreated, accountEntity)
}

func (ac *AccountController) ChangeStatus(c *gin.Context) {
	accountID := c.Param("id")
	idUser, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error: ": "Erro ao recuperar o ID",
		})
		return
	}

	userID := idUser.(uint)

	id, err := strconv.Atoi(accountID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "ID inváldo",
		})
		return
	}

	ac.accountService.ChangeStatus(uint(id), userID)
}

func (ac *AccountController) FindAllAccountsByID(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "ID não encontrado",
		})
		return
	}

	id := userID.(uint)

	accounts, err := ac.accountService.FindAllAccountsByUserID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Erro ao consultar contas, tente novamente mais tarde...",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}
