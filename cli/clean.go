package cli

import (
	"github.com/spf13/cobra"
	"goclip/storage"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "clean the history",
	Run: func(cmd *cobra.Command, args []string) {
		storage.CleanData()
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
