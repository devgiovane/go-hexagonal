package domain

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type Product struct {
	ID    string  `valid:"uuidv4"`
	Name  string  `valid:"required"`
	Price float64 `valid:"float"`
}

func NewProduct(name string, price float64) *Product {
	product := Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
	return &product
}

func (p *Product) Validate() (bool, error) {
	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, err
	}
	return true, nil
}
