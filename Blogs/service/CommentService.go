package service
import (
	"database-example/model"
	"database-example/repo"
	"fmt"
)

type CommentService struct {
	CommentRepo *repo.CommentRepository
}

func (service *CommentService) FindComment(id string) (*model.Comment, error) {
	comment, err := service.CommentRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &comment, nil
}

func (service *CommentService) FindAllComments() (*[]model.Comment, error) {
	comments, err := service.CommentRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("no users found")
	}
	return &comments, nil
}

func (service *CommentService) Create(comment *model.Comment) error {
	err := service.CommentRepo.CreateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func (service *CommentService) Update(comment *model.Comment) error {
	err := service.CommentRepo.UpdateComment(comment)
	if err != nil {
		return err
	}
	return nil
}