package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"errors"
)

type CategoryService struct {
	Repository repository.CategoryReporitory
}

func (service CategoryService) GetByID(id string) (*entity.Category, error) {
	category := service.Repository.FindByID(id)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}

}
