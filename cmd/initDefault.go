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

// initDefaultCmd represents the initDefault command
var initDefaultCmd = &cobra.Command{
	Use:   "initDefault",
	Short: "Initialise default partition",
	Long:  `Get default partition working.`,
	Run: func(cmd *cobra.Command, args []string) {

		partitionName, err := cmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		n, err := cmd.Flags().GetInt("size")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		var partition models.Partition
		partition.Name = partitionName

		partition, err = database.CreateTemplate(partition, n)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		partition, err = database.CreatePartition(partition)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Partition " + partitionName + " initialised")
	},
}

func init() {
	rootCmd.AddCommand(initDefaultCmd)
	initDefaultCmd.PersistentFlags().IntP("size", "n", 3, "size of initial partition")
}
