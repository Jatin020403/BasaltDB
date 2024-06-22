/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialise configuration for a partition",
	Long: `Get template for a configuration YAML file for creating a partition.`,
	Run: func(cmd *cobra.Command, args []string) {
		PartitionName, err := cmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
		}
		if PartitionName == "" || err != nil {
			fmt.Println("Invalid Partition. Provide a string")
			return
		}

		err = database.CreateTemplate(PartitionName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Partition " + PartitionName + " initialised")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
