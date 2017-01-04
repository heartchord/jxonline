package gameencoder

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"os"

	gmstruct "github.com/heartchord/jxonline/gamestruct"
)

// RoleBakHeader : a data struct of role bak data header
type RoleBakHeader struct {
	RoleNameLen uint32 // 角色名长度
	RoleNameGBK []byte // 角色名(GBK)格式
	RoleDataLen uint32 // 角色原始二进制数据长度
}

// RoleBakBody : a data struct of role bak data body
type RoleBakBody struct {
	RoleData []byte // Bak角色原始二进制数据
}

// RoleBakData : a data struct of role bak data
type RoleBakData struct {
	RoleBakHeader // Bak数据头 : 匿名字段
	RoleBakBody   // Bak数据体 : 匿名字段
}

// RoleBakEncoder : a data struct of role bak encoder and decoder
type RoleBakEncoder struct {
	BakData    RoleBakData
	RoleData   gmstruct.RoleData
	FSkillData []gmstruct.SkillData
	LSkillData []gmstruct.SkillData
	TaskData   []gmstruct.TaskData
	ItemData   []gmstruct.ItemData
}

// Decode : function to decode original role bak data
func (en *RoleBakEncoder) Decode(data []byte) bool {
	current := uint32(0)

	if !en.decodeBakHeader(data) {
		return false
	}

	// 角色基本信息解码
	if !en.decodeRoleBaseInfo(en.BakData.RoleData, &current) {
		return false
	}

	// 角色战斗技能解码
	if !en.decodeRoleFSkillData(en.BakData.RoleData, &current) {
		return false
	}

	// 角色生活技能解码
	if !en.decodeRoleLSkillData(en.BakData.RoleData, &current) {
		return false
	}

	// 角色任务变量解码
	if !en.decodeRoleTaskData(en.BakData.RoleData, &current) {
		return false
	}

	fmt.Printf("current = %d\n", current)
	if !en.decodeRoleItemData(en.BakData.RoleData, &current) {
		return false
	}
	return true
}

// PrintAllTaskData : function to print all task data
func (en RoleBakEncoder) PrintAllTaskData() {
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
func (en RoleBakEncoder) PrintAllFSkillData() {
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
func (en RoleBakEncoder) PrintAllLSkillData() {
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
func (en RoleBakEncoder) PrintAllItemData() {
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

	fmt.Printf("%d, %d, %d\n", en.ItemData[21].Bill.ExpiredTime, en.ItemData[21].Bill.CurrencyType, en.ItemData[21].Bill.ComeFromPlace)
	fmt.Print("\n")
}

// PrintAllSkillData : function to print all skill data
func (en RoleBakEncoder) PrintAllSkillData() {
	en.PrintAllFSkillData()
	en.PrintAllLSkillData()
}

func (en *RoleBakEncoder) decodeBakHeader(data []byte) bool {
	dataLen := uint32(len(data))
	current := uint32(0)

	if dataLen <= 4 { // 数据长度 <= 角色名数据头长度
		return false
	}

	// 获取角色名长度
	tmplen := uint32(0)
	tmpbuf := bytes.NewBuffer(data[current:4]) // [0, 3]存储角色名长度
	binary.Read(tmpbuf, binary.LittleEndian, &tmplen)
	en.BakData.RoleNameLen = tmplen
	current += 4

	// 打印调试信息
	//fmt.Printf("BakData.RoleNameLen = %d\n", en.BakData.RoleNameLen)

	// 获取角色名
	n := 4 + en.BakData.RoleNameLen
	if en.BakData.RoleNameLen <= 0 || dataLen <= n { // 角色名长度 <= 0 或 数据长度 <= 角色名数据头长度 + 角色名长度
		return false
	}

	en.BakData.RoleNameGBK = data[current:n] // [4, 4 + namelen]存储角色名
	current += en.BakData.RoleNameLen

	// 打印调试信息
	//mdecoder := mahonia.NewDecoder("GBK")
	//rolename := string(en.BakData.RoleNameGBK[:])
	//rolename = mdecoder.ConvertString(rolename)
	//fmt.Printf("BakData.RoleName = %s\n", rolename)

	// 获取角色原始数据长度
	n = 4 + en.BakData.RoleNameLen + 4
	if dataLen <= n { // 数据长度 <= 角色名数据头长度 + 角色名长度 + 角色数据长度
		return false
	}
	tmpbuf = bytes.NewBuffer(data[current:n])
	binary.Read(tmpbuf, binary.LittleEndian, &tmplen)
	en.BakData.RoleDataLen = tmplen // [4 + namelen, 4 + namelen + 4]存储角色原始数据长度
	current += 4

	// 打印调试信息
	//fmt.Printf("BakData.RoleDataLen = %d\n", en.BakData.RoleDataLen)

	// 获取角色原始数据
	n = 4 + en.BakData.RoleNameLen + 4 + en.BakData.RoleDataLen
	if dataLen < n {
		return false
	}
	en.BakData.RoleData = data[current:n]

	// 打印调试信息
	//fmt.Printf("len(BakData.Data) = %d\n", len(en.BakData.RoleData))
	return true
}

func (en *RoleBakEncoder) decodeRoleBaseInfo(data []byte, current *uint32) bool {
	dataTmp := data[*current:]
	dataLen := uint32(len(dataTmp))

	// 数据长度 < 角色名数据头长度
	len := uint32(binary.Size(en.RoleData))

	if dataLen < len {
		return false
	}

	buf := bytes.NewBuffer(dataTmp[:len])
	binary.Read(buf, binary.LittleEndian, &en.RoleData)
	*current += len

	return true
}

func (en *RoleBakEncoder) decodeRoleFSkillData(data []byte, current *uint32) bool {
	dataTmp := data[*current:]
	dataLen := uint32(len(dataTmp))

	ret, skillCount := en.getFSkillCount()
	if !ret {
		return false
	}
	if skillCount == 0 {
		return true
	}
	en.FSkillData = make([]gmstruct.SkillData, skillCount)

	// 数据长度 < 技能数据长度
	len := uint32(binary.Size(en.FSkillData))
	if dataLen < len {
		return false
	}

	structLen := uint32(binary.Size(en.FSkillData[0]))
	start := uint32(0)
	end := start + structLen

	for i := uint32(0); i < skillCount; i++ {
		buf := bytes.NewBuffer(dataTmp[start:end])
		binary.Read(buf, binary.LittleEndian, &en.FSkillData[i])
		start += structLen
		end += structLen
	}

	*current += len
	return true
}

func (en *RoleBakEncoder) decodeRoleLSkillData(data []byte, current *uint32) bool {
	dataTmp := data[*current:]
	dataLen := uint32(len(dataTmp))

	ret, skillCount := en.getLSkillCount()
	if !ret {
		return false
	}
	if skillCount == 0 {
		return true
	}
	en.LSkillData = make([]gmstruct.SkillData, skillCount)

	// 数据长度 < 技能数据长度
	len := uint32(binary.Size(en.LSkillData))
	if dataLen < len {
		return false
	}

	structLen := uint32(binary.Size(en.LSkillData[0]))
	start := uint32(0)
	end := start + structLen

	for i := uint32(0); i < skillCount; i++ {
		buf := bytes.NewBuffer(dataTmp[start:end])
		binary.Read(buf, binary.LittleEndian, &en.LSkillData[i])
		start += structLen
		end += structLen
	}

	*current += len
	return true
}

func (en *RoleBakEncoder) decodeRoleTaskData(data []byte, current *uint32) bool {
	dataTmp := data[*current:]
	dataLen := uint32(len(dataTmp))

	ret, taskCount := en.getTaskCount()
	if !ret {
		return false
	}
	if taskCount == 0 {
		return true
	}
	en.TaskData = make([]gmstruct.TaskData, taskCount)

	// 数据长度 < 任务变量数据长度
	len := uint32(binary.Size(en.TaskData))
	if dataLen < len {
		return false
	}

	structLen := uint32(binary.Size(en.TaskData[0]))
	start := uint32(0)
	end := start + structLen

	for i := uint32(0); i < taskCount; i++ {
		buf := bytes.NewBuffer(dataTmp[start:end])
		binary.Read(buf, binary.LittleEndian, &en.TaskData[i])
		start += structLen
		end += structLen
	}

	*current += len
	return true
}

func (en *RoleBakEncoder) decodeRoleItemData(data []byte, current *uint32) bool {
	dataTmp := data[*current:]
	counter := int16(0)
	var header gmstruct.DataHead

	// 角色身上没有物品，不解析
	if en.RoleData.ItemCount <= 0 {
		return true
	}
	en.ItemData = make([]gmstruct.ItemData, en.RoleData.ItemCount)

	end := uint32(0)
	start := uint32(0)
	structLen := uint32(0)

	for counter < en.RoleData.ItemCount {
		// 解析DataHead
		structLen = uint32(binary.Size(header))
		end = start + structLen
		buf := bytes.NewBuffer(dataTmp[start:end])
		binary.Read(buf, binary.LittleEndian, &header)
		start += structLen
		*current += structLen

		fmt.Printf("ItemCount = %d\n", en.RoleData.ItemCount)
		fmt.Printf("DataType = %d\n", header.DataType)
		fmt.Printf("DataCount = %d\n", header.DataCount)

		for i := int16(0); i < header.DataCount; i++ {
			en.ItemData[counter].HasStandard = (header.DataType&0xffff)&1 != 0
			en.ItemData[counter].HasLockSoul = (header.DataType&0xffff)&2 != 0
			en.ItemData[counter].HasBill = (header.DataType&0xffff)&4 != 0
			en.ItemData[counter].HasExtend = (header.DataType&0xffff)&8 != 0

			if en.ItemData[counter].HasStandard {
				structLen = uint32(binary.Size(en.ItemData[counter].Standard))
				end = start + structLen
				buf = bytes.NewBuffer(dataTmp[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].Standard)
				start += structLen
				*current += structLen
			}

			if en.ItemData[counter].HasLockSoul {
				structLen = uint32(binary.Size(en.ItemData[counter].LockSoul))
				end = start + structLen
				buf = bytes.NewBuffer(dataTmp[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].LockSoul)
				start += structLen
				*current += structLen
			}

			if en.ItemData[counter].HasBill {
				structLen = uint32(binary.Size(en.ItemData[counter].Bill))
				end = start + structLen
				buf = bytes.NewBuffer(dataTmp[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].Bill)
				start += structLen
				*current += structLen
			}

			if en.ItemData[counter].HasExtend {
				structLen = uint32(binary.Size(en.ItemData[counter].Extend))
				end = start + structLen
				buf = bytes.NewBuffer(dataTmp[start:end])
				binary.Read(buf, binary.LittleEndian, &en.ItemData[counter].Extend)
				start += structLen
				*current += structLen
			}

			counter++
		}
	}

	return true
}

func (en *RoleBakEncoder) getFSkillCount() (bool, uint32) {
	if en.RoleData.LSkillOffset < en.RoleData.FSkillOffset {
		return false, 0
	}

	var skill gmstruct.SkillData
	skillDataSize := uint32(binary.Size(skill))
	return true, (en.RoleData.LSkillOffset - en.RoleData.FSkillOffset) / skillDataSize
}

func (en *RoleBakEncoder) getLSkillCount() (bool, uint32) {
	if en.RoleData.TaskOffset < en.RoleData.LSkillOffset {
		return false, 0
	}

	var skill gmstruct.SkillData
	skillDataSize := uint32(binary.Size(skill))
	return true, (en.RoleData.TaskOffset - en.RoleData.LSkillOffset) / skillDataSize
}

func (en *RoleBakEncoder) getTaskCount() (bool, uint32) {
	if en.RoleData.ItemOffset < en.RoleData.TaskOffset {
		return false, 0
	}

	var task gmstruct.TaskData
	taskDataSize := uint32(binary.Size(task))
	return true, (en.RoleData.ItemOffset - en.RoleData.TaskOffset) / taskDataSize
}
