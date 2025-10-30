package gui

import (
	"fyne.io/systray"
	"github.com/gen2brain/beeep"
)

func OnTray() {
	systray.Run(OnReady, nil)
}

func OnReady() {
	//icon := fyne.CurrentApp().Icon()
	//if icon != nil {
	//	//从这里提取二进制数据
	//	systray.SetIcon(icon.Content())
	//}

	systray.SetTitle("goclip")
	systray.SetTooltip("goclip")

	//获取ico数据
	icon := GetImageByte()

	systray.SetIcon(icon)

	hShow := systray.AddMenuItem("历史记录", "查看历史记录")
	qShow := systray.AddMenuItem("退出", "退出程序")
	//监听托盘事件
	go func() {
		for {
			select {
			case <-hShow.ClickedCh:
			case <-qShow.ClickedCh:
				systray.Quit()
				beeep.Notify("退出", "程序已经退出", "")
				CLIPAPP.Quit()
			}

		}
	}()
}
