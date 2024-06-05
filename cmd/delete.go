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

		var key string
		var err error

		key, err = cmd.Flags().GetString("key")
		if err != nil {
			fmt.Println(err.Error())
		}

		if key == "" {
			if len(args) != 1 || err != nil {
				fmt.Println("Invalid input format. \nPass key with key, -k or as single string")
				cmd.Help()
				return
			}
			key = args[0]
		}

		partition, err := rootCmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
		}

		if err = database.DeleteOne(partition, key); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("delete success")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("key", "k", "", "Input key")
}
