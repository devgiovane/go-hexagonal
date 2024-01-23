package contract

import "hexagonal/domain"

type ProductRepositoryInterface interface {
	Save(product *domain.Product) error
	Get(id string) (*domain.Product, error)
}
