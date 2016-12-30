package gameencoder

import (
	"bytes"
	"encoding/binary"

	gmstruct "github.com/heartchord/jxonline/gamestruct"
)

// RoleBakData : a data struct of role bak data
type RoleBakData struct {
	RoleNameLen    uint32 // 角色名长度
	RoleName       []byte // 角色名
	RoleOriginData []byte // 角色原始二进制数据
}

// RoleBakEncoder :
type RoleBakEncoder struct {
	BakData  RoleBakData
	RoleData gmstruct.RoleData
}

// Decode :
func (pEncoder *RoleBakEncoder) Decode(data []byte) bool {
	dataLen := uint32(len(data))

	// 数据长度 < 角色名数据头长度
	if dataLen < 4 {
		return false
	}

	// 获取角色名长度
	len := uint32(0)
	buf := bytes.NewBuffer(data[:4])
	binary.Read(buf, binary.LittleEndian, &len)
	pEncoder.BakData.RoleNameLen = len

	// 角色名长度 < 0 或 数据长度 <= 角色名数据头长度 + 角色名长度
	if len <= 0 || dataLen <= 4+len {
		return false
	}
	pEncoder.BakData.RoleName = data[4 : 4+len]
	pEncoder.BakData.RoleOriginData = data[4+len:]

	// 这里跳过了4个字节数据，含义待查明
	if !pEncoder.decodeRoleBaseInfo(pEncoder.BakData.RoleOriginData[4:]) {
		return false
	}

	return true
}

func (pEncoder *RoleBakEncoder) decodeRoleBaseInfo(data []byte) bool {
	dataLen := uint32(len(data))

	// 数据长度 < 角色名数据头长度
	len := uint32(binary.Size(pEncoder.RoleData))
	if dataLen < len {
		return false
	}

	buf := bytes.NewBuffer(data[:len])
	binary.Read(buf, binary.LittleEndian, &pEncoder.RoleData)

	return true
}
