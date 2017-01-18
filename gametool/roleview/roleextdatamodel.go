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

// RoleExtDataOfTransNimbusItem :
type RoleExtDataOfTransNimbusItem struct {
	Index         int    // 索引
	Name          string // 数据名称
	Content       string // 数据内容
	Comment       string // 数据说明
	OriginContent string // 数据原始名称（数据内容可能改变）
	checked       bool   // 是否选中
}

// RoleExtDataOfTransNimbusModel :
type RoleExtDataOfTransNimbusModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	itemIcon   *walk.Icon
	items      []*RoleExtDataOfTransNimbusItem
}

// NewRoleExtDataOfTransNimbusModel :
func NewRoleExtDataOfTransNimbusModel() *RoleExtDataOfTransNimbusModel {
	var err error

	m := new(RoleExtDataOfTransNimbusModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {

	}

	m.ResetRows(nil)
	return m
}

// RowCount :
func (m *RoleExtDataOfTransNimbusModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleExtDataOfTransNimbusModel) Value(row, col int) interface{} {
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
func (m *RoleExtDataOfTransNimbusModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleExtDataOfTransNimbusModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleExtDataOfTransNimbusModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *RoleExtDataOfTransNimbusModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleExtDataOfTransNimbusModel) Less(i int, j int) bool {
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
func (m *RoleExtDataOfTransNimbusModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleExtDataOfTransNimbusModel) Image(row int) interface{} {
	return m.itemIcon
}

// ResetRows :
func (m *RoleExtDataOfTransNimbusModel) ResetRows(data *gamestruct.RoleExtDataOfTransNimbus) {
	if data == nil {
		return
	}

	fieldCount := int(goblazer.GetStructFieldNum(*data)) // 获取角色基础数据成员个数
	fieldNames := goblazer.GetStructFieldNames(*data)    // 获取角色基础数据成员名称
	fieldStrings := getStructFieldStrings(*data)         // 获取角色基础数据成员内容
	fieldTags := goblazer.GetStructFieldTags(*data)

	if m.items == nil {
		m.items = make([]*RoleExtDataOfTransNimbusItem, fieldCount)
	}

	for i := 0; i < fieldCount; i++ {
		m.items[i] = &RoleExtDataOfTransNimbusItem{
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
func (m *RoleExtDataOfTransNimbusModel) SetRowContent(idx int, content string) {
	//m.items[idx].Content = content
	//roleBakPage.roleBaseDataModel.PublishRowChanged(idx)
}

// SwitchRowCheckedState :
func (m *RoleExtDataOfTransNimbusModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleExtDataOfTransNimbusModel) Items() []*RoleExtDataOfTransNimbusItem {
	return m.items
}

// RoleExtDataOfBreakItem :
type RoleExtDataOfBreakItem struct {
	Index         int    // 索引
	Name          string // 数据名称
	Content       string // 数据内容
	Comment       string // 数据说明
	OriginContent string // 数据原始名称（数据内容可能改变）
	checked       bool   // 是否选中
}

// RoleExtDataOfBreakModel :
type RoleExtDataOfBreakModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	itemIcon   *walk.Icon
	items      []*RoleExtDataOfBreakItem
}

// NewRoleExtDataOfBreakModel :
func NewRoleExtDataOfBreakModel() *RoleExtDataOfBreakModel {
	var err error

	m := new(RoleExtDataOfBreakModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {

	}

	m.ResetRows(nil)
	return m
}

// RowCount :
func (m *RoleExtDataOfBreakModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleExtDataOfBreakModel) Value(row, col int) interface{} {
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
func (m *RoleExtDataOfBreakModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleExtDataOfBreakModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleExtDataOfBreakModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *RoleExtDataOfBreakModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleExtDataOfBreakModel) Less(i int, j int) bool {
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
func (m *RoleExtDataOfBreakModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleExtDataOfBreakModel) Image(row int) interface{} {
	return m.itemIcon
}

// ResetRows :
func (m *RoleExtDataOfBreakModel) ResetRows(data *gamestruct.RoleExtDataOfBreak) {
	if data == nil {
		return
	}

	fieldCount := int(goblazer.GetStructFieldNum(*data)) // 获取角色基础数据成员个数
	fieldNames := goblazer.GetStructFieldNames(*data)    // 获取角色基础数据成员名称
	fieldStrings := getStructFieldStrings(*data)         // 获取角色基础数据成员内容
	fieldTags := goblazer.GetStructFieldTags(*data)

	if m.items == nil {
		m.items = make([]*RoleExtDataOfBreakItem, fieldCount)
	}

	for i := 0; i < fieldCount; i++ {
		m.items[i] = &RoleExtDataOfBreakItem{
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
func (m *RoleExtDataOfBreakModel) SetRowContent(idx int, content string) {
	//m.items[idx].Content = content
	//roleBakPage.roleBaseDataModel.PublishRowChanged(idx)
}

// SwitchRowCheckedState :
func (m *RoleExtDataOfBreakModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleExtDataOfBreakModel) Items() []*RoleExtDataOfBreakItem {
	return m.items
}

// RoleExtDataOfEquipComposeItem :
type RoleExtDataOfEquipComposeItem struct {
	Index         int    // 索引
	Name          string // 数据名称
	Content       string // 数据内容
	Comment       string // 数据说明
	OriginContent string // 数据原始名称（数据内容可能改变）
	checked       bool   // 是否选中
}

// RoleExtDataOfEquipComposeModel :
type RoleExtDataOfEquipComposeModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	itemIcon   *walk.Icon
	items      []*RoleExtDataOfEquipComposeItem
}

// NewRoleExtDataOfEquipComposeModel :
func NewRoleExtDataOfEquipComposeModel() *RoleExtDataOfEquipComposeModel {
	var err error

	m := new(RoleExtDataOfEquipComposeModel)
	m.itemIcon, err = walk.NewIconFromFile("../../gameresource/img/right-arrow2.ico")
	if err != nil {

	}

	m.ResetRows(nil)
	return m
}

// RowCount :
func (m *RoleExtDataOfEquipComposeModel) RowCount() int {
	return len(m.items)
}

// Value  :
func (m *RoleExtDataOfEquipComposeModel) Value(row, col int) interface{} {
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
func (m *RoleExtDataOfEquipComposeModel) Checked(row int) bool {
	return m.items[row].checked
}

// SetChecked :
func (m *RoleExtDataOfEquipComposeModel) SetChecked(row int, checked bool) error {
	m.items[row].checked = checked

	return nil
}

// Sort :
func (m *RoleExtDataOfEquipComposeModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *RoleExtDataOfEquipComposeModel) Len() int {
	return len(m.items)
}

// Less : 数据排序小于运算符
func (m *RoleExtDataOfEquipComposeModel) Less(i int, j int) bool {
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
func (m *RoleExtDataOfEquipComposeModel) Swap(i, j int) {
	m.items[i], m.items[j] = m.items[j], m.items[i]
}

// Image : 获取数据Item图标
func (m *RoleExtDataOfEquipComposeModel) Image(row int) interface{} {
	return m.itemIcon
}

// ResetRows :
func (m *RoleExtDataOfEquipComposeModel) ResetRows(data *gamestruct.RoleExtDataOfEquipCompose) {
	if data == nil {
		return
	}

	fieldCount := int(goblazer.GetStructFieldNum(*data)) // 获取角色基础数据成员个数
	fieldNames := goblazer.GetStructFieldNames(*data)    // 获取角色基础数据成员名称
	fieldStrings := getStructFieldStrings(*data)         // 获取角色基础数据成员内容
	fieldTags := goblazer.GetStructFieldTags(*data)

	if m.items == nil {
		m.items = make([]*RoleExtDataOfEquipComposeItem, fieldCount)
	}

	for i := 0; i < fieldCount; i++ {
		m.items[i] = &RoleExtDataOfEquipComposeItem{
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
func (m *RoleExtDataOfEquipComposeModel) SetRowContent(idx int, content string) {
	//m.items[idx].Content = content
	//roleBakPage.roleBaseDataModel.PublishRowChanged(idx)
}

// SwitchRowCheckedState :
func (m *RoleExtDataOfEquipComposeModel) SwitchRowCheckedState(idx int) {
	checked := m.Checked(idx)
	m.SetChecked(idx, !checked)
	m.PublishRowChanged(idx)
}

// Items :
func (m *RoleExtDataOfEquipComposeModel) Items() []*RoleExtDataOfEquipComposeItem {
	return m.items
}
