package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type AccountService struct {
	AccountRepo *repo.AccountRepository
}

func (service *AccountService) FindAccount(id string) (*model.Account, error) {
	account, err := service.AccountRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &account, nil
}

func (service *AccountService) FindAllAccounts() (*[]model.Account, error) {
	accounts, err := service.AccountRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("no accounts found")
	}
	return &accounts, nil
}

func (service *AccountService) Create(account *model.Account) error {
	err := service.AccountRepo.CreateAccount(account)
	if err != nil {
		return err
	}
	return nil
}
