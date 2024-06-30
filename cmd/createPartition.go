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
		partitionName, err := cmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
		}
		if partitionName == "" || err != nil {
			fmt.Println("Invalid Partition. Provide a string")
			return
		}

		partition, err := database.CollectPartition(partitionName)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		partition, err = database.CreatePartition(partition)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Partition " + partitionName + " created")
	},
}

func init() {
	rootCmd.AddCommand(createPartitionCmd)
}
