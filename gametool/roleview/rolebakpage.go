package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/heartchord/goblazer"
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// BakFileInfoBindData :
type BakFileInfoBindData struct {
	BakURL string
}

// RoleBakPage : role bak analyse page
type RoleBakPage struct {
	*walk.TabPage
	bakFilePathText *walk.LineEdit
	db              *walk.DataBinder
	treeView        *walk.TreeView
	tableView       *walk.TableView
	treeModel       *DirectoryTreeModel
	tableModel      *FileInfoModel
}

// Create creates a new RoleBakPage instance
func (pg *RoleBakPage) Create() *dcl.TabPage {
	// create DirectoryTreeModel
	var err error
	pg.treeModel, err = NewDirectoryTreeModel()
	if err != nil {
		log.Fatal(err)
	}

	// create FileInfoModel
	bp.tableModel = NewFileInfoModel()

	var ep walk.ErrorPresenter
	tab := &dcl.TabPage{
		AssignTo: &pg.TabPage,
		Title:    "Role Bak",
		Layout:   dcl.HBox{},
		DataBinder: dcl.DataBinder{
			AssignTo:       &bp.db,
			DataSource:     bakdb,
			ErrorPresenter: dcl.ErrorPresenterRef{ErrorPresenter: &ep},
		},
		Children: []dcl.Widget{
			dcl.HSplitter{
				Children: []dcl.Widget{
					// 目录树控件
					dcl.TreeView{
						AssignTo:             &pg.treeView,
						MinSize:              dcl.Size{Width: 100, Height: 0},
						Font:                 dcl.Font{Family: "微软雅黑", PointSize: 10},
						Model:                pg.treeModel,
						OnCurrentItemChanged: bp.onCurrentTreeViewItemChanged,
						OnSizeChanged:        bp.onCurrentTreeViewSizeChanged,
					},
					dcl.TableView{
						AssignTo:              &bp.tableView,
						Font:                  dcl.Font{Family: "微软雅黑", PointSize: 10},
						Model:                 bp.tableModel,
						StretchFactor:         2,
						OnCurrentIndexChanged: bp.onCurrentTableViewItemChanged,
						OnItemActivated:       bp.onCurrentTableViewItemActivated,
						OnKeyDown:             bp.onCurrentTableViewKeyDown,
						Columns: []dcl.TableViewColumn{
							dcl.TableViewColumn{
								DataMember: "Name",
								Width:      120,
							},
							dcl.TableViewColumn{
								DataMember: "Size",
								Format:     "%d",
								Alignment:  dcl.AlignFar,
								Width:      64,
							},
							dcl.TableViewColumn{
								DataMember: "Modified",
								Format:     "2006-01-02 15:04:05",
								Width:      120,
							},
						},
					},
				},
			},
			dcl.Composite{ // 这里重新布局
				MinSize: dcl.Size{Width: 450, Height: 0},
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
								AssignTo:   &bp.bakFilePathText,
								ColumnSpan: 1,
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
								Text:       "Remarks",
							},
						},
					},

					dcl.VSpacer{Size: 5},

					dcl.Composite{ // 这里重新布局
						MinSize: dcl.Size{Width: 0, Height: 500},
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

// Tree View Event Handler
func (pg *RoleBakPage) onCurrentTreeViewItemChanged() {
	dir := bp.treeView.CurrentItem().(*DirectoryNode)
	err := bp.tableModel.SetDirPath(dir.Path())
	if err != nil {
		walk.MsgBox(mw, "Error", err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)
	}
}

func (pg *RoleBakPage) onCurrentTreeViewSizeChanged() {
}

// Table View Event Handler
func (pg *RoleBakPage) onCurrentTableViewItemChanged() {
	var url string
	if index := bp.tableView.CurrentIndex(); index > -1 {
		name := bp.tableModel.items[index].Name
		dir := bp.treeView.CurrentItem().(*DirectoryNode)
		url = filepath.Join(dir.Path(), name)
	}

	bp.bakFilePathText.SetText(url)
}

func (pg *RoleBakPage) onCurrentTableViewItemActivated() {

	tlvIndex := bp.tableView.CurrentIndex()
	if tlvIndex <= -1 {
		return
	}

	curItem := bp.treeView.CurrentItem()
	curNode := curItem.(*DirectoryNode)

	name := bp.tableModel.items[tlvIndex].Name
	trvIndex := curNode.FindChild(name)
	if trvIndex <= -1 {
		return
	}

	path := filepath.Join(curNode.Path(), name)
	if !goblazer.IsFileDirectory(path) { // 如果不是目录，返回
		return
	}

	// 更新目录树
	bp.treeView.SetExpanded(curItem, true)
	child := curNode.ChildAt(trvIndex)
	bp.treeView.SetCurrentItem(child)

	err := bp.tableModel.SetDirPath(path)
	if err != nil {
		walk.MsgBox(mw, "Error", err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)
	}
}

func (pg *RoleBakPage) onCurrentTableViewKeyDown(key walk.Key) {

	switch key {
	case walk.KeyBack:
		{
			curItem := bp.treeView.CurrentItem()
			if curItem == nil {
				return
			}

			parentItem := curItem.Parent()
			if parentItem == nil {
				return
			}
			parentNode := parentItem.(*DirectoryNode)

			// 更新目录树
			bp.treeView.SetExpanded(parentItem, true)
			bp.treeView.SetExpanded(curItem, false)
			bp.treeView.SetCurrentItem(parentItem)

			err := bp.tableModel.SetDirPath(parentNode.Path())
			if err != nil {
				walk.MsgBox(mw, "Error", err.Error(),
					walk.MsgBoxOK|walk.MsgBoxIconError)
			}
		}
	}
}
