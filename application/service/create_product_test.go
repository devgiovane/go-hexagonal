package service_test

import (
	"github.com/stretchr/testify/require"
	"hexagonal/application/service"
	"hexagonal/infra/database"
	"hexagonal/infra/repository"
	"testing"
)

func TestCreateProductService_Execute(t *testing.T) {
	connection := database.NewMemoryConnection()
	productRepository := repository.NewProductRepositoryMemory(connection)
	createProductService := service.NewCreateProductService(productRepository)
	id := createProductService.Execute(service.InputCreateProduct{Name: "Beer", Price: 1.99})
	require.NotEmpty(t, id)
	getProductService := service.NewGetProductService(productRepository)
	product := getProductService.Execute(service.InputGetProduct{ID: id})
	require.Equal(t, "Beer", product.Name)
	require.Equal(t, 1.99, product.Price)
}
