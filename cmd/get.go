package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get from Database",
	Long:  `This command gets value from database.`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Length of args must be 1")
			return
		}
		res, err := database.GetOne(args[0])

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	rootCmd.MarkFlagRequired("key")
}
