package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "BasaltDB",
	Short: " A Key-Value Database with persistant storage.",
	Long:  `BasaltDB is a SQLite and LevelDB inspired key-value database. It runs from a single executable with a CLI interface and persistant storage.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var key string
var value string

func init() {
	rootCmd.Flags().StringVarP(&key, "key", "k", "default", "key (required)")
	rootCmd.Flags().StringVarP(&value, "value", "v", "default", "value (required)")

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
