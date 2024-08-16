package repository

import (
	"belajar-golang-unit-test/entity"

	"github.com/stretchr/testify/mock"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindByID(ID string) *entity.Category {
	arguments := repository.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
