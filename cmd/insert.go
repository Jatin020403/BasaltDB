package cmd

import (
	"fmt"

	"github.com/Jatin020403/BasaltDB/database"
	"github.com/Jatin020403/BasaltDB/utils"
	"github.com/spf13/cobra"
)

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert into Database",
	Long:  `This command inserts into database.`,

	Run: func(cmd *cobra.Command, args []string) {

		var key, value string
		var err error

		key, err = cmd.Flags().GetString("key")
		if err != nil {
			fmt.Println(err.Error())
		}
		value, err = cmd.Flags().GetString("value")
		if err != nil {
			fmt.Println(err.Error())
		}

		if key == "" || value == "" {
			if len(args) != 2 || err != nil {
				fmt.Println("Invalid input format. \nPass input as key value pair, or with -k -v flags")
				cmd.Help()
				return
			}
			key = args[0]
			value = args[1]
		}

		var partition string

		partition, err = rootCmd.Flags().GetString("use")
		if err != nil {
			fmt.Println(err.Error())
		}

		hashedKey := utils.MurmurHashInt(key)

		if err = database.InsertOne(partition, hashedKey, value); err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("insert successful")
		}
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)
	insertCmd.Flags().StringP("key", "k", "", "Input key")
	insertCmd.Flags().StringP("value", "v", "", "Input value")
	insertCmd.MarkFlagsRequiredTogether("key", "value")
}
