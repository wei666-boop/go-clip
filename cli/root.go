package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"goclip/pkg"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "clip",
	Short: "clip is a command which get help for how to use go-clip",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stdout, "welcome to go-clip")
		pkg.WriteInfoLog("service.log", "go-clip service ")
	},
}

func CmdExecute() {
	rootCmd.Execute()
}
