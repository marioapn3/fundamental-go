package repository

import "belajar-golang-unit-test/entity"

type CategoryReporitory interface {
	FindByID(ID string) *entity.Category
}
