package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func GUIClip() {
	ClipApp := app.New()
	ClipApp.SetIcon(nil)
	ClipWindow := ClipApp.NewWindow("goclip")
	historyButton := widget.NewButton("历史", HistoryButton)

	startButton := widget.NewButton("启动服务", StartButton)

	stopButton := widget.NewButton("停止服务", StopButton)

	ClipWindow.SetContent(container.NewHBox(historyButton, startButton, stopButton))

	ClipWindow.ShowAndRun()
}

func main() {
	GUIClip()
}
