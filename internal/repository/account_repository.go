package repository

import (
	"propay/internal/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{DB: db}
}

// Criar conta
func (r *AccountRepository) Create(account *model.Account) error {
	return r.DB.Create(account).Error
}

/*
Função responsável por buscar contas com base no ID do usuário
Essa função recebe um ID do tipo uint
E retorna uma lista de model.Account
*/
func (r *AccountRepository) FindAllAccountByUserID(userID uint) ([]model.Account, error) {
	var accounts []model.Account
	err := r.DB.Where("user_id = ?", userID).Find(&accounts).Error

	return accounts, err
}

/*
Está função muda o status da conta dependendo do status atual
if status for pending ele altera para paid
if status for paid altera para pending
*/
func (r *AccountRepository) ChangeStatus(accountID uint) error {
	var account *model.Account

	if err := r.DB.Where("id = ?", accountID).First(&account).Error; err != nil {
		return err
	}

	//Inversão de status
	if account.Status == "pending" {
		account.Status = "paid"
	} else {
		account.Status = "pending"
	}

	return r.DB.Save(account).Error
}

/*
Função para deletar uma conta pelo ID informado, caso não existe uma conta com o ID informado será retornado um erro
*/
func (r *AccountRepository) DeleteAccountByID(accountID uint) error {
	var account *model.Account
	if err := r.DB.Where("id = ?", accountID).First(&account).Error; err != nil {
		return err
	}

	return r.DB.Unscoped().Delete(&account).Error
}

func (r *AccountRepository) FindAccountByID(accountID uint, userID uint) (*model.Account, error) {
	var account *model.Account

	err := r.DB.Where("id = ? and user_id = ?", accountID, userID).First(&account).Error

	if err != nil {
		return nil, err
	}

	return account, nil
}
