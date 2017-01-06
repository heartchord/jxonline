package gameencoder

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	gmstruct "github.com/heartchord/jxonline/gamestruct"
)

// 123
const (
	roleExtDataOfItem = iota
	roleExtDataOfBase
	roleExtDataOfLingLongLock
	roleExtDataTypeOfHangerOn
	roleExtDataTypeOfTransNimbus
	roleExtDataTypeOfBreak
	roleExtDataTypeOfEquipCompose
	roleExtDataTypeCount
)

// playerExtDataReader :
type playerExtDataReader func(data []byte, current *uint32) bool

// RoleEncoder : a data struct of role bak encoder and decoder
type RoleEncoder struct {
	RoleBaseData       gmstruct.RoleBaseData       // 角色基础数据
	FSkillData         []gmstruct.SkillData        // 战斗技能数据
	LSkillData         []gmstruct.SkillData        // 生活技能数据
	TaskData           []gmstruct.TaskData         // 任务变量数据
	ItemData           []gmstruct.ItemData         // 装备物品数据
	SkillState         []gmstruct.SkillState       // 技能状态数据
	SkillCD            []gmstruct.SkillCD          // 技能冷却数据
	FeatureInfo        []gmstruct.FeatureInfo      // 角色外观数据
	PlayerEvent        []gmstruct.PlayerEvent      // 角色事件数据
	PlayerTitle        []gmstruct.RoleTitle        // 角色称号数据
	CustomStructHeader []gmstruct.CustomDataHeader // 自定义数据头
	RoleExtData        gmstruct.RoleExtData        // 角色扩展数据
	CRC32Cal           uint32                      // 根据角色原始数据计算的CRC32码
	CRC32Read          uint32                      // 从角色原始数据末尾读的CRC32码

	extDataReader map[int32]playerExtDataReader // 角色扩展数据解析函数
}

// Init : 123
func (en *RoleEncoder) Init() bool {
	// 初始化ReadFunction
	en.extDataReader = make(map[int32]playerExtDataReader)
	en.extDataReader[roleExtDataOfItem] = en.decodeRoleExtDataOfItem
	en.extDataReader[roleExtDataOfBase] = en.decodeRoleExtDataOfBase
	en.extDataReader[roleExtDataOfLingLongLock] = en.decodeRoleExtDataOfLingLongLock
	en.extDataReader[roleExtDataTypeOfHangerOn] = en.decodeRoleExtDataOfHangerOn
	en.extDataReader[roleExtDataTypeOfTransNimbus] = en.decodeRoleExtDataOfTransNimbus
	en.extDataReader[roleExtDataTypeOfBreak] = en.decodeRoleExtDataOfBreak
	en.extDataReader[roleExtDataTypeOfEquipCompose] = en.decodeRoleExtDataOfEquipCompose

	return true
}

// Decode : function to decode original role bak data
func (en *RoleEncoder) Decode(data []byte) bool {
	current := uint32(0)

	// 计算CRC32
	dataLen := len(data)
	en.CRC32Cal = CRC32(0, data[:dataLen-4])

	// 角色基本信息解码
	if !en.decodeRoleBaseInfo(data, &current) {
		return false
	}

	fmt.Printf("CurrentPos = %-4d, SkillOffset   = %-4d\n", current, en.RoleBaseData.FSkillOffset)

	// 角色战斗技能解码
	if !en.decodeRoleFSkillData(data, &current) {
		return false
	}

	fmt.Printf("CurrentPos = %-4d, LSkillOffset  = %-4d\n", current, en.RoleBaseData.LSkillOffset)

	// 角色生活技能解码
	if !en.decodeRoleLSkillData(data, &current) {
		return false
	}

	fmt.Printf("CurrentPos = %-4d, TaskOffset    = %-4d\n", current, en.RoleBaseData.TaskOffset)

	// 角色任务变量解码
	if !en.decodeRoleTaskData(data, &current) {
		return false
	}

	fmt.Printf("CurrentPos = %-4d, ItemOffset    = %-4d\n", current, en.RoleBaseData.ItemOffset)

	// 角色装备道具解码
	if !en.decodeRoleItemData(data, &current) {
		return false
	}

	fmt.Printf("CurrentPos = %-4d, StateOffset   = %-4d\n", current, en.RoleBaseData.StateOffset)

	if !en.decodeRoleStateList(data, &current) {
		return false
	}

	fmt.Printf("CurrentPos = %-4d, ExtBuffOffset = %-4d\n", current, en.RoleBaseData.ExtBuffOffset)

	if !en.decodeRoleExtData(data, &current) {
		return false
	}

	fmt.Printf("CurrentPos = %-4d, RoleDataLen   = %-4d\n", current, en.RoleBaseData.DataLen)

	return true
}

// PrintAllTaskData : function to print all task data
func (en RoleEncoder) PrintAllTaskData() {
	count := len(en.TaskData)

	fmt.Println("=================================[TASK VALUE]=================================")
	fmt.Printf("Total = %d\n", count)

	for i := 0; i < count; i++ {
		fmt.Printf("Task[ %-4d ] = %-10d", en.TaskData[i].TaskID, en.TaskData[i].TaskValue)
		if i%2 == 1 {
			fmt.Println("")
		}
		if i%2 == 0 {
			fmt.Print("\t")
		}
	}
	fmt.Print("\n")
}

// PrintAllFSkillData : function to print all fight skill data
func (en RoleEncoder) PrintAllFSkillData() {
	count := len(en.FSkillData)

	fmt.Println("==============================[FIGHT SKILL DATA]==============================")
	fmt.Printf("Total = %d\n", count)

	for i := 0; i < count; i++ {
		fmt.Printf("Skill[ %-4d ] = { %-2d, %-10d }", en.FSkillData[i].SkillID, en.FSkillData[i].SkillLv, en.FSkillData[i].SkillExp)
		if i%2 == 1 {
			fmt.Println("")
		}
		if i%2 == 0 {
			fmt.Print("\t")
		}
	}
	fmt.Print("\n")
}

// PrintAllLSkillData : function to print all life skill data
func (en RoleEncoder) PrintAllLSkillData() {
	count := len(en.LSkillData)

	fmt.Println("==============================[LIFE SKILL DATA]===============================")
	fmt.Printf("Total = %d\n", count)

	for i := 0; i < count; i++ {
		fmt.Printf("Skill[ %-4d ] = { %-2d, %-10d }", en.LSkillData[i].SkillID, en.LSkillData[i].SkillLv, en.LSkillData[i].SkillExp)
		if i%2 == 1 {
			fmt.Println("")
		}
		if i%2 == 0 {
			fmt.Print("\t")
		}
	}
	fmt.Print("\n")
}

// PrintAllItemData : function to print all Item data
func (en RoleEncoder) PrintAllItemData() {
	count := len(en.ItemData)

	fmt.Println("=================================[Item DATA]==================================")
	fmt.Printf("Total = %d\n", count)

	for i := 0; i < count; i++ {
		fmt.Printf("Item[ %-3d ] = { G = %d, D = %d, P = %-4d, Lv = %-2d, Place = %-2d }\n", i,
			en.ItemData[i].Standard.ClassCode&0x0000FFFF, en.ItemData[i].Standard.DetailType, en.ItemData[i].Standard.ParticularType,
			en.ItemData[i].Standard.Level, en.ItemData[i].Standard.Place)

		if i%20 == 19 {
			reader := bufio.NewReader(os.Stdin)
			reader.ReadLine()
		}
	}
}

// PrintAllSkillData : function to print all skill data
func (en RoleEncoder) PrintAllSkillData() {
	en.PrintAllFSkillData()
	en.PrintAllLSkillData()
}

func (en *RoleEncoder) decodeRoleBaseInfo(data []byte, current *uint32) bool {

	dataLen := uint32(len(data))
	structLen := uint32(binary.Size(en.RoleBaseData))

	start := *current
	if start+structLen > dataLen { // 数据长度 < 角色名数据头长度
		fmt.Println(start)
		fmt.Println(structLen)
		fmt.Println(dataLen)
		return false
	}

	end := *current + structLen
	buf := bytes.NewBuffer(data[start:end])
	binary.Read(buf, binary.LittleEndian, &en.RoleBaseData)

	*current += structLen
	return true
}

func (en *RoleEncoder) decodeRoleFSkillData(data []byte, current *uint32) bool {

	ret, skillCount := en.getFSkillCount()
	if !ret {
		return false
	}
	if skillCount == 0 {
		return true
	}
	en.FSkillData = make([]gmstruct.SkillData, skillCount)

	start := *current
	dataLen := uint32(len(data))
	totalLen := uint32(binary.Size(en.FSkillData))

	if start+totalLen > dataLen { // 数据长度 < 技能数据长度
		return false
	}

	structLen := uint32(binary.Size(en.FSkillData[0]))
	end := start + structLen
	for i := uint32(0); i < skillCount; i++ {
		buf := bytes.NewBuffer(data[start:end])
		binary.Read(buf, binary.LittleEndian, &en.FSkillData[i])
		start += structLen
		end += structLen
	}

	*current += totalLen
	return true
}

func (en *RoleEncoder) decodeRoleLSkillData(data []byte, current *uint32) bool {

	ret, skillCount := en.getLSkillCount()
	if !ret {
		return false
	}
	if skillCount == 0 {
		return true
	}
	en.LSkillData = make([]gmstruct.SkillData, skillCount)

	start := *current
	dataLen := uint32(len(data))
	totalLen := uint32(binary.Size(en.LSkillData))

	if start+totalLen > dataLen { // 数据长度 < 技能数据长度
		return false
	}

	structLen := uint32(binary.Size(en.LSkillData[0]))
	end := start + structLen

	for i := uint32(0); i < skillCount; i++ {
		buf := bytes.NewBuffer(data[start:end])
		binary.Read(buf, binary.LittleEndian, &en.LSkillData[i])
		start += structLen
		end += structLen
	}

	*current += totalLen
	return true
}

func (en *RoleEncoder) decodeRoleTaskData(data []byte, current *uint32) bool {

	ret, taskCount := en.getTaskCount()
	if !ret {
		return false
	}
	if taskCount == 0 {
		return true
	}
	en.TaskData = make([]gmstruct.TaskData, taskCount)

	start := *current
	dataLen := uint32(len(data))
	totalLen := uint32(binary.Size(en.TaskData))

	if start+totalLen > dataLen { // 数据长度 < 任务变量数据长度
		return false
	}

	structLen := uint32(binary.Size(en.TaskData[0]))
	end := start + structLen

	for i := uint32(0); i < taskCount; i++ {
		buf := bytes.NewBuffer(data[start:end])
		binary.Read(buf, binary.LittleEndian, &en.TaskData[i])
		start += structLen
		end += structLen
	}

	*current += totalLen
	return true
}

func (en *RoleEncoder) decodeRoleItemData(data []byte, current *uint32) bool {

	if en.RoleBaseData.ItemCount <= 0 { // 角色身上没有物品，不解析
		return true
	}
	en.ItemData = make([]gmstruct.ItemData, en.RoleBaseData.ItemCount)

	end := *current
	start := *current
	counter := int16(0)
	structLen := uint32(0)

	var header gmstruct.DataHead
	for counter < en.RoleBaseData.ItemCount {
		// 解析DataHead
		structLen = uint32(binary.Size(header))
		end = start + structLen
		buf := bytes.NewBuffer(data[start:end])
		binary.Read(buf, binary.LittleEndian, &header)
		start += structLen
		*current += structLen

		for i := int16(0); i < header.DataCount; i++ {
			en.ItemData[counter].HasStandard = (header.DataType&0xffff)&1 != 0
			en.ItemData[counter].HasLockSoul = (header.DataType&0xffff)&2 != 0
			en.ItemData[counter].HasBill = (header.DataType&0xffff)&4 != 0
			en.ItemData[counter].HasExtend = (header.DataType&0xffff)&8 != 0

			if en.ItemData[counter].HasStandard {
				structLen = uint32(binary.Size(en.ItemData[counter].Standard))
				end = start + structLen
				buf = bytes.NewBuffer(data[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].Standard)
				start += structLen
				*current += structLen
			}

			if en.ItemData[counter].HasLockSoul {
				structLen = uint32(binary.Size(en.ItemData[counter].LockSoul))
				end = start + structLen
				buf = bytes.NewBuffer(data[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].LockSoul)
				start += structLen
				*current += structLen
			}

			if en.ItemData[counter].HasBill {
				structLen = uint32(binary.Size(en.ItemData[counter].Bill))
				end = start + structLen
				buf = bytes.NewBuffer(data[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].Bill)
				start += structLen
				*current += structLen
			}

			if en.ItemData[counter].HasExtend {
				structLen = uint32(binary.Size(en.ItemData[counter].Extend))
				end = start + structLen
				buf = bytes.NewBuffer(data[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].Extend)
				start += structLen
				*current += structLen
			}

			counter++
		}
	}

	return true
}

func (en *RoleEncoder) decodeRoleStateList(data []byte, current *uint32) bool {
	// 角色身上没有状态信息，不解析
	if en.RoleBaseData.StateCount <= 0 {
		return true
	}

	var stateData gmstruct.StateData
	var end uint32
	var start = *current

	for i := int16(0); i < en.RoleBaseData.StateCount; i++ {
		// 解码StateData
		stateDataLen := uint32(binary.Size(stateData))
		end = start + stateDataLen
		buf := bytes.NewBuffer(data[start:end])
		binary.Read(buf, binary.LittleEndian, &stateData)

		// 根据类型进行解码
		switch stateData.Type {
		case gmstruct.SkillStateType:
			{
				var state gmstruct.SkillState
				structLen := uint32(binary.Size(state))
				buf := bytes.NewBuffer(stateData.Data[0:structLen])
				binary.Read(buf, binary.LittleEndian, &state)
				en.SkillState = append(en.SkillState, state)

				start += stateDataLen
				*current += stateDataLen
			}
		case gmstruct.SkillCDType:
			{
				var cd gmstruct.SkillCD
				structLen := uint32(binary.Size(cd))
				buf := bytes.NewBuffer(stateData.Data[0:structLen])
				binary.Read(buf, binary.LittleEndian, &cd)
				en.SkillCD = append(en.SkillCD, cd)

				start += stateDataLen
				*current += stateDataLen
			}
		case gmstruct.FeatureInfoType:
			{
				var info gmstruct.FeatureInfo
				structLen := uint32(binary.Size(info))
				buf := bytes.NewBuffer(stateData.Data[0:structLen])
				binary.Read(buf, binary.LittleEndian, &info)
				en.FeatureInfo = append(en.FeatureInfo, info)

				start += stateDataLen
				*current += stateDataLen
			}
		case gmstruct.PlayerEventInfoType:
			{
				var event gmstruct.PlayerEvent
				structLen := uint32(binary.Size(event))
				buf := bytes.NewBuffer(stateData.Data[0:structLen])
				binary.Read(buf, binary.LittleEndian, &event)
				en.PlayerEvent = append(en.PlayerEvent, event)

				start += stateDataLen
				*current += stateDataLen
			}
		case gmstruct.PlayerTitleType:
			{
				var title gmstruct.RoleTitle
				structLen := uint32(binary.Size(title))
				buf := bytes.NewBuffer(stateData.Data[0:structLen])
				binary.Read(buf, binary.LittleEndian, &title)
				en.PlayerTitle = append(en.PlayerTitle, title)

				start += stateDataLen
				*current += stateDataLen
			}
		case gmstruct.CustomStructType:
			{ // 用户自定义数据头，真正数据在数据头之后
				// 用户自定义数据可能比gmstruct.CustomStructHeader.Data小
				// 如果用户数据太短，这里会发生问题
				var custom gmstruct.CustomDataHeader
				structLen := uint32(binary.Size(custom))
				buf := bytes.NewBuffer(stateData.Data[0:structLen])
				binary.Read(buf, binary.LittleEndian, &custom)
				en.CustomStructHeader = append(en.CustomStructHeader, custom)

				// 处理用户自定义数据体
				switch custom.Type {
				case gmstruct.CustomDataTypeOfPartner:
					{
						en.decodeCustomDataOfPartner(data[start+1 : start+1+custom.Size])
					}
				}

				// 跳过用户自定义数据体
				start += custom.Size + 1    // 用户自定义数据体长度 + gmstruct.CustomDataHeader.Type
				*current += custom.Size + 1 // 用户自定义数据体长度 + gmstruct.CustomDataHeader.Type
			}
		default:
			{
				fmt.Printf("DecodeRoleStateList : unexpected type - %d!\n", stateData.Type)
				return false
			}
		}
	}
	return true
}

func (en *RoleEncoder) decodeCustomDataOfPartner(data []byte) bool {
	var header gmstruct.CustomDataOfPartnerHeader

	start := uint32(0)
	structLen := uint32(binary.Size(header))
	end := start + structLen
	dataLen := uint32(len(data))

	if dataLen < structLen {
		return false
	}

	buf := bytes.NewBuffer(data[start:end])
	binary.Read(buf, binary.LittleEndian, &header)
	start += structLen

	if header.PartnerCount <= 0 {
		return true
	}

	for i := byte(0); i < header.PartnerCount; i++ {
		// 解析同伴数据，由于目前没有同伴数据，这里暂时不解析了
	}

	return true
}

func (en *RoleEncoder) decodeRoleExtData(data []byte, current *uint32) bool {
	var header gmstruct.DataHead

	if *current != en.RoleBaseData.ExtBuffOffset && en.RoleBaseData.ExtBuffOffset > 0 {
		// 如果偏移出错，使用ExtBuffOffset修正
		*current = en.RoleBaseData.ExtBuffOffset
	}

	dataLen := uint32(len(data))
	headerSize := uint32(binary.Size(header))

	start := *current
	end := *current

	for start+headerSize <= dataLen {
		// 解析gmstruct.DataHead
		end = start + headerSize
		buf := bytes.NewBuffer(data[start:end])
		binary.Read(buf, binary.LittleEndian, &header)
		start += headerSize
		*current += headerSize

		// 获取数据类型
		t := header.DataType >> 16
		if t >= roleExtDataTypeCount {
			return false
		}

		// 解析角色扩展数据
		if !en.extDataReader[t](data, current) {
			return false
		}

		// 解析完跳过数据体
		start = *current
	}

	return true
}

func (en *RoleEncoder) decodeRoleExtDataOfItem(data []byte, current *uint32) bool {
	//fmt.Println("RoleEncoder.decodeRoleExtDataOfItem")
	return true
}

func (en *RoleEncoder) decodeRoleExtDataOfBase(data []byte, current *uint32) bool {
	//fmt.Println("RoleEncoder.decodeRoleExtDataOfBase")

	dataLen := uint32(len(data))
	structLen := uint32(binary.Size(en.RoleExtData.Base))

	if *current+structLen > dataLen {
		return false
	}

	buf := bytes.NewBuffer(data[*current : *current+structLen])
	binary.Read(buf, binary.LittleEndian, &en.RoleExtData.Base)
	*current += structLen

	return true
}

func (en *RoleEncoder) decodeRoleExtDataOfLingLongLock(data []byte, current *uint32) bool {
	//fmt.Println("RoleEncoder.decodeRoleExtDataOfLingLongLock")

	dataLen := uint32(len(data))
	structLen := uint32(binary.Size(en.RoleExtData.LingLongLock))

	if *current+structLen > dataLen {
		return false
	}

	buf := bytes.NewBuffer(data[*current : *current+structLen])
	binary.Read(buf, binary.LittleEndian, &en.RoleExtData.LingLongLock)
	*current += structLen

	return true
}

func (en *RoleEncoder) decodeRoleExtDataOfHangerOn(data []byte, current *uint32) bool {
	//fmt.Println("RoleEncoder.decodeRoleExtDataOfHangerOn")

	dataLen := uint32(len(data))
	structLen := uint32(binary.Size(en.RoleExtData.HangerOn))

	if *current+structLen > dataLen {
		return false
	}

	buf := bytes.NewBuffer(data[*current : *current+structLen])
	binary.Read(buf, binary.LittleEndian, &en.RoleExtData.HangerOn)
	*current += structLen

	return true
}

func (en *RoleEncoder) decodeRoleExtDataOfTransNimbus(data []byte, current *uint32) bool {
	//fmt.Println("RoleEncoder.decodeRoleExtDataOfTransNimbus")

	dataLen := uint32(len(data))
	structLen := uint32(binary.Size(en.RoleExtData.TransNimbus))

	if *current+structLen > dataLen {
		return false
	}

	buf := bytes.NewBuffer(data[*current : *current+structLen])
	binary.Read(buf, binary.LittleEndian, &en.RoleExtData.TransNimbus)
	*current += structLen

	return true
}

func (en *RoleEncoder) decodeRoleExtDataOfBreak(data []byte, current *uint32) bool {
	//fmt.Println("RoleEncoder.decodeRoleExtDataOfBreak")

	dataLen := uint32(len(data))
	structLen := uint32(binary.Size(en.RoleExtData.Break))

	if *current+structLen > dataLen {
		return false
	}

	buf := bytes.NewBuffer(data[*current : *current+structLen])
	binary.Read(buf, binary.LittleEndian, &en.RoleExtData.Break)
	*current += structLen

	return true
}

func (en *RoleEncoder) decodeRoleExtDataOfEquipCompose(data []byte, current *uint32) bool {
	//fmt.Println("RoleEncoder.decodeRoleExtDataOfEquipCompose")

	dataLen := uint32(len(data))
	structLen := uint32(binary.Size(en.RoleExtData.EquipCompose))

	if *current+structLen > dataLen {
		return false
	}

	buf := bytes.NewBuffer(data[*current : *current+structLen])
	binary.Read(buf, binary.LittleEndian, &en.RoleExtData.EquipCompose)
	*current += structLen

	return true
}

func (en *RoleEncoder) getFSkillCount() (bool, uint32) {
	if en.RoleBaseData.LSkillOffset < en.RoleBaseData.FSkillOffset {
		return false, 0
	}

	var skill gmstruct.SkillData
	skillDataSize := uint32(binary.Size(skill))
	return true, (en.RoleBaseData.LSkillOffset - en.RoleBaseData.FSkillOffset) / skillDataSize
}

func (en *RoleEncoder) getLSkillCount() (bool, uint32) {
	if en.RoleBaseData.TaskOffset < en.RoleBaseData.LSkillOffset {
		return false, 0
	}

	var skill gmstruct.SkillData
	skillDataSize := uint32(binary.Size(skill))
	return true, (en.RoleBaseData.TaskOffset - en.RoleBaseData.LSkillOffset) / skillDataSize
}

func (en *RoleEncoder) getTaskCount() (bool, uint32) {
	if en.RoleBaseData.ItemOffset < en.RoleBaseData.TaskOffset {
		return false, 0
	}

	var task gmstruct.TaskData
	taskDataSize := uint32(binary.Size(task))
	return true, (en.RoleBaseData.ItemOffset - en.RoleBaseData.TaskOffset) / taskDataSize
}
