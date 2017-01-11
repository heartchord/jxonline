// Copyright 2013 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"log"

	"github.com/lxn/walk"

	dcl "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

// MyMainWindow :
type MyMainWindow struct {
	*walk.MainWindow
	prevFilePath string
}

var mw = new(MyMainWindow)
var roleBakPage = new(RoleBakPage)
var roleDbPage = new(RoleDbPage)
var bakBindData = new(BakFileInfoBindData)

func main() {
	dcl.MustRegisterCondition("isSpecialMode", isSpecialMode)

	var openAction, showAboutBoxAction *walk.Action
	var recentMenu *walk.Menu

	if err := (dcl.MainWindow{
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

		ContextMenuItems: []dcl.MenuItem{
			dcl.ActionRef{Action: &showAboutBoxAction},
		},
		MinSize: dcl.Size{Width: 1200, Height: 900},
		Layout:  dcl.VBox{},

		Children: []dcl.Widget{
			dcl.TabWidget{
				Pages: []dcl.TabPage{
					*roleBakPage.Create(),
					*roleDbPage.Create(),
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	addRecentFileActions := func(texts ...string) {
		for _, text := range texts {
			a := walk.NewAction()
			a.SetText(text)
			a.Triggered().Attach(mw.openActionTriggered)
			recentMenu.Actions().Add(a)
		}
	}

	addRecentFileActions("Foo", "Bar", "Baz")

	mw.Run()
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
