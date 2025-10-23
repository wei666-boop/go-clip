package cli

import (
	"github.com/spf13/cobra"
	"goclip/internal"
	"goclip/pkg"
)

var exportCmd = &cobra.Command{
	Use:   "export [mode] [name]",
	Short: "export your history to a json_file or csv_file",
	Args:  cobra.ExactArgs(2), //有一个参数
	Run: func(cmd *cobra.Command, args []string) {
		var mode string = args[0]
		var name string = args[1]
		switch mode {
		case "json":
			internal.ExportJson(name)
		case "csv":
			internal.ExportCsv(name)
		default:
			pkg.Prompt(1, "please enter json or csv")
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
}
