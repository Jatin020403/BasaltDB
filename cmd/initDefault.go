/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/handlers"
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

		handlers.InitDefaultHandler(partitionName, n)
	},
}

func init() {
	rootCmd.AddCommand(initDefaultCmd)
	initDefaultCmd.PersistentFlags().IntP("size", "n", 3, "size of initial partition")
}
