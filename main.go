package main

import (
	"goclip/cli"
	"goclip/conf"
	GUI "goclip/gui/ui"
	"os"
)

func main() {
	//命令行执行入口
	if len(os.Args) > 1 {
		conf.Mode = conf.CLIMODE
		cli.CmdExecute()
	} else {
		conf.Mode = conf.GUIMODE
		GUI.GUIClip()
	}
}
