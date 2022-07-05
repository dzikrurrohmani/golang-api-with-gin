package usecase

import (
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/model"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/repository"
)

type CategoryUseCase interface {
	AddNewCategory(newCategory *model.Category) error
	GetAllCategory() []model.Category
}

type categoryUseCase struct {
	categoryRepo repository.CategoryRepository
}

func (c *categoryUseCase) AddNewCategory(newCategory *model.Category) error {
	return c.categoryRepo.Add(newCategory)
}

func (c *categoryUseCase) GetAllCategory() []model.Category {
	return c.categoryRepo.Retrieve()
}

func NewCategoryUseCase(categoryRepo repository.CategoryRepository) CategoryUseCase {
	usecase := new(categoryUseCase)
	usecase.categoryRepo = categoryRepo
	return usecase
}