package gameencoder

import (
	"bytes"
	"encoding/binary"

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
}

// Decode : function to decode original role bak data
func (en *RoleBakEncoder) Decode(data []byte) bool {
	current := uint32(0)

	if !en.decodeBakHeader(data) {
		return false
	}

	if !en.decodeRoleBaseInfo(en.BakData.RoleData, &current) {
		return false
	}

	if !en.decodeRoleFSkillData(en.BakData.RoleData, &current) {
		return false
	}

	if !en.decodeRoleLSkillData(en.BakData.RoleData, &current) {
		return false
	}

	if !en.decodeRoleTaskData(en.BakData.RoleData, &current) {
		return false
	}

	return true
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
