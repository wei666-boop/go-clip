package service

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func StopService() {
	ReadConfig()
	data, err := os.ReadFile(viper.GetString("pid") + viper.GetString("dir"))
	if err != nil {
		return
	}
	//读取进程号
	var pid int
	fmt.Sscanf(string(data), "%d", &pid)
	//找到该进程
	proc, err := os.FindProcess(pid)
	if err != nil {
		return
	}
	//杀死该进程
	err = proc.Kill()
	if err != nil {
		return
	}
	//移除该文件
	os.Remove("./../log/goclip.pid")

}
