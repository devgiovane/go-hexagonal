package repository

import (
	"database/sql"
	"hexagonal/domain"
	"hexagonal/infra/database"
)

type ProductRepository struct {
	Connection database.Connection[*sql.DB]
}

func NewProductRepository(connection database.Connection[*sql.DB]) *ProductRepository {
	return &ProductRepository{
		Connection: connection,
	}
}

func (p *ProductRepository) Save(product *domain.Product) error {
	conn := p.Connection.GetConnection()
	stmt, err := conn.Prepare("INSERT INTO product(id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductRepository) Get(id string) (*domain.Product, error) {
	conn := p.Connection.GetConnection()
	stmt, err := conn.Prepare("SELECT id, name, price FROM product WHERE id = ?")
	if err != nil {
		return &domain.Product{}, err
	}
	var product domain.Product
	result := stmt.QueryRow(id)
	err = result.Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return &domain.Product{}, nil
	}
	err = stmt.Close()
	if err != nil {
		return &domain.Product{}, nil
	}
	return &product, nil
}
