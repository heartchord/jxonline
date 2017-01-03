package gameencoder

import (
	"bytes"
	"encoding/binary"

	"fmt"

	gmstruct "github.com/heartchord/jxonline/gamestruct"
	"github.com/henrylee2cn/mahonia"
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
	Header RoleBakHeader // Bak数据头
	Body   RoleBakBody   // Bak数据体
}

// RoleBakEncoder : a data struct of role bak encoder and decoder
type RoleBakEncoder struct {
	BakData   RoleBakData
	RoleData  gmstruct.RoleData
	SkillData []gmstruct.SkillData
}

// Decode : function to decode original role bak data
func (pEncoder *RoleBakEncoder) Decode(data []byte) bool {
	if !pEncoder.decodeBakHeader(data) {
		return false
	}

	current := uint32(0)
	if !pEncoder.decodeRoleBaseInfo(pEncoder.BakData.Body.RoleData, &current) {
		return false
	}

	pEncoder.decodeRoleSkillData(pEncoder.BakData.Body.RoleData, &current)

	return true
}

func (pEncoder *RoleBakEncoder) decodeBakHeader(data []byte) bool {
	dataLen := uint32(len(data))
	current := uint32(0)

	if dataLen <= 4 { // 数据长度 <= 角色名数据头长度
		return false
	}

	// 获取角色名长度
	tmplen := uint32(0)
	tmpbuf := bytes.NewBuffer(data[current:4]) // [0, 3]存储角色名长度
	binary.Read(tmpbuf, binary.LittleEndian, &tmplen)
	pEncoder.BakData.Header.RoleNameLen = tmplen
	current += 4

	// 打印调试信息
	fmt.Printf("Header.RoleNameLen = %d\n", pEncoder.BakData.Header.RoleNameLen)

	// 获取角色名
	n := 4 + pEncoder.BakData.Header.RoleNameLen
	if pEncoder.BakData.Header.RoleNameLen <= 0 || dataLen <= n { // 角色名长度 <= 0 或 数据长度 <= 角色名数据头长度 + 角色名长度
		return false
	}
	pEncoder.BakData.Header.RoleNameGBK = data[current:n] // [4, 4 + namelen]存储角色名
	current += pEncoder.BakData.Header.RoleNameLen

	// 打印调试信息
	mdecoder := mahonia.NewDecoder("GBK")
	rolename := string(pEncoder.BakData.Header.RoleNameGBK[:])
	rolename = mdecoder.ConvertString(rolename)
	fmt.Printf("Header.RoleName = %s\n", rolename)

	// 获取角色原始数据长度
	n = 4 + pEncoder.BakData.Header.RoleNameLen + 4
	if dataLen <= n { // 数据长度 <= 角色名数据头长度 + 角色名长度 + 角色数据长度
		return false
	}
	tmpbuf = bytes.NewBuffer(data[current:n])
	binary.Read(tmpbuf, binary.LittleEndian, &tmplen)
	pEncoder.BakData.Header.RoleDataLen = tmplen // [4 + namelen, 4 + namelen + 4]存储角色原始数据长度
	current += 4

	// 打印调试信息
	fmt.Printf("Header.RoleOriginDataLen = %d\n", pEncoder.BakData.Header.RoleDataLen)

	// 获取角色原始数据
	n = 4 + pEncoder.BakData.Header.RoleNameLen + 4 + pEncoder.BakData.Header.RoleDataLen
	if dataLen < n {
		return false
	}
	pEncoder.BakData.Body.RoleData = data[current:n]

	// 打印调试信息
	fmt.Printf("len(Body.Data) = %d\n", len(pEncoder.BakData.Body.RoleData))
	return true
}

func (pEncoder *RoleBakEncoder) decodeRoleBaseInfo(data []byte, current *uint32) bool {
	dataTmp := data[*current:]
	dataLen := uint32(len(dataTmp))

	// 数据长度 < 角色名数据头长度
	len := uint32(binary.Size(pEncoder.RoleData))

	if dataLen < len {
		return false
	}

	buf := bytes.NewBuffer(dataTmp[:len])
	binary.Read(buf, binary.LittleEndian, &pEncoder.RoleData)
	*current += len

	return true
}

func (pEncoder *RoleBakEncoder) decodeRoleSkillData(data []byte, current *uint32) bool {
	dataTmp := data[*current:]
	dataLen := uint32(len(dataTmp))

	ret, skillCount := pEncoder.getFightSkillCount()
	if !ret {
		return false
	}
	pEncoder.SkillData = make([]gmstruct.SkillData, skillCount)

	// 数据长度 < 技能数据长度
	len := uint32(binary.Size(pEncoder.SkillData))
	if dataLen < len {
		return false
	}

	structLen := uint32(binary.Size(pEncoder.SkillData[0]))
	start := uint32(0)
	end := start + structLen

	for i := uint32(0); i < skillCount; i++ {
		buf := bytes.NewBuffer(dataTmp[start:end])
		binary.Read(buf, binary.LittleEndian, &pEncoder.SkillData[i])
		start += structLen
		end += structLen
	}

	*current += len
	return true
}

func (pEncoder *RoleBakEncoder) getFightSkillCount() (bool, uint32) {
	if pEncoder.RoleData.LSkillOffset < pEncoder.RoleData.FSkillOffset {
		return false, 0
	}

	var skill gmstruct.SkillData
	skillDataSize := uint32(binary.Size(skill))
	return true, (pEncoder.RoleData.LSkillOffset - pEncoder.RoleData.FSkillOffset) / skillDataSize
}
