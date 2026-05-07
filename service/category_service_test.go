package service

import (
	"Belajar-Go-UnitTest/entity"
	"Belajar-Go-UnitTest/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepositoryMock = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepositoryMock}

func TestCategoryService_GetCategory(t *testing.T) {
	categoryRepositoryMock.Mock.On("FindByID", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetCategorySuccess(t *testing.T) {
	category := entity.Category{Id: "1", Name: "Gadget"}
	categoryRepositoryMock.Mock.On("FindByID", "1").Return(category)
	categoryResult, err := categoryService.Get("1")
	assert.Nil(t, err)
	assert.NotNil(t, categoryResult)
	assert.Equal(t, category.Id, categoryResult.Id)
	assert.Equal(t, category.Name, categoryResult.Name)
}
