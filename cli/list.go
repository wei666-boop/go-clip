package cli

import (
	"github.com/spf13/cobra"
	"goclip/storage"
)

var listCmd = &cobra.Command{
	Use:   "list [limit]",
	Short: "execute command of list can show your history of clipboard",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit := args[0]
		storage.List(limit)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
