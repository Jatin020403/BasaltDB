/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// deletePartitionCmd represents the deletePartition command
var deletePartitionCmd = &cobra.Command{
	Use:   "deletePartition",
	Short: "Deletes Partitions",
	Long:  `Enter the partition name to delete the partition if it exists`,
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

		err = database.DeletePartition(partition)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println("Partition " + partitionName + " deleted")

	},
}

func init() {
	rootCmd.AddCommand(deletePartitionCmd)
}
