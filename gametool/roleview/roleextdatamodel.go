package main

import (
	"sort"

	"github.com/heartchord/goblazer"
	"github.com/heartchord/jxonline/gamestruct"
	"github.com/lxn/walk"
)

// RoleExtDataOfBaseItem :
type RoleExtDataOfBaseItem struct {
	Index         int    // 索引
	Name          string // 数据名称
	Content       string // 数据内容
	Comment       string // 数据说明
	OriginContent string // 数据原始名称（数据内容可能改变）
	checked       bool   // 是否选中
}

// RoleExtDataOfBaseModel :
type RoleExtDataOfBaseModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	itemIcon   *walk.Icon
	items      []*RoleExtDataOfBaseItem
}

// NewRoleExtDataOfBaseModel :
func NewRoleExtDataOfBaseModel() *RoleExtDataOfBaseModel {
	var err error

	m := new(RoleExtDataOfBaseModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {

	}

	m.ResetRows(nil)
	return m
}

// RowCount :
func (m *RoleExtDataOfBaseModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleExtDataOfBaseModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.Name

	case 2:
		return item.Content

	case 3:
		return item.Comment
	}

	panic("unexpected col")
}

// Checked :
func (m *RoleExtDataOfBaseModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleExtDataOfBaseModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleExtDataOfBaseModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *RoleExtDataOfBaseModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleExtDataOfBaseModel) Less(i int, j int) bool {
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
		return f(a.Name < b.Name)

	case 2:
		return f(a.Content < b.Content)

	case 3:
		return f(a.Comment < b.Comment)
	}

	panic("Unreachable Column Index!")
}

// Swap : 数据交换
func (m *RoleExtDataOfBaseModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleExtDataOfBaseModel) Image(row int) interface{} {
	return m.itemIcon
}

// ResetRows :
func (m *RoleExtDataOfBaseModel) ResetRows(data *gamestruct.RoleExtDataOfBase) {
	if data == nil {
		return
	}

	fieldCount := int(goblazer.GetStructFieldNum(*data)) // 获取角色基础数据成员个数
	fieldNames := goblazer.GetStructFieldNames(*data)    // 获取角色基础数据成员名称
	fieldStrings := getStructFieldStrings(*data)         // 获取角色基础数据成员内容
	fieldTags := goblazer.GetStructFieldTags(*data)

	if m.items == nil {
		m.items = make([]*RoleExtDataOfBaseItem, fieldCount)
	}

	for i := 0; i < fieldCount; i++ {
		m.items[i] = &RoleExtDataOfBaseItem{
			Index:         i,
			Name:          fieldNames[i],
			Content:       fieldStrings[i],
			OriginContent: fieldStrings[i],
			Comment:       fieldTags[i],
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
}

// SetRowContent :
func (m *RoleExtDataOfBaseModel) SetRowContent(idx int, content string) {
	//m.items[idx].Content = content
	//roleBakPage.roleBaseDataModel.PublishRowChanged(idx)
}

// SwitchRowCheckedState :
func (m *RoleExtDataOfBaseModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleExtDataOfBaseModel) Items() []*RoleExtDataOfBaseItem {
	return m.items
}

// RoleExtDataOfLingLongLockItem :
type RoleExtDataOfLingLongLockItem struct {
	Index         int    // 索引
	Name          string // 数据名称
	Content       string // 数据内容
	Comment       string // 数据说明
	OriginContent string // 数据原始名称（数据内容可能改变）
	checked       bool   // 是否选中
}

// RoleExtDataOfLingLongLockModel :
type RoleExtDataOfLingLongLockModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	itemIcon   *walk.Icon
	items      []*RoleExtDataOfLingLongLockItem
}

// NewRoleExtDataOfLingLongLockModel :
func NewRoleExtDataOfLingLongLockModel() *RoleExtDataOfLingLongLockModel {
	var err error

	m := new(RoleExtDataOfLingLongLockModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {

	}

	m.ResetRows(nil)
	return m
}

// RowCount :
func (m *RoleExtDataOfLingLongLockModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleExtDataOfLingLongLockModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Index

	case 1:
		return item.Name

	case 2:
		return item.Content

	case 3:
		return item.Comment
	}

	panic("unexpected col")
}

// Checked :
func (m *RoleExtDataOfLingLongLockModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleExtDataOfLingLongLockModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleExtDataOfLingLongLockModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *RoleExtDataOfLingLongLockModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleExtDataOfLingLongLockModel) Less(i int, j int) bool {
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
		return f(a.Name < b.Name)

	case 2:
		return f(a.Content < b.Content)

	case 3:
		return f(a.Comment < b.Comment)
	}

	panic("Unreachable Column Index!")
}

// Swap : 数据交换
func (m *RoleExtDataOfLingLongLockModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleExtDataOfLingLongLockModel) Image(row int) interface{} {
	return m.itemIcon
}

// ResetRows :
func (m *RoleExtDataOfLingLongLockModel) ResetRows(data *gamestruct.RoleExtDataOfLingLongLock) {
	if data == nil {
		return
	}

	fieldCount := int(goblazer.GetStructFieldNum(*data)) // 获取角色基础数据成员个数
	fieldNames := goblazer.GetStructFieldNames(*data)    // 获取角色基础数据成员名称
	fieldStrings := getStructFieldStrings(*data)         // 获取角色基础数据成员内容
	fieldTags := goblazer.GetStructFieldTags(*data)

	if m.items == nil {
		m.items = make([]*RoleExtDataOfLingLongLockItem, fieldCount)
	}

	for i := 0; i < fieldCount; i++ {
		m.items[i] = &RoleExtDataOfLingLongLockItem{
			Index:         i,
			Name:          fieldNames[i],
			Content:       fieldStrings[i],
			OriginContent: fieldStrings[i],
			Comment:       fieldTags[i],
		}
	}

	// Notify TableView and other interested parties about the reset.
	m.PublishRowsReset()
	m.Sort(m.sortColumn, m.sortOrder)
}

// SetRowContent :
func (m *RoleExtDataOfLingLongLockModel) SetRowContent(idx int, content string) {
	//m.items[idx].Content = content
	//roleBakPage.roleBaseDataModel.PublishRowChanged(idx)
}

// SwitchRowCheckedState :
func (m *RoleExtDataOfLingLongLockModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleExtDataOfLingLongLockModel) Items() []*RoleExtDataOfLingLongLockItem {
	return m.items
}
