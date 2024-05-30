package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete from Database",
	Long:  `This command deletes from database.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Length of args must be 1")
			return
		}
		err := database.DeleteNode(args[0])
		if err {
			fmt.Println("delete success")
		} else {
			fmt.Println("delete failed")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	rootCmd.MarkFlagRequired("key")
}
