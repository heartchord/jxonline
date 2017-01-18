package main

import (
	"sort"

	"github.com/heartchord/jxonline/gamestruct"
	"github.com/lxn/walk"
)

// RoleTaskDataItem :
type RoleTaskDataItem struct {
	Index     int    // 索引
	TaskID    string // 数据名称
	TaskValue string // 数据内容
	checked   bool   // 是否选中
}

// RoleTaskDataModel :
type RoleTaskDataModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	itemIcon   *walk.Icon
	items      []*RoleTaskDataItem
}

// NewRoleTaskDataModel :
func NewRoleTaskDataModel() *RoleTaskDataModel {
	var err error

	m := new(RoleTaskDataModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {
	}

	m.ResetRows(nil)
	return m
}

// RowCount :
func (m *RoleTaskDataModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleTaskDataModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.TaskID

	case 2:
		return item.TaskValue
	}

	panic("unexpected col")
}

// Checked :
func (m *RoleTaskDataModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleTaskDataModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleTaskDataModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

// Len :
func (m *RoleTaskDataModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleTaskDataModel) Less(i int, j int) bool {
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
		return f(a.TaskID < b.TaskID)

	case 2:
		return f(a.TaskValue < b.TaskValue)
	}

	panic("Unreachable Column Index!")
}

// Swap : 数据交换
func (m *RoleTaskDataModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleTaskDataModel) Image(row int) interface{} {
	return m.itemIcon
}

// ResetRows :
func (m *RoleTaskDataModel) ResetRows(data []gamestruct.TaskData) {
	if data == nil {
		return
	}

	dataCount := len(data)
	if dataCount <= 0 {
		return
	}

	m.items = make([]*RoleTaskDataItem, dataCount)

	for i := 0; i < dataCount; i++ {
		fieldStrings := getStructFieldStrings(data[i])
		for j := 0; j < len(fieldStrings); j++ {
		}

		m.items[i] = &RoleTaskDataItem{
			Index:     i,
			TaskID:    fieldStrings[0],
			TaskValue: fieldStrings[1],
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
}

// SwitchRowCheckedState :
func (m *RoleTaskDataModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleTaskDataModel) Items() []*RoleTaskDataItem {
	return m.items
}
