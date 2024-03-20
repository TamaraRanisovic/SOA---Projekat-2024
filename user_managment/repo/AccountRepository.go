package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type AccountRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *AccountRepository) FindById(id string) (model.Account, error) {
	account := model.Account{}
	dbResult := repo.DatabaseConnection.First(&account, "id = ?", id)
	if dbResult != nil {
		return account, dbResult.Error
	}
	return account, nil
}

func (repo *AccountRepository) FindAll() ([]model.Account, error) {
	var accounts = []model.Account{}
	dbResult := repo.DatabaseConnection.Find(&accounts)
	if dbResult != nil {
		return accounts, dbResult.Error
	}
	return accounts, nil
}

func (repo *AccountRepository) CreateAccount(account *model.Account) error {
	dbResult := repo.DatabaseConnection.Create(account)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}