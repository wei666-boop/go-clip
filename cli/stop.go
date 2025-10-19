package cli

import (
	"fmt"
	kservice "github.com/kardianos/service"
	"github.com/spf13/cobra"
	"goclip/service"
	"os"
	"runtime"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "this arg can stop go-clip server",
	Run: func(cmd *cobra.Command, args []string) {
		switch runtime.GOOS {
		case "windows":
			var s kservice.Service
			s = service.GetService()
			service.StopWService(s)
		case "linux":
			service.StopService()
		default:
			fmt.Fprintf(os.Stderr, "Unknown OS: %s\n", runtime.GOOS)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
