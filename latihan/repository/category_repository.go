package repository

import (
	"errors"

	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/model"
)

type CategoryRepository interface {
	Add(newCategory *model.Category) error
	Retrieve() []model.Category
}

type categoryRepository struct {
	categoryDb []model.Category
}

func (p *categoryRepository) Add(newCategory *model.Category) error {
	p.categoryDb = append(p.categoryDb, *newCategory)
	if p.categoryDb[len(p.categoryDb)-1].CategoryId != (*newCategory).CategoryId {
		return errors.New("category tidak bertambah")
	}
	return nil
}

func (p *categoryRepository) Retrieve() []model.Category {
	return p.categoryDb
}

func NewCategoryRepository() CategoryRepository {
	repo := new(categoryRepository)
	repo.categoryDb = []model.Category{
		{
			CategoryId:   "C0001",
			CategoryName: "Food",
		},
		{
			CategoryId:   "C0002",
			CategoryName: "Drink",
		},
		{
			CategoryId:   "C0003",
			CategoryName: "Other",
		},
	}
	return repo
}
