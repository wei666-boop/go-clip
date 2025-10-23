package cli

import (
	"github.com/spf13/cobra"
	"goclip/pkg"
)

var rootCmd = &cobra.Command{
	Use:   "clip",
	Short: "clip is a command which get help for how to use go-clip",
	Run: func(cmd *cobra.Command, args []string) {
		pkg.Prompt(1, "欢迎来到clip")
		pkg.WriteInfoLog("service.log", "go-clip service ")
	},
}

func CmdExecute() {
	rootCmd.Execute()
}
