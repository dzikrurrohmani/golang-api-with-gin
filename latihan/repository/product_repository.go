package repository

import (
	"errors"

	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/model"
)

type ProductRepository interface {
	Add(newProduct *model.Product) error
	Retrieve() []model.Product
	FindById(id string, dest *model.Product) error
}

type productRepository struct {
	productDb []model.Product
}

func (p *productRepository) Add(newProduct *model.Product) error {
	p.productDb = append(p.productDb, *newProduct)
	if p.productDb[len(p.productDb)-1].ProductId != (*newProduct).ProductId {
		return errors.New("product tidak bertambah")
	}
	return nil
}

func (p *productRepository) Retrieve() []model.Product {
	return p.productDb
}

func (p *productRepository) FindById(id string, dest *model.Product) error {
	for _, product := range p.productDb {
		if product.ProductId == id {
			*dest = product
			return nil
		}
	}
	return errors.New("product not found")
}

func NewProductRepository() ProductRepository {
	repo := new(productRepository)
	return repo
}
