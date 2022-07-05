package usecase

import (
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/model"
	"github.com/dzikrurrohmani/golang-api-with-gin/latihan/repository"
)

type ProductUseCase interface {
	AddNewProduct(newProduct *model.Product) error
	GetAllProduct() []model.Product
	GetProductById(id string, dest *model.Product) error
}

type productUseCase struct {
	productRepo repository.ProductRepository
}

func (p *productUseCase) AddNewProduct(newProduct *model.Product) error {
	return p.productRepo.Add(newProduct)
}

func (p *productUseCase) GetAllProduct() []model.Product {
	return p.productRepo.Retrieve()
}

func (p *productUseCase) GetProductById(id string, dest *model.Product) error {
	return p.productRepo.FindById(id, dest)
}

func NewProductUseCase(productRepo repository.ProductRepository) ProductUseCase {
	usecase := new(productUseCase)
	usecase.productRepo = productRepo
	return usecase
}