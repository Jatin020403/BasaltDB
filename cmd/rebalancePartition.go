/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// rebalancePartitionCmd represents the rebalancePartition command
var rebalancePartitionCmd = &cobra.Command{
	Use:   "rebalancePartition",
	Short: "rebalance the partition",
	Long: `Rebalances the partition to the required number. 
To rebalance create a new_config.yaml file to with the required 
configuration in the current Partition directory.`,
	Run: func(cmd *cobra.Command, args []string) {

		partitionName, err := cmd.Flags().GetString("use")
		partition, err := database.CollectPartition(partitionName)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		err = database.RebalancePartition(partition)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(rebalancePartitionCmd)
}
