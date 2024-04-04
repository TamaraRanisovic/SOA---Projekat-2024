package service

import (
	"fmt"

	"blogservice.com/model"
	"blogservice.com/repo"
)

type BlogService struct {
	BlogRepo *repo.BlogRepository
}

func (service *BlogService) FindBlog(id string) (*model.Blog, error) {
	blog, err := service.BlogRepo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("menu item with id %s not found", id))
	}
	return &blog, nil
}

func (service *BlogService) FindAllBlogs() (*[]model.Blog, error) {
	blogs, err := service.BlogRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("no blogs found")
	}
	return &blogs, nil
}

func (service *BlogService) Create(blog *model.Blog) error {
	err := service.BlogRepo.CreateBlog(blog)
	if err != nil {
		return err
	}
	return nil
}
