package main

import (
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// RoleExtDataDialog :
type RoleExtDataDialog struct {
	*walk.Dialog

	LockSoulDataModel     *DataModel1
	LingLongLockDataModel *DataModel1
	TransNimbusDataModel  *DataModel1
	RoleBreakDataModel    *DataModel1
	EquipComposeDataModel *DataModel1
	LockSoulDataTV        *walk.TableView
	LingLongLockDataTV    *walk.TableView
	TransNimbusDataTV     *walk.TableView
	RoleBreakDataTV       *walk.TableView
	EquipComposeDataTV    *walk.TableView
}

// CreateInstance :
func (dlg *RoleExtDataDialog) CreateInstance(parent walk.Form) bool {
	dlg.LockSoulDataModel = NewDataModel1("../../gameresource/img/right-arrow2.ico")
	dlg.LingLongLockDataModel = NewDataModel1("../../gameresource/img/right-arrow2.ico")
	dlg.TransNimbusDataModel = NewDataModel1("../../gameresource/img/right-arrow2.ico")
	dlg.RoleBreakDataModel = NewDataModel1("../../gameresource/img/right-arrow2.ico")
	dlg.EquipComposeDataModel = NewDataModel1("../../gameresource/img/right-arrow2.ico")

	return dlg.CreateRoleExtDataDialog(parent)
}

// CreateRoleExtDataDialog :
func (dlg *RoleExtDataDialog) CreateRoleExtDataDialog(parent walk.Form) bool {
	var acceptPB, cancelPB *walk.PushButton

	o := &dcl.Dialog{
		AssignTo:      &dlg.Dialog,
		Title:         "角色扩展数据信息",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       dcl.Size{Width: 1000, Height: 750},
		Font:          dcl.Font{Family: "微软雅黑", PointSize: 10},
		Layout:        dcl.VBox{},
		Children: []dcl.Widget{
			dcl.Composite{ // 这里重新布局
				MinSize: dcl.Size{Width: 0, Height: 400},
				Font:    dcl.Font{Family: "微软雅黑", PointSize: 10},
				Layout:  dcl.Grid{Columns: 2, Spacing: 10},
				Children: []dcl.Widget{
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【角色锁魂数据信息】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
					},
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【角色玲珑锁数据信息】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
					},
					dcl.Composite{ // 这里重新布局
						ColumnSpan: 1,
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:     dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.TableView{
								AssignTo:         &dlg.LockSoulDataTV,
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
								Model: dlg.LockSoulDataModel,
								OnItemActivated: func() {
									//	idx := pg.roleBaseDataTV.CurrentIndex()
									//	pg.roleBaseDataModel.SwitchRowCheckedState(idx)
								},
							},
						},
					},
					dcl.Composite{ // 这里重新布局
						ColumnSpan: 1,
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:     dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.TableView{
								AssignTo:         &dlg.LingLongLockDataTV,
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
								Model: dlg.LingLongLockDataModel,
								OnItemActivated: func() {
									//	idx := pg.roleBaseDataTV.CurrentIndex()
									//	pg.roleBaseDataModel.SwitchRowCheckedState(idx)
								},
							},
						},
					},
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【角色转灵数据信息】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
					},
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【角色突破数据信息】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
					},
					dcl.Composite{ // 这里重新布局
						ColumnSpan: 1,
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:     dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.TableView{
								AssignTo:         &dlg.TransNimbusDataTV,
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
								Model: dlg.TransNimbusDataModel,
								OnItemActivated: func() {
									//	idx := pg.roleBaseDataTV.CurrentIndex()
									//	pg.roleBaseDataModel.SwitchRowCheckedState(idx)
								},
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						ColumnSpan: 1,
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:     dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.TableView{
								AssignTo:         &dlg.RoleBreakDataTV,
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
								Model: dlg.RoleBreakDataModel,
								OnItemActivated: func() {
									//	idx := pg.roleBaseDataTV.CurrentIndex()
									//	pg.roleBaseDataModel.SwitchRowCheckedState(idx)
								},
							},
						},
					},
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【角色锻造数据信息】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
					},
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【角色门客数据信息】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
					},
					dcl.Composite{ // 这里重新布局
						ColumnSpan: 1,
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:     dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.TableView{
								AssignTo:         &dlg.EquipComposeDataTV,
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
								Model: dlg.EquipComposeDataModel,
								OnItemActivated: func() {
									//	idx := pg.roleBaseDataTV.CurrentIndex()
									//	pg.roleBaseDataModel.SwitchRowCheckedState(idx)
								},
							},
						},
					},
				},
			},
			dcl.Composite{
				Layout: dcl.HBox{},
				Children: []dcl.Widget{
					dcl.HSpacer{},
					dcl.PushButton{
						AssignTo: &acceptPB,
						Text:     "OK",
						OnClicked: func() {
							dlg.Accept()
						},
					},
					dcl.PushButton{
						AssignTo:  &cancelPB,
						Text:      "Cancel",
						OnClicked: func() { dlg.Cancel() },
					},
				},
			},
		},
	}

	err := o.Create(parent)
	if err != nil {
		return false
	}

	return true
}
