//go:build linux

package service

import (
	"github.com/sevlyar/go-daemon"
	"github.com/spf13/viper"
	"goclip/internal"
	"goclip/storage"
)

func StartService() {
	ReadConfig()
	//创建一个守护进程
	cntxt := &daemon.Context{
		PidFileName: viper.GetString("pid"), //存放进程ID的文件名
		PidFilePerm: 0644,                   //PID文件的权限(0644表示所有者可读写,其他用户只读)
		LogFileName: viper.GetString("log"), //守护进程的标准输出日志文件
		LogFilePerm: 0640,                   //
		WorkDir:     viper.GetString("dir"), //服务的工作目录
		Umask:       027,                    //设置文件创建时的权限掩码
	}

	//ToDo 这里的模型不适用于windows，记得以后更换
	d, err := cntxt.Reborn() //fork出一个子进程
	if err != nil {
		panic(err)
	}
	//如果d不是空的，那么就退出
	if d != nil {
		//父进程返回，子进程开始运行
		return
	}
	//释放资源
	defer cntxt.Release()

	//初始化数据库
	_, DBErr := storage.NewDB("./../clip_history.db")
	if DBErr != nil {
		panic(DBErr)
	}
	internal.WatchClipBoard()
}
