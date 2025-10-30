package main

import (
	"github.com/spf13/viper"
	"goclip/conf"
	"goclip/pkg"
	"runtime"
)

//根据操作系统来读取配置文件

func ReadConfig() {
	switch runtime.GOOS {
	case "windows":
		viper.SetConfigName("./conf/serviceWindows.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			pkg.WriteErrorLog("conf.log", "读取配置文件发生错误")
			return
		}
	case "linux":
		viper.SetConfigName("./conf/serviceLinux.yaml")
		err := viper.ReadInConfig()
		if err != nil {
			pkg.WriteErrorLog("conf.log", "读取配置文件出错")
			return
		}
	default:
		pkg.Prompt(conf.Mode, "unknow OS,please use windows or linux")
	}

}
