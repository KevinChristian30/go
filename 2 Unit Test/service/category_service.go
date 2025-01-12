package service

import (
	"errors"
	"main/entity"
	"main/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return category, errors.New("not found")
	} else {
		return category, nil
	}
}
