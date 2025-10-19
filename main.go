package main

import (
	"goclip/cli"
	"os"
)

func main() {
	//命令行执行入口
	if len(os.Args) > 1 {
		cli.CmdExecute()
	} else {

	}
}
