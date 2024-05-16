package repo

import (
	"blogservice/model"

	"gorm.io/gorm"
)

type BlogRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *BlogRepository) FindById(id string) (model.Blog, error) {
	blog := model.Blog{}
	dbResult := repo.DatabaseConnection.Preload("Pictures").First(&blog, "id = ?", id)
	if dbResult.Error != nil {
		return blog, dbResult.Error
	}
	return blog, nil
}

func (repo *BlogRepository) FindAll() ([]model.Blog, error) {
	var blogs []model.Blog

	// Preload Pictures for each Blog
	dbResult := repo.DatabaseConnection.Preload("Pictures").Find(&blogs)
	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return blogs, nil
}

func (repo *BlogRepository) CreateBlog(blog *model.Blog) error {
	dbResult := repo.DatabaseConnection.Create(blog)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}
