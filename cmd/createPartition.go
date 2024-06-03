/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// createPartitionCmd represents the createPartition command
var createPartitionCmd = &cobra.Command{
	Use:   "createPartition",
	Short: "Creates partition",
	Long:  `Enter the partition name to create partition`,
	Run: func(cmd *cobra.Command, args []string) {
		PartitionName, err := cmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
		}
		if PartitionName == "" || err != nil {
			fmt.Println("Invalid Partition. Provide a string")
			return
		}

		err = database.CreatePartition(PartitionName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Partition " + PartitionName + " created")

	},
}

func init() {
	rootCmd.AddCommand(createPartitionCmd)
}
