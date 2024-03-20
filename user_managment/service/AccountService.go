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

func (service *AccountService) BlockAccount(id string) error {
	err := service.AccountRepo.BlockById(id)
	if err != nil {
		return err
	}
	return nil
}
func (service *AccountService) FindAccountByUsernameAndPassword(username, password string) (*model.Account, error) {
	// Perform a database query to find the account by username and password
	account, err := service.AccountRepo.FindByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}
	return account, nil
}
