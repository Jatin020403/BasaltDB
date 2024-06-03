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

func init() {
	rootCmd.PersistentFlags().StringP("use", "p", "default", "Set partition")
}
