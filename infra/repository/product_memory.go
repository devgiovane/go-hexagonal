package repository

import (
	"errors"
	"hexagonal/domain"
	"hexagonal/infra/database"
)

type ProductRepositoryMemory struct {
	Connection database.Connection[string]
	Products   []*domain.Product
}

func NewProductRepositoryMemory(connection database.Connection[string]) *ProductRepositoryMemory {
	return &ProductRepositoryMemory{
		Connection: connection,
		Products:   []*domain.Product{},
	}
}

func (p *ProductRepositoryMemory) Save(product *domain.Product) error {
	p.Products = append(p.Products, product)
	return nil
}

func (p *ProductRepositoryMemory) Get(id string) (*domain.Product, error) {
	for _, p := range p.Products {
		if p.ID == id {
			return p, nil
		}
	}
	return &domain.Product{}, errors.New("product not found")
}
