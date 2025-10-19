package internal

import (
	"fmt"
	"github.com/atotto/clipboard"
	"goclip/storage"
	"os"
	"time"
)

func WatchClipBoard() {
	var last string
	fmt.Fprintln(os.Stdout, "服务启动")
	//开始监听剪切板变化
	for {
		text, err := clipboard.ReadAll()
		if err != nil {
			continue
		}
		if text != last {
			last = text
			fmt.Println(last)
			storage.AddDate(last)
		}
		//0.5秒轮训一次
		time.Sleep(time.Millisecond * 500)
	}
}
