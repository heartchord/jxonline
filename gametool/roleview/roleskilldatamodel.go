package main

import (
	"sort"

	"github.com/heartchord/jxonline/gamestruct"
	"github.com/lxn/walk"
)

// RoleSkillDataItem :
type RoleSkillDataItem struct {
	DataModelItemBase
	SkillID  string // 数据名称
	SkillLv  string // 数据内容
	SkillExp string // 数据说明
}

// RoleSkillDataModel :
type RoleSkillDataModel struct {
	DataModelBase
	items []*RoleSkillDataItem
}

// NewRoleSkillDataModel :
func NewRoleSkillDataModel() *RoleSkillDataModel {
	var err error

	m := new(RoleSkillDataModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {
	}

	m.ResetRows(nil)
	return m
}

// RowCount :
func (m *RoleSkillDataModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleSkillDataModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.SkillID

	case 2:
		return item.SkillLv

	case 3:
		return item.SkillExp
	}

	panic("unexpected col")
}

// Checked :
func (m *RoleSkillDataModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleSkillDataModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleSkillDataModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

// Len :
func (m *RoleSkillDataModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleSkillDataModel) Less(i int, j int) bool {
	a, b := m.items[i], m.items[j]

	f := func(ls bool) bool {
		if m.sortOrder == walk.SortAscending { // 升序
			return ls
		}

		// 降序
		return !ls
	}

	switch m.sortColumn {
	case 0:
		return f(a.Index < b.Index)

	case 1:
		return f(a.SkillID < b.SkillID)

	case 2:
		return f(a.SkillLv < b.SkillLv)

	case 3:
		return f(a.SkillExp < b.SkillExp)
	}

	panic("Unreachable Column Index!")
}

// Swap : 数据交换
func (m *RoleSkillDataModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleSkillDataModel) Image(row int) interface{} {
	return m.itemIcon
}

// ResetRows :
func (m *RoleSkillDataModel) ResetRows(data []gamestruct.SkillData) {
	if data == nil {
		return
	}

	dataCount := len(data)
	if dataCount <= 0 {
		return
	}

	m.items = make([]*RoleSkillDataItem, dataCount)

	for i := 0; i < dataCount; i++ {
		fieldStrings := getStructFieldStrings(data[i])

		m.items[i] = &RoleSkillDataItem{}
		m.items[i].Index = i
		m.items[i].SkillID = fieldStrings[0]
		m.items[i].SkillLv = fieldStrings[1]
		m.items[i].SkillExp = fieldStrings[2]
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
}

// SwitchRowCheckedState :
func (m *RoleSkillDataModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleSkillDataModel) Items() []*RoleSkillDataItem {
	return m.items
}
