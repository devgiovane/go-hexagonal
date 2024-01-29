package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"hexagonal/application/service"
	"hexagonal/infra/database"
	"hexagonal/infra/repository"
	"strconv"
)

var connection = database.NewMySQLConnection()

var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A cli for execute commands in application",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch action {
		case "create":
			productRepository := repository.NewProductRepository(connection)
			createProductService := service.NewCreateProductService(productRepository)
			price, _ := strconv.ParseFloat(params[1], 64)
			createProductService.Execute(service.InputCreateProduct{Name: params[0], Price: price})
			fmt.Println("User created")
		case "get":
			productRepository := repository.NewProductRepository(connection)
			getProductService := service.NewProductService(productRepository)
			response := getProductService.Execute(params[0])
			fmt.Println("User", response)
		default:
			fmt.Println("Invalid action use -a")
		}
	},
}

var action string
var params []string

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "", "Create/Get product")
	cliCmd.Flags().StringSliceVarP(&params, "params", "p", []string{}, "")
}
