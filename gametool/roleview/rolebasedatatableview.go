package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"

	"github.com/heartchord/goblazer"
	"github.com/heartchord/jxonline/gamestruct"
	"github.com/henrylee2cn/mahonia"
	"github.com/lxn/walk"
)

// RoleBaseDataItem :
type RoleBaseDataItem struct {
	Index         int    // 索引
	Name          string // 数据名称
	Content       string // 数据内容
	Comment       string // 数据说明
	OriginContent string // 数据原始名称（数据内容可能改变）
	checked       bool   // 是否选中
}

// RoleBaseDataModel :
type RoleBaseDataModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	itemIcon   *walk.Icon
	items      []*RoleBaseDataItem
}

// NewRoleBaseDataModel :
func NewRoleBaseDataModel() *RoleBaseDataModel {
	var err error

	m := new(RoleBaseDataModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {

	}

	m.ResetRows(&gamestruct.RoleBaseData{})
	return m
}

// RowCount :
func (m *RoleBaseDataModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleBaseDataModel) Value(row, col int) interface{} {
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
func (m *RoleBaseDataModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleBaseDataModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleBaseDataModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *RoleBaseDataModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleBaseDataModel) Less(i int, j int) bool {
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
func (m *RoleBaseDataModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleBaseDataModel) Image(row int) interface{} {
	return m.itemIcon
}

func getStructFieldStrings(s interface{}) []string {
	var ret []string

	o := reflect.ValueOf(s)
	if o.Kind() != reflect.Struct {
		panic("Unexpected Data Type")
	}

	mdecoder := mahonia.NewDecoder("GBK")

	count := o.NumField()
	for i := 0; i < count; i++ {
		v := o.Field(i)
		if v.Kind() == reflect.Struct {
			sub := getStructFieldStrings(v.Interface())
			ret = append(ret, sub...)
			continue
		}

		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			var slice []byte

			str := ""
			length := v.Len()

			for i := 0; i < length; i++ {
				valStr := fmt.Sprintf("%v", v.Index(i))
				value, err := strconv.Atoi(valStr)

				if err != nil {
					panic("Wrong String Content")
				}

				if value != 0 {
					slice = append(slice, byte(value))
				}
				str = string(slice)
			}

			str = mdecoder.ConvertString(str)
			ret = append(ret, str)

		} else {
			ret = append(ret, fmt.Sprintf("%v", v))
		}
	}

	return ret
}

// ResetRows :
func (m *RoleBaseDataModel) ResetRows(data *gamestruct.RoleBaseData) {
	fieldCount := int(goblazer.GetStructFieldNum(*data)) // 获取角色基础数据成员个数
	fieldNames := goblazer.GetStructFieldNames(*data)    // 获取角色基础数据成员名称
	fieldStrings := getStructFieldStrings(*data)         // 获取角色基础数据成员内容
	fieldTags := goblazer.GetStructFieldTags(*data)

	if m.items == nil {
		m.items = make([]*RoleBaseDataItem, fieldCount)
	}

	for i := 0; i < fieldCount; i++ {
		m.items[i] = &RoleBaseDataItem{
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
func (m *RoleBaseDataModel) SetRowContent(idx int, content string) {
	m.items[idx].Content = content
	roleBakPage.roleBaseDataModel.PublishRowChanged(idx)
}

// SwitchRowCheckedState :
func (m *RoleBaseDataModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleBaseDataModel) Items() []*RoleBaseDataItem {
	return m.items
}
