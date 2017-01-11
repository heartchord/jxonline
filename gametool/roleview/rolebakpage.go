package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/heartchord/goblazer"
	"github.com/heartchord/jxonline/gameencoder"
	"github.com/henrylee2cn/mahonia"
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// BakFileInfoBindData :
type BakFileInfoBindData struct {
	BakFilePath string
}

// RoleBakPage : role bak analyse page
type RoleBakPage struct {
	*walk.TabPage
	bakDataBinder *walk.DataBinder
	treeView      *walk.TreeView
	tableView     *walk.TableView
	treeModel     *DirectoryTreeModel
	tableModel    *FileInfoModel

	bakFilePathText        *walk.LineEdit
	bakFileRoleNameLenText *walk.LineEdit
	bakFileRoleDataLenText *walk.LineEdit
	bakFileRoleNameText    *walk.LineEdit
	bakFileProcessLogText  *walk.TextEdit
	bakFileCRC1            *walk.LineEdit
	bakFileCRC2            *walk.LineEdit
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
	roleBakPage.tableModel = NewFileInfoModel()

	var ep walk.ErrorPresenter
	tab := &dcl.TabPage{
		AssignTo: &pg.TabPage,
		Title:    "Role Bak",
		Layout:   dcl.HBox{},
		DataBinder: dcl.DataBinder{
			AssignTo:   &roleBakPage.bakDataBinder,
			DataSource: bakBindData,
			AutoSubmit: true,
			//OnSubmitted: func() {
			//	fmt.Println("OnSubmitted")
			//},
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
						OnCurrentItemChanged: roleBakPage.onCurrentTreeViewItemChanged,
						OnSizeChanged:        roleBakPage.onCurrentTreeViewSizeChanged,
					},
					dcl.TableView{
						AssignTo:              &roleBakPage.tableView,
						Font:                  dcl.Font{Family: "微软雅黑", PointSize: 10},
						Model:                 roleBakPage.tableModel,
						StretchFactor:         2,
						OnCurrentIndexChanged: roleBakPage.onCurrentTableViewItemChanged,
						OnItemActivated:       roleBakPage.onCurrentTableViewItemActivated,
						OnKeyDown:             roleBakPage.onCurrentTableViewKeyDown,
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
								AssignTo:   &roleBakPage.bakFilePathText,
								Text:       dcl.Bind("BakFilePath"),
								ColumnSpan: 1,
								ReadOnly:   true,
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 12},
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "解析",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
								OnClicked:  roleBakPage.onDecodeROleBakData,
							},
							dcl.TextEdit{
								AssignTo:   &roleBakPage.bakFileProcessLogText,
								ColumnSpan: 3,
								MinSize:    dcl.Size{Width: 100, Height: 20},
								Text:       "",
								ReadOnly:   true,
								OnSizeChanged: func() {
								},
							},
						},
					},

					dcl.VSpacer{Size: 1},

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
								AssignTo:   &roleBakPage.bakFileRoleNameLenText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色数据长度:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &roleBakPage.bakFileRoleDataLenText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &roleBakPage.bakFileRoleNameText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},

							dcl.Label{
								ColumnSpan: 1,
								Text:       "CRC1:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &roleBakPage.bakFileCRC1,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "CRC2:",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &roleBakPage.bakFileCRC2,
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
	dir := roleBakPage.treeView.CurrentItem().(*DirectoryNode)
	err := roleBakPage.tableModel.SetDirPath(dir.Path())
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
	if index := roleBakPage.tableView.CurrentIndex(); index > -1 {
		name := roleBakPage.tableModel.items[index].Name
		dir := roleBakPage.treeView.CurrentItem().(*DirectoryNode)
		url = filepath.Join(dir.Path(), name)
		roleBakPage.bakFilePathText.SetText(url)
	}
}

func (pg *RoleBakPage) onCurrentTableViewItemActivated() {

	tlvIndex := roleBakPage.tableView.CurrentIndex()
	if tlvIndex <= -1 {
		return
	}

	curItem := roleBakPage.treeView.CurrentItem()
	curNode := curItem.(*DirectoryNode)

	name := roleBakPage.tableModel.items[tlvIndex].Name
	trvIndex := curNode.FindChild(name)
	if trvIndex <= -1 {
		return
	}

	path := filepath.Join(curNode.Path(), name)
	if !goblazer.IsFileDirectory(path) { // 如果不是目录，返回
		return
	}

	// 更新目录树
	roleBakPage.treeView.SetExpanded(curItem, true)
	child := curNode.ChildAt(trvIndex)
	roleBakPage.treeView.SetCurrentItem(child)

	err := roleBakPage.tableModel.SetDirPath(path)
	if err != nil {
		walk.MsgBox(mw, "Error", err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)
	}
}

func (pg *RoleBakPage) onCurrentTableViewKeyDown(key walk.Key) {

	switch key {
	case walk.KeyBack:
		{
			curItem := roleBakPage.treeView.CurrentItem()
			if curItem == nil {
				return
			}

			parentItem := curItem.Parent()
			if parentItem == nil {
				return
			}
			parentNode := parentItem.(*DirectoryNode)

			// 更新目录树
			roleBakPage.treeView.SetExpanded(parentItem, true)
			roleBakPage.treeView.SetExpanded(curItem, false)
			roleBakPage.treeView.SetCurrentItem(parentItem)

			err := roleBakPage.tableModel.SetDirPath(parentNode.Path())
			if err != nil {
				walk.MsgBox(mw, "Error", err.Error(),
					walk.MsgBoxOK|walk.MsgBoxIconError)
			}
		}
	}
}

func (pg *RoleBakPage) onDecodeROleBakData() {
	roleBakPage.bakFileProcessLogText.SetText("")
	go BakDecodeRoutineFunction(bakBindData.BakFilePath)
}

// WriteLog :
func (pg *RoleBakPage) WriteLog(format string, a ...interface{}) (n int, err error) {
	ts := time.Now().Unix()
	tm := time.Unix(ts, 0)
	t := tm.Format("2006-01-02 15:04:05")

	format = t + " : " + format + "\r\n"
	log := fmt.Sprintf(format, a...)
	roleBakPage.bakFileProcessLogText.AppendText(log)
	return roleBakPage.bakFileProcessLogText.TextLength(), nil
}

// BakDecodeRoutineFunction :
func BakDecodeRoutineFunction(path string) {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	data, err := ioutil.ReadAll(fi)
	encoder := gameencoder.NewRoleBakEncoder()
	encoder.SetLogger(roleBakPage.WriteLog)
	encoder.Decode(data)

	mdecoder := mahonia.NewDecoder("GBK")

	roleName := string(encoder.BakData.RoleNameGBK)
	roleName = mdecoder.ConvertString(roleName)
	roleNameLen := fmt.Sprintf("%d", encoder.BakData.RoleNameLen)
	roleDataLen := fmt.Sprintf("%d", encoder.BakData.RoleDataLen)

	roleBakPage.bakFileRoleNameText.SetText(roleName)
	roleBakPage.bakFileRoleNameLenText.SetText(roleNameLen)
	roleBakPage.bakFileRoleDataLenText.SetText(roleDataLen)

	crc1 := fmt.Sprintf("%X", encoder.CRC32Cal)
	roleBakPage.bakFileCRC1.SetText(crc1)

	crc2 := fmt.Sprintf("%X", encoder.CRC32Read)
	roleBakPage.bakFileCRC2.SetText(crc2)
}
