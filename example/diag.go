package main

import (
	_ "embed"
	"fmt"
	"image/color"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/flopp/go-findfont"
)

//go:embed 1.png
var icn []byte

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simfang.ttf") {
			fmt.Println(path)
			os.Setenv("FYNE_FONT", path) // 设置环境变量  // 取消环境变量 os.Unsetenv("FYNE_FONT")
			break
		}
	}
	fmt.Println("=============")
}

// On Ubuntu you need install some packages:
// apt-get install libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev

func initMainDialog(app fyne.App, ch chan bool) *MainDialog {
	newDialog := &MainDialog{app: app, window: app.NewWindow("来聊"), ch: ch}

	newDialog.Init()

	return newDialog
}

func ShowDialog(ch chan bool) {
	myApp := app.New()

	mainDialog := initMainDialog(myApp, ch)

	mainDialog.Show()

	myApp.Run()

	displayExited()
}

func displayExited() {
	fmt.Println("Exited")
}

type MainDialog struct {
	app    fyne.App
	window fyne.Window
	ch     chan bool
}

func (dialog *MainDialog) Init() {
	// icon1 := widget.NewIcon(fyne.NewStaticResource("icons", icn)) // 创建一个icon图标
	// icon2 := widget.NewFileIcon(nil) // 创建一个file图标可添加路径的
	//  icon2.SetURI(fyne.CurrentApp().Storage().RootURI()) // 设置图标路径
	//icon2.SetSelected(true) // 选中图标
	//lable := widget.NewLabel("") // 标签控件
	dialog.window.SetIcon(fyne.NewStaticResource("111", icn)) // 给主程序设置图标

	infoText := widget.NewLabel("已检测到新版本是否要更新?")

	centered := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), infoText, layout.NewSpacer())

	button1 := widget.NewButton("  是  ", func() {
		//.SetText("Excelent!")
		fmt.Println(true)
		dialog.ch <- true
		dialog.Close()
	})

	button2 := widget.NewButton("  否  ", func() {
		//infoText.SetText("You're wrong!")
		fmt.Println(false)
		dialog.ch <- false
		dialog.Close()
	})
	// 两边对齐
	hcontainer5 := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		red_button(button1),
		layout.NewSpacer(),
		button2)

	dialog.window.SetContent(container.NewVBox(centered, hcontainer5))
}

func (dialog *MainDialog) Show() {
	dialog.window.Show()
}
func (dialog *MainDialog) Close() {
	dialog.window.Close()
}


func red_button(btn *widget.Button) *fyne.Container { // return type
	// btn := widget.NewButton("", nil) // button widget
	// button color
	btn_color := canvas.NewRectangle(
		color.NRGBA{R: 64, G: 158, B: 255, A: 255},
	)
	//color.NRGB{}
	// container for colored button
	container1 := container.New(
		// layout of container
		layout.NewMaxLayout(),
		// first use btn color
		btn_color,
		// 2nd btn widget
		btn,
	)
	// our button is ready
	return container1
}
