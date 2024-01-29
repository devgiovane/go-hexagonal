package service

import "hexagonal/application/contract"

type OutputGetProduct struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type GetProductService struct {
	productRepository contract.ProductRepositoryInterface
}

func NewGetProductService(productRepository contract.ProductRepositoryInterface) *GetProductService {
	return &GetProductService{
		productRepository: productRepository,
	}
}

type InputGetProduct struct {
	ID string
}

func (g *GetProductService) Execute(input InputGetProduct) OutputGetProduct {
	product, err := g.productRepository.Get(input.ID)
	if err != nil {
		panic(err)
	}
	return OutputGetProduct{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}
