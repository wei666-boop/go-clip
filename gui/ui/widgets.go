package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GUIClip() {
	clipApp := app.New()
	w := clipApp.NewWindow("clip")
	//设置图标
	icon := fyne.NewStaticResource("clip.png", resourceClipPngData)
	clipApp.SetIcon(icon)

	//创建右侧页面组件
	welcomePage := widget.NewLabel("欢迎来到clip")
	historyPage := widget.NewLabel("历史记录显示")
	settingPage := widget.NewLabel("设置界面")
	startPage := widget.NewLabel("服务启动")
	stopPage := widget.NewLabel("服务停止")

	//右侧主内容
	rightContent := container.NewCenter(welcomePage)

	//控制历史记录页面的状态
	showingHistory := false

	//左侧菜单
	btnHistory := widget.NewButton("历史记录", func() {
		if showingHistory {
			rightContent.Objects = []fyne.CanvasObject{welcomePage}
			showingHistory = !showingHistory
		} else {
			rightContent.Objects = []fyne.CanvasObject{historyPage}
			showingHistory = !showingHistory
		}
	})
	btnSettings := widget.NewButton("设置", func() {
		rightContent.Objects = []fyne.CanvasObject{settingPage}
	})
	btnStart := widget.NewButton("开始", func() {
		rightContent.Objects = []fyne.CanvasObject{startPage}
	})
	btnStop := widget.NewButton("结束", func() {
		rightContent.Objects = []fyne.CanvasObject{stopPage}
	})

	leftMenu := container.NewBorder(
		container.NewVBox(
			btnHistory, btnSettings, btnStart, btnStop),
		nil, nil, nil)

	//主布局
	mainWidget := container.NewBorder(nil, nil, leftMenu, nil, rightContent)

	w.SetContent(mainWidget)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}
