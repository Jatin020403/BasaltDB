/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// initDefaultCmd represents the initDefault command
var initDefaultCmd = &cobra.Command{
	Use:   "initDefault",
	Short: "Initialise default partition",
	Long:  `Get default partition working.`,
	Run: func(cmd *cobra.Command, args []string) {

		default_partition := "default"

		database.CreateTemplate(default_partition)
		database.CreatePartition(default_partition)

		fmt.Println("Partition " + default_partition + " initialised")
	},
}

func init() {
	rootCmd.AddCommand(initDefaultCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initDefaultCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initDefaultCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
