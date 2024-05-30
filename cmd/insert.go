package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert into Database",
	Long:  `This command inserts into database.`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 2 {
			fmt.Println("Length of args must be 2")
			return
		}
		if database.InsertOne(args[0], args[1]) {
			fmt.Println("insert success")
		} else {
			fmt.Println("insert failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	rootCmd.MarkFlagRequired("key")
	rootCmd.MarkFlagRequired("value")
}
