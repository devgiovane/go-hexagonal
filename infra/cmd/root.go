package cmd

import (
	"github.com/spf13/cobra"
	"hexagonal/infra/database"
	"os"
)

var connection = database.NewMySQLConnection()

var rootCmd = &cobra.Command{
	Use:   "hexagonal",
	Short: "Application for study architecture hexagonal",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
