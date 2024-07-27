/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/handlers"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise configuration for a partition",
	Long:  `Get template for a configuration YAML file for creating a partition.`,
	Run: func(cmd *cobra.Command, args []string) {
		partitionName, err := cmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
		}

		n, err := cmd.Flags().GetInt("size")
		if err != nil {
			fmt.Println(err.Error())
		}

		handlers.InitHandler(partitionName, n)

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().IntP("size", "n", 3, "size of initial partition")

}
