package gameencoder

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// LogWriter :
type LogWriter func(format string, a ...interface{}) (n int, err error)

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
	BakData     RoleBakData // 原始角色数据
	RoleEncoder             // 角色解码器
	logger      LogWriter   // 日志函数
}

// NewRoleBakEncoder : 123
func NewRoleBakEncoder() (en *RoleBakEncoder) {
	en = new(RoleBakEncoder)
	en.logger = fmt.Printf

	// 初始化ReadFunction
	en.RoleEncoder.Init()
	en.RoleEncoder.SetLogger(fmt.Printf)
	return en
}

// SetLogger :
func (en *RoleBakEncoder) SetLogger(logger LogWriter) {
	en.logger = logger
	en.RoleEncoder.SetLogger(logger)
}

// Decode : function to decode original role bak data
func (en *RoleBakEncoder) Decode(data []byte) bool {

	if !en.decodeBakHeader(data) {
		return false
	}

	if !en.RoleEncoder.Decode(en.BakData.RoleData) {
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

	// 获取角色名长度(包含'\0'结束符)
	tmplen := uint32(0)
	tmpbuf := bytes.NewBuffer(data[current:4]) // [0, 3]存储角色名长度
	binary.Read(tmpbuf, binary.LittleEndian, &tmplen)
	en.BakData.RoleNameLen = tmplen
	current += 4

	// 获取角色名
	n := 4 + en.BakData.RoleNameLen - 1              // 要去掉'\0'字符
	if en.BakData.RoleNameLen <= 0 || dataLen <= n { // 角色名长度 <= 0 或 数据长度 <= 角色名数据头长度 + 角色名长度
		return false
	}

	en.BakData.RoleNameGBK = data[current:n] // [4, 4 + namelen]存储角色名
	current += en.BakData.RoleNameLen

	// 获取角色原始数据长度
	n = 4 + en.BakData.RoleNameLen + 4
	if dataLen <= n { // 数据长度 <= 角色名数据头长度 + 角色名长度 + 角色数据长度
		return false
	}
	tmpbuf = bytes.NewBuffer(data[current:n])
	binary.Read(tmpbuf, binary.LittleEndian, &tmplen)
	en.BakData.RoleDataLen = tmplen // [4 + namelen, 4 + namelen + 4]存储角色原始数据长度
	current += 4

	// 获取角色原始数据
	n = 4 + en.BakData.RoleNameLen + 4 + en.BakData.RoleDataLen
	if dataLen < n {
		return false
	}
	en.BakData.RoleData = data[current:n]

	return true
}
