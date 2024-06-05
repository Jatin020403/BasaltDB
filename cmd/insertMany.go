/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// insertManyCmd represents the insertMany command
var insertManyCmd = &cobra.Command{
	Use:   "insertMany",
	Short: "Insert many values",
	Long:  `Insert an array of key value pairs into the database`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("insertMany called")
		// partition, err := rootCmd.Flags().GetString("use")
		key, err := cmd.Flags().GetStringSlice("keyArr")
		if err != nil {
			fmt.Println(err.Error())
		}
		value, err := cmd.Flags().GetStringSlice("valueArr")

		if(len(key) != len(value)){
			fmt.Println("length of key and value must be equal")
			return;
		}

		// fmt.Println(partition)
		fmt.Println(key)
		fmt.Println(value)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(insertManyCmd)

	insertManyCmd.Flags().StringSliceP("keyArr", "w", []string{}, "Insert array of keys")
	insertManyCmd.Flags().StringSliceP("valueArr", "q", []string{}, "Insert array of values")
	insertCmd.MarkFlagsRequiredTogether("key", "value")
}
