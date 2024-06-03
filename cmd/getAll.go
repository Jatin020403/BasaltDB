package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/spf13/cobra"
)

// getAllcmd represents the insert command
var getAllcmd = &cobra.Command{
	Use:   "getAll",
	Short: "getAll for Database",
	Long:  `This command inserts into database.`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) != 0 {
			fmt.Println("Length of args must be 0")
			return
		}

		partition, err := rootCmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
		}

		database.GetAll(partition)

	},
}

func init() {
	rootCmd.AddCommand(getAllcmd)
}
