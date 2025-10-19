//go:build windows

package service

import (
	"github.com/kardianos/service"
	"goclip/internal"
)

// 使用program来实现service接口
type program struct{}

//服务启动要干的事情

var stopFlag = make(chan bool)

func (p *program) Start(s service.Service) error {
	//初始化停止标志
	stopFlag <- false
	//go internal.WatchClipBoard()
	go internal.CheckHistoryCount()
	for {
		select {
		case <-stopFlag:
			return nil
		default:
			internal.WatchClipBoard()
		}
	}
}

//服务停止要干的事情

//ToDo 如何才能停止这一个进程？

func (p *program) Stop(s service.Service) error {
	//ToDo 如何停止该服务
	stopFlag <- true
	return nil
}

//ToDo 加上去重设置,确保在没有意外出现的情况下,只会创建一次服务

func GetService() service.Service {
	serviceConfig := &service.Config{
		Name:        "goclip",
		DisplayName: "goclip",
		Description: "A background service for clipboard synchronization",
	}

	prg := &program{}

	//创建一个新服务
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		return nil
	}
	return s
}

func StartWService(s service.Service) {
	s.Start()
}

func StopWService(s service.Service) {
	s.Stop()
}
