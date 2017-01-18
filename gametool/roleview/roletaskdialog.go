package main

import (
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// RoleTaskDialog :
type RoleTaskDialog struct {
	*walk.Dialog

	RoleTaskDataModel *RoleTaskDataModel
	RoleTaskDataTV    *walk.TableView
}

// CreateInstance :
func (dlg *RoleTaskDialog) CreateInstance(parent walk.Form) bool {
	dlg.RoleTaskDataModel = NewRoleTaskDataModel()
	return dlg.CreateRoleTaskDialog(parent)
}

// CreateRoleTaskDialog :
func (dlg *RoleTaskDialog) CreateRoleTaskDialog(parent walk.Form) bool {
	var acceptPB, cancelPB *walk.PushButton

	o := &dcl.Dialog{
		AssignTo:      &dlg.Dialog,
		Title:         "角色任务变量信息",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       dcl.Size{Width: 1000, Height: 750},
		Font:          dcl.Font{Family: "微软雅黑", PointSize: 10},
		Layout:        dcl.VBox{},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.Grid{Columns: 1, Spacing: 10},
				Children: []dcl.Widget{
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【任务变量】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
					},
					dcl.Composite{ // 这里重新布局
						ColumnSpan: 1,
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:     dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.TableView{
								AssignTo:         &dlg.RoleTaskDataTV,
								ColumnSpan:       1,
								CheckBoxes:       true,
								ColumnsOrderable: true,
								MultiSelection:   true,
								MinSize:          dcl.Size{Width: 0, Height: 300},
								Columns: []dcl.TableViewColumn{
									{Title: "数据索引"},
									{Title: "任务变量ID"},
									{Title: "任务变量值"},
								},
								Model: dlg.RoleTaskDataModel,
								OnItemActivated: func() {
									idx := dlg.RoleTaskDataTV.CurrentIndex()
									dlg.RoleTaskDataModel.SwitchRowCheckedState(idx)
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
