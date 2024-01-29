package service

import (
	"github.com/google/uuid"
	"hexagonal/application/contract"
	"hexagonal/domain"
)

type CreateProductService struct {
	productRepository contract.ProductRepositoryInterface
}

func NewCreateProductService(productRepository contract.ProductRepositoryInterface) *CreateProductService {
	return &CreateProductService{
		productRepository: productRepository,
	}
}

type InputCreateProduct struct {
	Name  string
	Price float64
}

func (c *CreateProductService) Execute(input InputCreateProduct) {
	product := &domain.Product{
		ID:    uuid.New().String(),
		Name:  input.Name,
		Price: input.Price,
	}
	_, err := product.Validate()
	if err != nil {
		panic(err)
	}
	err = c.productRepository.Save(product)
	if err != nil {
		panic(err)
	}
}
