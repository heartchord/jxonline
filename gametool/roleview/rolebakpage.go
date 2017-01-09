package main

import (
	"log"

	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// RoleBakPage : role bak analyse page
type RoleBakPage struct {
	parent    *MyMainWindow
	self      *walk.TabPage
	treeView  *walk.TreeView
	webView   *walk.WebView
	treeModel *DirectoryTreeModel
}

// Create creates a new RoleBakPage instance
func (pg *RoleBakPage) Create(parent *MyMainWindow) *dcl.TabPage {
	var err error
	pg.treeModel, err = NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}
	pg.parent = parent

	tab := &dcl.TabPage{
		AssignTo: &pg.self,
		Title:    "Role Bak",
		Layout:   dcl.HBox{},
		Children: []dcl.Widget{
			dcl.HSplitter{
				Children: []dcl.Widget{
					// 目录树控件
					dcl.TreeView{
						AssignTo: &pg.treeView,
						MinSize:  dcl.Size{Width: 100},
						Font:     dcl.Font{Family: "微软雅黑", PointSize: 10},
						Model:    pg.treeModel,
						OnCurrentItemChanged: func() {
							dir := pg.treeView.CurrentItem().(*DirectoryNode)
							path := dir.Path()
							pg.webView.SetURL(path)
						},
						OnSizeChanged: func() {
							pg.webView.Refresh()
						},
					},
					// 视图控件
					dcl.WebView{
						AssignTo:      &pg.webView,
						MinSize:       dcl.Size{Width: 100},
						Font:          dcl.Font{Family: "微软雅黑", PointSize: 10},
						StretchFactor: 2,
					},
					dcl.VSplitter{
						MinSize: dcl.Size{Width: 500},
						Children: []dcl.Widget{
							dcl.PushButton{
								Text: "test1",
							},
							dcl.PushButton{
								Text: "test2",
							},
						},
						OnSizeChanged: func() {
							pg.webView.Refresh()
						},
					},
				},
			},
		},
	}

	return tab
}
