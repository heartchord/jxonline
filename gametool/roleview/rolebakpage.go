package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
	"github.com/lxn/win"
)

// RoleBakPage : role bak analyse page
type RoleBakPage struct {
	*walk.TabPage
	parent *MyMainWindow

	treeView  *walk.TreeView
	webView   *walk.WebView
	treeModel *DirectoryTreeModel
}

type MyWebView struct {
	*walk.WebView
}

func NewMyWebView(wv *walk.WebView) (*MyWebView, error) {
	mwv := &MyWebView{wv}

	if err := walk.InitWrapperWindow(mwv); err != nil {
		return nil, err
	}

	return mwv, nil
}

func (wv *MyWebView) WndProc(hwnd win.HWND, msg uint32, wParam, lParam uintptr) uintptr {
	fmt.Println("here")
	switch msg {
	case win.WM_SIZE, win.WM_SIZING:

	}

	return wv.WidgetBase.WndProc(hwnd, msg, wParam, lParam)
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
		AssignTo: &pg.TabPage,
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
						OnMouseDown: func(x int, y int, button walk.MouseButton) {
							fmt.Printf("Mouse Down, x = %d, y = %d, button = %d\n", x, y, button)
						},
						OnMouseMove: func(x int, y int, button walk.MouseButton) {
							fmt.Printf("Mouse Down, x = %d, y = %d, button = %d\n", x, y, button)
						},
					},
					// 视图控件
					dcl.WebView{
						AssignTo:      &pg.webView,
						MinSize:       dcl.Size{Width: 100, Height: 0},
						Font:          dcl.Font{Family: "微软雅黑", PointSize: 10},
						StretchFactor: 2,
						OnSizeChanged: func() {
							fmt.Printf("OnSizeChanged")
						},
						OnMouseDown: func(x int, y int, button walk.MouseButton) {
							fmt.Printf("Mouse Down, x = %d, y = %d, button = %d\n", x, y, button)
						},
						OnMouseMove: func(x int, y int, button walk.MouseButton) {
							fmt.Printf("Mouse Down, x = %d, y = %d, button = %d\n", x, y, button)
						},
					},
				},
			},
			dcl.Composite{ // 这里重新布局
				MinSize: dcl.Size{Width: 600, Height: 0},
				Font:    dcl.Font{Family: "微软雅黑", PointSize: 10},
				Layout:  dcl.Grid{Columns: 1},
				Children: []dcl.Widget{

					dcl.Composite{ // 这里重新布局
						Layout: dcl.Grid{Columns: 3},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 1,
								Text:       "Bak文件路径：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 12, Bold: true},
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 12},
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "解析",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
								OnClicked: func() {
									fmt.Printf("OnCilcked\n")
								},
							},
							dcl.TextEdit{
								ColumnSpan: 3,
								MinSize:    dcl.Size{Width: 100, Height: 20},
								Text:       dcl.Bind("Remarks"),
							},
						},
					},

					dcl.VSpacer{Size: 5},

					dcl.Composite{ // 这里重新布局
						MinSize: dcl.Size{Width: 600, Height: 500},
						Font:    dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:  dcl.Grid{Columns: 6, Spacing: 10},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 6,
								Text:       "Bak头部信息:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 12, Bold: true},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名长度:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色数据长度:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
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
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								ColumnSpan: 1,
								Text:       "asd",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "CRC2:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
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
			},
		},
	}

	return tab
}

func (pg *RoleBakPage) webViewItemClicked(x int, y int, button walk.MouseButton) {
	fmt.Printf("Mouse Down, x = %d, y = %d, button = %d\n", x, y, button)
}
