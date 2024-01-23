package main

import (
	"fmt"
	"hexagonal/application/service"
	"hexagonal/infra/database"
	"hexagonal/infra/repository"
)

func main() {
	connection := database.NewMySQLConnection()
	productRepository := repository.NewProductRepository(connection)
	getProductService := service.NewProductService(productRepository)
	response := getProductService.Execute("4c4a00c9-20ba-454f-9b1e-050322a999f8")
	fmt.Println(response)
}
