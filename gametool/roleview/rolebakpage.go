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
						MinSize:  dcl.Size{Width: 100, Height: 0},
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
						MinSize:       dcl.Size{Width: 100, Height: 0},
						Font:          dcl.Font{Family: "微软雅黑", PointSize: 10},
						StretchFactor: 2,
					},
				},
			},
			dcl.VSplitter{
				MinSize: dcl.Size{Width: 550, Height: 0},
				Font:    dcl.Font{Family: "微软雅黑", PointSize: 10},
				Children: []dcl.Widget{
					dcl.Composite{
						Layout: dcl.Grid{Columns: 3},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 1,
								Text:       "Bak文件路径:",
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "解析",
							},
							dcl.TextEdit{
								ColumnSpan: 3,
								MinSize:    dcl.Size{Width: 100, Height: 20},
								Text:       dcl.Bind("Remarks"),
							},
						},
					},
					dcl.Composite{
						Layout: dcl.Grid{Columns: 6, Spacing: 10},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 6,
								Text:       "Bak头部信息:",
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名长度:",
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色数据长度:",
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名:",
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},

							dcl.Label{
								ColumnSpan: 1,
								Text:       "CRC1:",
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "CRC2:",
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.TextEdit{
								ColumnSpan: 6,
								MinSize:    dcl.Size{Width: 100, Height: 20},
							},
						},
					},
				},

				OnSizeChanged: func() {
					pg.webView.Refresh()
				},
			},
		},
	}

	return tab
}
