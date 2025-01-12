package repository

import "main/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
