/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/handlers"
	"github.com/spf13/cobra"
)

// renamePartitionCmd represents the renamePartition command
var renamePartitionCmd = &cobra.Command{
	Use:   "renamePartition",
	Short: "Rename a partition",
	Long:  `Rename a partition`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("renamePartition called")

		oldP, err := cmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		newP, err := cmd.Flags().GetString("to")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if newP == "" {
			fmt.Println("empty to name provided")
			return
		}

		handlers.RenamePartitionHandler(oldP, newP)

	},
}

func init() {
	rootCmd.AddCommand(renamePartitionCmd)

	renamePartitionCmd.Flags().StringP("to", "t", "", "From partition name")

}
