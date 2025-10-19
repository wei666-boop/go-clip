package cli

import "github.com/spf13/cobra"

var visionCmd = &cobra.Command{
	Use:   "version",
	Short: "this command can show the vision of go-clip",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(visionCmd)
}
