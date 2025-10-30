package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"goclip/internal"
	"goclip/model"
	"goclip/storage"
	"os"
)

// ToDo系统托盘和主应用的启动顺序问题
var CLIPAPP fyne.App

func GUIClip() {
	var stopFlag = make(chan bool, 1)
	clipApp := app.New() //创建一个新的应用实例
	CLIPAPP = clipApp
	w := clipApp.NewWindow("clip") //创建一个新窗口
	//设置图标
	icon := fyne.NewStaticResource("clip.png", resourceClipPngData)

	clipApp.SetIcon(icon)

	//创建右侧页面组件(标签控件)
	welcomePage := widget.NewLabel("欢迎来到clip")

	var historyList []model.History
	historyList = storage.GetData()

	fmt.Fprintf(os.Stdout, "historyList: %v\n", historyList)

	//调用createHistoryTable函数传入历史数据生成一个表格控件
	historyPage := createHistoryTable(historyList)
	settingPage := widget.NewLabel("设置界面")
	settingPage.Alignment = fyne.TextAlignCenter
	startPage := widget.NewLabel("服务启动")
	stopPage := widget.NewLabel("服务停止")

	//反馈界面
	emailEntry := widget.NewEntry() //输入框
	contentEntry := widget.NewEntry()
	nameEntry := widget.NewEntry()
	//提交表单
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "邮箱地址", Widget: emailEntry},
			{Text: "内容", Widget: contentEntry},
			{Text: "姓名", Widget: nameEntry},
		},
	}
	btnSend := widget.NewButton("发送", func() {
		var email string
		var content string
		var name string
		email = emailEntry.Text
		content = contentEntry.Text
		name = nameEntry.Text
		fmt.Println("email:", email)
		fmt.Println("content:", content)
		fmt.Println("name", name)
		//ToDO发送邮箱
	})
	feedbackPage := container.NewVBox(widget.NewLabel("反馈"), form, btnSend)

	//创建一个最大容器,这种容器会填满所有可用空间,适合做主显示区
	rightContent := container.NewMax(welcomePage)

	//控制历史记录页面的状态
	showingHistory := false

	//左侧菜单
	btnHistory := widget.NewButton("历史记录", func() {
		if showingHistory {
			//更新历史数据
			newList := storage.GetData()
			historyPage = createHistoryTable(newList)
			rightContent.Objects = []fyne.CanvasObject{historyPage}
			//刷新页面
			rightContent.Refresh()
			showingHistory = !showingHistory
		} else {
			rightContent.Objects = []fyne.CanvasObject{welcomePage}
			rightContent.Refresh()
			showingHistory = !showingHistory
		}
	})
	btnSettings := widget.NewButton("设置", func() {
		rightContent.Objects = []fyne.CanvasObject{settingPage}
	})
	btnStart := widget.NewButton("开始", func() {
		rightContent.Objects = []fyne.CanvasObject{startPage}
		//switch runtime.GOOS {
		//case "windows":
		//	var s kservice.Service
		//	s = service.GetService()
		//	service.StartWService(s)
		//case "linux":
		//	service.StartService()
		//default:
		//	pkg.WriteInfoLog("service.log", "Unsupported OS")
		//}

		go internal.WatchClipBoard(stopFlag)
		go internal.WatchClipBoard(stopFlag)
	})
	btnStop := widget.NewButton("结束", func() {
		rightContent.Objects = []fyne.CanvasObject{stopPage}
		//switch runtime.GOOS {
		//case "windows":
		//	var s kservice.Service
		//	s = service.GetService()
		//	service.StopWService(s)
		//case "linux":
		//	service.StopService()
		//default:
		//	pkg.WriteInfoLog("service.log", "Unsupported OS")
		//}
		stopFlag <- true
	})
	btnFeedBack := widget.NewButton("反馈", func() {
		rightContent.Objects = []fyne.CanvasObject{feedbackPage}
	})

	//创建左侧菜单区域
	//使用container.NewBorder(left, right, top, bottom, center)布局
	leftMenu := container.NewBorder(
		//设置垂直布局
		container.NewVBox(
			btnHistory, btnSettings, btnStart, btnStop, btnFeedBack),
		nil, nil, nil)

	//主布局
	mainWidget := container.NewBorder(nil, nil, leftMenu, rightContent, nil)

	//将主布局内容设置为窗口内容
	w.SetContent(mainWidget)
	//设置窗口大小
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

// 返回可显示的UI组件
func createHistoryTable(data []model.History) fyne.CanvasObject {
	//创建一个表格,表格行数=数据长度+1,列数=2
	table := widget.NewTable(func() (int, int) {
		return len(data) + 1, 2
	},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		}, //单元格工厂函数,每个单元格都是一个空标签
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			//if len(data) == 0 {
			//	return widget.NewLabel("暂无数据")
			//}
			label := cell.(*widget.Label)
			if id.Row == 0 {
				if id.Col == 0 {
					label.SetText("时间")
				} else {
					label.SetText("内容")
				}
			} else {
				h := data[id.Row-1]
				if id.Col == 0 {
					label.SetText(h.Time.String()) //将时间转换为字符串类型
				} else {
					label.SetText(h.Content)
				}
			}
		},
	)
	//返回一个MAX容器
	return container.NewMax(table)
}
