/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/Jatin020403/BasaltDB/models"
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

		var partition models.Partition
		partition.Name = partitionName

		partition, err = database.CreateTemplate(partition, n)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Partition " + partitionName + " initialised")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.PersistentFlags().IntP("size", "n", 3, "size of initial partition")

}
