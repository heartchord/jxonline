package gamestruct

// RoleBaseData : a data struct of role base info
type RoleBaseData struct {
	RoleID              uint32   // 当前未使用
	RoleName            [32]byte // 角色名
	Sex                 byte     // 性别
	Alias               [32]byte // 当前未使用
	Account             [32]byte // 帐号名
	LastFaction         byte     // 上次加入门派
	CurFaction          byte     // 当前门派
	FightMode           byte     // 战斗状态
	UseRevive           byte     // 是否使用复活点复活
	IsExchanged         byte     // 是否处于跨服状态
	PkStatus            byte     // PK状态
	AddFactionTimes     int32    // 加入门总次数
	SectRole            int32    // 未知含义
	GroupCode           int32    // 当前未使用
	GroupRole           int32    // 当前未使用
	RevivalID           int32    // 重生点地图ID
	RevivalX            int32    // 重生点ID
	RevivalY            int32    // 0
	SubWorldID          int32    // 上次登出地图ID
	SubWorldMpsX        int32    // 上次登出地图X坐标
	SubWorldMpsY        int32    // 上次登出地图Y坐标
	PrimaryKey          [32]byte // 角色唯一标识，MD5
	BoxMoney            int32    // 储物箱金钱
	BagMoney            int32    // 身上金钱
	FiveElement         int32    // 五行
	Camp                int32    // 阵营
	RoleLevel           uint16   // 角色等级
	ExpHigh             int16    // 经验高位
	ExpLow              int32    // 经验低位
	LeadLevel           int32    // 统帅力等级
	LeadExp             int32    // 统帅力经验
	LiveExp             int32    // 当前未使用
	Strength            int32    // 力量
	Dexterity           int32    // 身法
	Vitality            int32    // 外功
	Energy              int32    // 内功
	Luck                int32    // 幸运值
	LifeMax             int32    // 最大生命
	StaminaMax          int32    // 最大体力
	ManaMax             int32    // 最大内力
	CurLife             int32    // 当前生命
	CurStamina          int32    // 当前体力
	CurMana             int32    // 当前内力
	PkValue             int32    // 当前PK值
	LeftPropPoint       int32    // 潜能点
	LeftSkillPoint      int32    // 技能点
	LeftLife            int32    // 当前未使用
	PlayGameTime        int32    // 角色游戏时间
	ArmorRes            int16    // 当前未使用
	Weaponres           int16    // 当前未使用
	HeadImage           int16    // 头像编号
	SectStat            int32    // 未知含义
	WorldStat           int32    // 未知含义
	KillPeopleNumber    int32    // 未知含义
	BitFlag             int32    // 未知含义
	TongID              uint32   // 帮会ID
	Repute              int32    // 当前未使用
	VotePoint           int32    // 当前未使用
	LastLogoutTime      uint32   // 上次登出时间
	PhysicsRes          int16    // 当前未使用
	ColdRes             int16    // 当前未使用
	PoisonRes           int16    // 当前未使用
	LightingRes         int16    // 当前未使用
	FireRes             int16    // 当前未使用
	ReLiveTime          int16    // 当前未使用
	ExtBox              byte     // 扩展箱状态：0x01 box1; 0x04 box2; 0x10 box3
	BoxPasswordParam    byte     // 储物箱密码参数
	Reserve13           byte     // 当前未使用
	Reserve14           byte     // 当前未使用
	BoxPassword         uint32   // 储物箱密码
	CatchTimeForAntiBot uint32   // 使用外挂被抓时间
	RefuseLoginCount    byte     // 已拒绝使用外挂的角色登录的次数
	HaveRefuseLogin     byte     // 已拒绝状态中
	IsExchangeServer    byte     // 是否处于跨服务器中
	RefuseLoginRe2      byte     // 当前未使用
	MapCopyIndex        int32    // 当前未使用
	RoleCreateTime      uint32   // 角色创建时间
	DataTransMark       byte     // 数据转换标记
	LastTransLifeLevel  byte     // 上次转生等级
	Reserve72           uint16   // 当前未使用
	ExBuffOffset        uint32   // 新扩充数据在RoleData的偏移
	Reserve9            uint32   // 当前未使用
	Reserve0            uint32   // 当前未使用
}

// RoleData : a data struct of role data
type RoleData struct {
	Version         uint32       // 角色数据版本
	BaseData        RoleBaseData // 角色基本数据
	BaseNeedUpdate  byte         // 通知是否需要更新
	FightSkillCount int16        // 战斗技能数量
	LiveSkillCount  int16        // 生活技能数量
	TaskCount       byte         // 任务变量数量
	ItemCount       int16        // 物品数量
	StateCount      int16        // 未知含义
	TaskOffset      uint32       // 任务变量数据偏移
	LSkillOffset    uint32       // 生活技能数据偏移
	FSkillOffset    uint32       // 战斗技能数据偏移
	ItemOffset      uint32       // 物品数据偏移
	StateOffset     uint32       // 未知含义
	DataLen         uint32       // 数据长度
}

// SkillData :
type SkillData struct {
	SkillID  int16  // 技能ID
	SkillLv  int16  // 技能等级
	SkillExp uint32 // 技能经验
}
