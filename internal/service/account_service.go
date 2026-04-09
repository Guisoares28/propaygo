package service

import (
	"propay/internal/model"
	"propay/internal/repository"
)

type AccountService struct {
	accountRepository *repository.AccountRepository
}

func NewAccountService(accountRepository *repository.AccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (as *AccountService) Create(userID uint, title string, amount float64) (*model.Account, error) {
	account := &model.Account{
		UserID: userID,
		Title:  title,
		Amount: amount,
	}
	err := as.accountRepository.Create(account)

	if err != nil {
		return nil, err
	}

	return account, nil
}
