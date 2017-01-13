package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
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
	bakDataBinder          *walk.DataBinder
	treeView               *walk.TreeView
	tableView              *walk.TableView
	roleBaseDataTV         *walk.TableView
	treeModel              *DirectoryTreeModel
	tableModel             *FileInfoModel
	roleBaseDataModel      *RoleBaseDataModel
	roleBaseDataComposite  *walk.Composite
	bakFilePathText        *walk.LineEdit
	bakFileRoleNameLenText *walk.LineEdit
	bakFileRoleDataLenText *walk.LineEdit
	bakFileRoleNameText    *walk.LineEdit
	bakFileProcessLogText  *walk.TextEdit
	bakFileCRC1            *walk.LineEdit
	bakFileCRC2            *walk.LineEdit
	decodeProcessFinished  bool
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
	pg.tableModel = NewFileInfoModel()
	pg.roleBaseDataModel = NewRoleBaseDataModel()
	pg.decodeProcessFinished = true

	var ep walk.ErrorPresenter
	tab := &dcl.TabPage{
		AssignTo: &pg.TabPage,
		Title:    "Role Bak",
		Layout:   dcl.HBox{},
		DataBinder: dcl.DataBinder{
			AssignTo:   &pg.bakDataBinder,
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
						OnCurrentItemChanged: pg.onCurrentTreeViewItemChanged,
						OnSizeChanged:        pg.onCurrentTreeViewSizeChanged,
					},
					dcl.TableView{
						AssignTo:              &pg.tableView,
						Font:                  dcl.Font{Family: "微软雅黑", PointSize: 10},
						Model:                 pg.tableModel,
						StretchFactor:         2,
						OnCurrentIndexChanged: pg.onCurrentTableViewItemChanged,
						OnItemActivated:       pg.onCurrentTableViewItemActivated,
						OnKeyDown:             pg.onCurrentTableViewKeyDown,
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
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.bakFilePathText,
								Text:       dcl.Bind("BakFilePath"),
								ColumnSpan: 1,
								ReadOnly:   true,
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "解析",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
								OnClicked:  pg.onDecodeROleBakData,
							},
							dcl.TextEdit{
								AssignTo:   &pg.bakFileProcessLogText,
								ColumnSpan: 3,
								MinSize:    dcl.Size{Width: 100, Height: 20},
								Text:       "",
								ReadOnly:   true,
								OnSizeChanged: func() {
								},
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						Font:   dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout: dcl.Grid{Columns: 6, Spacing: 10},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 6,
								Text:       "【Bak头部信息】",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名长度：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.bakFileRoleNameLenText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色数据长度：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.bakFileRoleDataLenText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.bakFileRoleNameText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "CRC1：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.bakFileCRC1,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "CRC2：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.bakFileCRC2,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						AssignTo: &pg.roleBaseDataComposite,
						MinSize:  dcl.Size{Width: 0, Height: 450},
						Font:     dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:   dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 1,
								Text:       "【角色基础数据信息】",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.TableView{
								AssignTo:         &pg.roleBaseDataTV,
								ColumnSpan:       1,
								CheckBoxes:       true,
								ColumnsOrderable: true,
								MultiSelection:   true,
								Columns: []dcl.TableViewColumn{
									{Title: "数据索引"},
									{Title: "数据名称"},
									{Title: "数据内容"},
									{Title: "数据说明"},
								},
								Model: pg.roleBaseDataModel,
								OnSelectedIndexesChanged: func() {
									fmt.Printf("SelectedIndexes: %v\n", pg.roleBaseDataTV.SelectedIndexes())
								},
								OnItemActivated: func() {
									idx := pg.roleBaseDataTV.CurrentIndex()
									pg.roleBaseDataModel.SwitchRowCheckedState(idx)
								},
								OnMouseDown: func(x, y int, button walk.MouseButton) {
									if button != walk.RightButton {
										return
									}

									idx := pg.roleBaseDataTV.CurrentIndex()
									fmt.Println(idx)

									// 打开选项
									//openAction := walk.NewAction()
									//err = openAction.SetText("打开(&o)")
									//if err != nil {
									//	return
									//}
									//openAction.Triggered().Attach(pg.notifyIconOpenActionHandler)
									//pg.roleBaseDataComposite.Layout().Container().ContextMenu().Actions().Add(openAction)
									//if err != nil {
									//	return
									//}
								},
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						Font:   dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
						Layout: dcl.Grid{Columns: 5, Spacing: 1},
						Children: []dcl.Widget{
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "角色技能",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								//OnClicked:  pg.onDecodeROleBakData,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "任务变量",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								//OnClicked:  pg.onDecodeROleBakData,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "角色物品",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								//OnClicked:  pg.onDecodeROleBakData,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "角色状态",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								//OnClicked:  pg.onDecodeROleBakData,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "扩展数据",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								//OnClicked:  pg.onDecodeROleBakData,
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
	dir := pg.treeView.CurrentItem().(*DirectoryNode)
	err := pg.tableModel.SetDirPath(dir.Path())
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
	if index := pg.tableView.CurrentIndex(); index > -1 {
		name := pg.tableModel.items[index].Name
		dir := pg.treeView.CurrentItem().(*DirectoryNode)
		url = filepath.Join(dir.Path(), name)
		pg.bakFilePathText.SetText(url)
	}
}

func (pg *RoleBakPage) onCurrentTableViewItemActivated() {

	tlvIndex := pg.tableView.CurrentIndex()
	if tlvIndex <= -1 {
		return
	}

	curItem := pg.treeView.CurrentItem()
	curNode := curItem.(*DirectoryNode)

	name := pg.tableModel.items[tlvIndex].Name
	trvIndex := curNode.FindChild(name)
	if trvIndex <= -1 {
		return
	}

	path := filepath.Join(curNode.Path(), name)
	if !goblazer.IsFileDirectory(path) { // 如果不是目录，返回
		return
	}

	// 更新目录树
	pg.treeView.SetExpanded(curItem, true)
	child := curNode.ChildAt(trvIndex)
	pg.treeView.SetCurrentItem(child)

	err := pg.tableModel.SetDirPath(path)
	if err != nil {
		walk.MsgBox(mw, "Error", err.Error(),
			walk.MsgBoxOK|walk.MsgBoxIconError)
	}
}

func (pg *RoleBakPage) onCurrentTableViewKeyDown(key walk.Key) {

	switch key {
	case walk.KeyBack:
		{
			curItem := pg.treeView.CurrentItem()
			if curItem == nil {
				return
			}

			parentItem := curItem.Parent()
			if parentItem == nil {
				return
			}
			parentNode := parentItem.(*DirectoryNode)

			// 更新目录树
			pg.treeView.SetExpanded(parentItem, true)
			pg.treeView.SetExpanded(curItem, false)
			pg.treeView.SetCurrentItem(parentItem)

			err := pg.tableModel.SetDirPath(parentNode.Path())
			if err != nil {
				walk.MsgBox(mw, "Error", err.Error(),
					walk.MsgBoxOK|walk.MsgBoxIconError)
			}
		}
	}
}

func (pg *RoleBakPage) onDecodeROleBakData() {
	if !pg.decodeProcessFinished {
		pg.WriteLog("Info - Last decode process hasn't finished, please wait...")
		return
	}

	pg.decodeProcessFinished = false
	pg.bakFileProcessLogText.SetText("")
	go pg.BakDecodeRoutineFunction(bakBindData.BakFilePath)
}

// WriteLog :
func (pg *RoleBakPage) WriteLog(format string, a ...interface{}) (n int, err error) {
	ts := time.Now().Unix()
	tm := time.Unix(ts, 0)
	t := tm.Format("2006-01-02 15:04:05")

	format = t + " : " + format + "\r\n"
	log := fmt.Sprintf(format, a...)
	pg.bakFileProcessLogText.AppendText(log)
	return pg.bakFileProcessLogText.TextLength(), nil
}

// BakDecodeRoutineFunction :
func (pg *RoleBakPage) BakDecodeRoutineFunction(filePath string) {
	defer func() {
		pg.decodeProcessFinished = true
	}()

	// 文件存在性判断
	if !goblazer.IsFileExisted(filePath) {
		pg.WriteLog("Error - File not existed, the input path is [%s]!", filePath)
		return
	}

	// 文件后缀名判断
	fileName := path.Base(filePath)
	fileSuffix := path.Ext(fileName)
	if fileSuffix != ".bak" {
		pg.WriteLog("Error - Not role bak file!")
		return
	}

	fi, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	data, err := ioutil.ReadAll(fi)
	encoder := gameencoder.NewRoleBakEncoder()
	encoder.SetLogger(pg.WriteLog)
	encoder.Decode(data)

	mdecoder := mahonia.NewDecoder("GBK")

	roleName := string(encoder.BakData.RoleNameGBK)
	roleName = mdecoder.ConvertString(roleName)
	roleNameLen := fmt.Sprintf("%d", encoder.BakData.RoleNameLen)
	roleDataLen := fmt.Sprintf("%d", encoder.BakData.RoleDataLen)

	pg.bakFileRoleNameText.SetText(roleName)
	pg.bakFileRoleNameLenText.SetText(roleNameLen)
	pg.bakFileRoleDataLenText.SetText(roleDataLen)

	crc1 := fmt.Sprintf("%X", encoder.CRC32Cal)
	pg.bakFileCRC1.SetText(crc1)

	crc2 := fmt.Sprintf("%X", encoder.CRC32Read)
	pg.bakFileCRC2.SetText(crc2)

	pg.roleBaseDataModel.ResetRows(&encoder.RoleBaseData)

	//time.Sleep(time.Second * 3)
}

func (pg *RoleBakPage) notifyIconOpenActionHandler() {
	mw.Show()
}
