package main

import (
	"fmt"
	kservice "github.com/kardianos/service"
	"goclip/service"
	"os"
	"runtime"
)

//根据操作系统来确定不同的行为

func StartButton() {
	switch runtime.GOOS {
	case "windows":
		var s kservice.Service
		s = service.GetService()
		service.StartWService(s)
	case "linux":
		service.StopService()
	default:
		fmt.Fprintf(os.Stderr, "Unknown OS: %s\n", runtime.GOOS)
	}
}

func StopButton() {
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
}

func HistoryButton() {

}
