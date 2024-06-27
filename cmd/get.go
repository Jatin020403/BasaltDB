package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/Jatin020403/BasaltDB/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get from Database",
	Long:  `This command gets value from database.`,

	Run: func(cmd *cobra.Command, args []string) {

		var key string
		var err error

		key, err = cmd.Flags().GetString("key")
		if err != nil {
			fmt.Println(err.Error())
		}

		if key == "" || err != nil {
			if len(args) != 1 {
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

		hashedKey := utils.MurmurHashInt(key)

		res, err := database.GetOne(partition, hashedKey)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(res)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringP("key", "k", "", "Input key")
}
