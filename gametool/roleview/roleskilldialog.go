package main

import (
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// RoleSkillDialog :
type RoleSkillDialog struct {
	*walk.Dialog

	RoleFSkillDataModel *RoleSkillDataModel
	RoleFSkillDataTV    *walk.TableView
}

// CreateInstance :
func (dlg *RoleSkillDialog) CreateInstance(parent walk.Form) bool {
	dlg.RoleFSkillDataModel = NewRoleSkillDataModel()
	return dlg.CreateRoleSkillDialog(parent)
}

// CreateRoleSkillDialog :
func (dlg *RoleSkillDialog) CreateRoleSkillDialog(parent walk.Form) bool {
	var acceptPB, cancelPB *walk.PushButton

	o := &dcl.Dialog{
		AssignTo:      &dlg.Dialog,
		Title:         "角色技能数据信息",
		DefaultButton: &acceptPB,
		CancelButton:  &cancelPB,
		MinSize:       dcl.Size{Width: 800, Height: 600},
		Font:          dcl.Font{Family: "微软雅黑", PointSize: 10},
		Layout:        dcl.VBox{},
		Children: []dcl.Widget{
			dcl.Composite{
				Layout: dcl.Grid{Columns: 1, Spacing: 10},
				Children: []dcl.Widget{
					dcl.Label{
						ColumnSpan: 1,
						Text:       "【战斗技能】",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
					},
					dcl.Composite{ // 这里重新布局
						ColumnSpan: 1,
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:     dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.TableView{
								AssignTo:         &dlg.RoleFSkillDataTV,
								ColumnSpan:       1,
								CheckBoxes:       true,
								ColumnsOrderable: true,
								MultiSelection:   true,
								Columns: []dcl.TableViewColumn{
									{Title: "数据索引"},
									{Title: "技能索引"},
									{Title: "技能等级"},
									{Title: "技能经验"},
								},
								Model: dlg.RoleFSkillDataModel,
								OnItemActivated: func() {
									idx := dlg.RoleFSkillDataTV.CurrentIndex()
									dlg.RoleFSkillDataModel.SwitchRowCheckedState(idx)
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
