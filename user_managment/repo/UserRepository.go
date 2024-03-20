package repo

import (
	"database-example/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *UserRepository) FindById(id string) (model.User, error) {
	user := model.User{}
	dbResult := repo.DatabaseConnection.First(&user, "id = ?", id)
	if dbResult != nil {
		return user, dbResult.Error
	}
	return user, nil
}

func (repo *UserRepository) FindAll() ([]model.User, error) {
	var users = []model.User{}
	dbResult := repo.DatabaseConnection.Find(&users)
	if dbResult != nil {
		return users, dbResult.Error
	}
	return users, nil
}

func (repo *UserRepository) CreateUser(user *model.User) error {
	dbResult := repo.DatabaseConnection.Create(user)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}