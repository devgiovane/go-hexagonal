package service

import "hexagonal/application/contract"

type Output struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type GetProductService struct {
	productRepository contract.ProductRepositoryInterface
}

func NewProductService(productRepository contract.ProductRepositoryInterface) *GetProductService {
	return &GetProductService{
		productRepository: productRepository,
	}
}

func (g *GetProductService) Execute(id string) Output {
	product, err := g.productRepository.Get(id)
	if err != nil {
		panic(err)
	}
	return Output{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}
