package repo

import (
	"blogservice.com/model"
	"gorm.io/gorm"
)

type CommentRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CommentRepository) FindById(id string) (model.Comment, error) {
	comment := model.Comment{}
	dbResult := repo.DatabaseConnection.First(&comment, "id = ?", id)
	if dbResult != nil {
		return comment, dbResult.Error
	}
	return comment, nil
}

func (repo *CommentRepository) FindAll() ([]model.Comment, error) {
	var comments = []model.Comment{}
	dbResult := repo.DatabaseConnection.Find(&comments)
	if dbResult != nil {
		return comments, dbResult.Error
	}
	return comments, nil
}

func (repo *CommentRepository) CreateComment(comment *model.Comment) error {
	dbResult := repo.DatabaseConnection.Create(comment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}

func (repo *CommentRepository) UpdateComment(comment *model.Comment) error {
	dbResult := repo.DatabaseConnection.Save(comment)
	if dbResult.Error != nil {
		return dbResult.Error
	}
	println("Rows affected: ", dbResult.RowsAffected)
	return nil
}