// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/lxn/walk"

	dcl "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

// MyMainWindow :
type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string

	roleBakPage *RoleBakPage
	roleDbPage  *RoleDbPage
}

var mw = new(MyMainWindow)

func main() {
	rand.Seed(time.Now().UnixNano())
	dcl.MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw.roleBakPage = new(RoleBakPage)
	mw.roleDbPage = new(RoleDbPage)

	ret := mw.createWindow()
	if !ret {
		return
	}

	mw.setIcon("../../gameresource/img/sword1_classic.ico")
	mw.setNotifyIcon("../../gameresource/img/sword1_classic.ico")

	mw.Closing().Attach(mw.closeEventHandler)
	mw.Disposing().Attach(mw.disposingEventHandler)

	mw.Run()
}

func (mw *MyMainWindow) setIcon(path string) {
	icon, err := walk.NewIconFromFile(path)

	if err != nil {
		return
	}

	if mw != nil {
		mw.SetIcon(icon)
	}
}

func (mw *MyMainWindow) setNotifyIcon(path string) {
	icon, err := walk.NewIconFromFile(path)
	if err != nil {
		return
	}

	ni, err := walk.NewNotifyIcon()
	if err != nil {
		return
	}

	err = ni.SetIcon(icon)
	if err != nil {
		return
	}

	err = ni.SetToolTip("单击图标显示程序窗口，右键图标显示或关闭程序")
	if err != nil {
		return
	}
	ni.MouseDown().Attach(mw.notifyIconMouseDownHandler)

	// 结束选项
	exitAction := walk.NewAction()
	err = exitAction.SetText("退出(&x)")
	if err != nil {
		return
	}
	exitAction.Triggered().Attach(mw.notifyIconExitActionHandler)

	err = ni.ContextMenu().Actions().Add(exitAction)
	if err != nil {
		return
	}

	// 打开选项
	openAction := walk.NewAction()
	err = openAction.SetText("打开(&o)")
	if err != nil {
		return
	}
	openAction.Triggered().Attach(mw.notifyIconOpenActionHandler)

	err = ni.ContextMenu().Actions().Add(openAction)
	if err != nil {
		return
	}

	// 显示通知栏图标
	err = ni.SetVisible(true)
	if err != nil {
		return
	}

	err = ni.ShowInfo("角色存档信息查看工具", "点击通知栏图标可以重新显示程序窗口.")
	if err != nil {
		log.Fatal(err)
	}
}

func (mw *MyMainWindow) openFile() error {
	dlg := new(walk.FileDialog)

	dlg.FilePath = mw.prevFilePath
	dlg.Filter = "Image Files (*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff)|*.emf;*.bmp;*.exif;*.gif;*.jpeg;*.jpg;*.png;*.tiff"
	dlg.Title = "Select an Image"

	if ok, err := dlg.ShowOpen(mw); err != nil {
		return err
	} else if !ok {
		return nil
	}

	mw.prevFilePath = dlg.FilePath

	return nil
}

func (mw *MyMainWindow) openActionTriggered() {
	if err := mw.openFile(); err != nil {
		log.Print(err)
	}
}

func (mw *MyMainWindow) newActionTriggered() {
	walk.MsgBox(mw, "New", "Newing something up... or not.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) changeViewActionTriggered() {
	walk.MsgBox(mw, "Change View", "By now you may have guessed it. Nothing changed.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) showAboutBoxActionTriggered() {
	walk.MsgBox(mw, "About", "Walk Actions Example", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) specialActionTriggered() {
	walk.MsgBox(mw, "Special", "Nothing to see here.", walk.MsgBoxIconInformation)
}

func (mw *MyMainWindow) disposingEventHandler() {
	fmt.Println("Disposing Event Handler")
}

func (mw *MyMainWindow) closeEventHandler(canceled *bool, reason walk.CloseReason) {
	mw.Hide()
	*canceled = true
}

func (mw *MyMainWindow) notifyIconMouseDownHandler(x, y int, button walk.MouseButton) {
	if button != walk.LeftButton {
		return
	}

	mw.Show()
}

func (mw *MyMainWindow) notifyIconExitActionHandler() {
	walk.App().Exit(0)
}

func (mw *MyMainWindow) notifyIconOpenActionHandler() {
	mw.Show()
}

func (mw *MyMainWindow) createWindow() bool {
	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu

	_mw := &dcl.MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "角色存档信息查看工具",
		MenuItems: []dcl.MenuItem{
			dcl.Menu{
				// &符号：Alt快捷键，快捷键即&后面字符
				Text: "文件(&F)",
				Items: []dcl.MenuItem{
					dcl.Action{
						AssignTo: &openAction,
						Text:     "Open",
						Image:    "../../gameresource/img/open.png",
						// Shortcut字段：设置Action的快捷键
						Shortcut:    dcl.Shortcut{Modifiers: walk.ModControl, Key: walk.KeyO},
						OnTriggered: mw.openActionTriggered,
					},
					dcl.Menu{
						AssignTo: &recentMenu,
						Text:     "Recent",
					},
					dcl.Separator{},
					dcl.Action{
						Text:        "Exit",
						OnTriggered: func() { mw.Close() },
					},
				},
			},
			dcl.Menu{
				Text: "帮助(&H)",
				Items: []dcl.MenuItem{
					dcl.Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: mw.showAboutBoxActionTriggered,
					},
				},
			},
		},

		ToolBar: dcl.ToolBar{
			ButtonStyle: dcl.ToolBarButtonImageBeforeText,
			Items: []dcl.MenuItem{
				//dcl.ActionRef{Action: &openAction},
				dcl.Menu{
					Text:  "New A",
					Image: "../../gameresource/img/document-new.png",
					Items: []dcl.MenuItem{
						dcl.Action{
							Text:        "A",
							OnTriggered: mw.newActionTriggered,
						},
						dcl.Action{
							Text:        "B",
							OnTriggered: mw.newActionTriggered,
						},
						dcl.Action{
							Text:        "C",
							OnTriggered: mw.newActionTriggered,
						},
					},
					OnTriggered: mw.newActionTriggered,
				},
				dcl.Separator{},
				dcl.Menu{
					Text:  "View",
					Image: "../../gameresource/img/document-properties.png",
					Items: []dcl.MenuItem{
						dcl.Action{
							Text:        "X",
							OnTriggered: mw.changeViewActionTriggered,
						},
						dcl.Action{
							Text:        "Y",
							OnTriggered: mw.changeViewActionTriggered,
						},
						dcl.Action{
							Text:        "Z",
							OnTriggered: mw.changeViewActionTriggered,
						},
					},
				},
				dcl.Separator{},
				dcl.Action{
					Text:        "Special",
					Image:       "../../gameresource/img/system-shutdown.png",
					OnTriggered: mw.specialActionTriggered,
				},
			},
		},

		//ContextMenuItems: []dcl.MenuItem{
		//	dcl.ActionRef{Action: &showAboutBoxAction},
		//},
		MinSize: dcl.Size{Width: 1200, Height: 900},
		Layout:  dcl.VBox{},

		Children: []dcl.Widget{
			dcl.TabWidget{
				Pages: []dcl.TabPage{
					*mw.roleBakPage.Create(),
					*mw.roleDbPage.Create(),
				},
			},
		},
	}

	err := _mw.Create()
	if err != nil {
		return false
	}

	func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openActionTriggered)
			recentMenu.Actions().Add(a)
		}
	}("Foo", "Bar", "Baz")

	return true
}
