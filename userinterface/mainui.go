package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/sxn0508/fyne_zh_CN"
)

func Mainui() {
	myApp := app.New()
	myWindow := myApp.NewWindow("按钮页面")

	// 创建按钮页面
	buttonsPage := createButtonsPage(myWindow)

	myWindow.SetContent(buttonsPage)
	myWindow.ShowAndRun()
}

func createButtonsPage(w fyne.Window) fyne.CanvasObject {
	// 创建一个网格布局
	grid := container.NewGridWithColumns(4)

	// 创建12个按钮
	for i := 0; i < 12; i++ {
		button := widget.NewButton("按钮", func() {
			// 按钮按下后跳转到另一个界面
			nextPage := createNextPage(w, i+1)
			w.SetContent(nextPage)
		})
		grid.AddObject(button)
	}

	return grid
}

func createNextPage(w fyne.Window, buttonNumber int) fyne.CanvasObject {
	// 创建一个垂直布局
	vbox := container.NewVBox()

	// 添加返回按钮
	backButton := widget.NewButton("返回", func() {
		// 返回按钮按下后返回到按钮页面
		buttonsPage := createButtonsPage(w)
		w.SetContent(buttonsPage)
	})
	vbox.Add(backButton)

	// 添加一行文字
	label := widget.NewLabel("这是第 " + fmt.Sprintf("%d", buttonNumber) + " 个按钮的界面")
	vbox.Add(label)

	return vbox
}
