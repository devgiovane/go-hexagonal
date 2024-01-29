package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hexagonal/application/service"
	"hexagonal/infra/repository"
	"strconv"
)

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A cli for execute commands in application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch action {
		case "create":
			productRepository := repository.NewProductRepository(connection)
			createProductService := service.NewCreateProductService(productRepository)
			name := product[0]
			price, _ := strconv.ParseFloat(product[1], 64)
			response := createProductService.Execute(service.InputCreateProduct{Name: name, Price: price})
			fmt.Println("User created", response)
		case "get":
			productRepository := repository.NewProductRepository(connection)
			getProductService := service.NewGetProductService(productRepository)
			response := getProductService.Execute(service.InputGetProduct{ID: id})
			fmt.Println("User", response)
		default:
			fmt.Println("Invalid action use -a")
		}
	},
}

var action string
var id string
var product []string

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "", "Create/Get product")
	cliCmd.Flags().StringVarP(&id, "id", "i", "", "Product ID")
	cliCmd.Flags().StringSliceVarP(&product, "params", "p", []string{}, "Product")
}
