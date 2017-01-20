package gamestruct

import "fmt"

// RoleBaseInfo : 角色基础数据中的基础信息部分，包含角色最基本的属性状态
type RoleBaseInfo struct {
	RoleID              uint32   "当前未使用"
	RoleName            [32]byte "角色名"
	Sex                 byte     "性别：0-男性，1-女性"
	Alias               [32]byte "当前未使用"
	Account             [32]byte "帐号名"
	LastFaction         byte     "上次加入门派：-1(255)-未加入门派，0-少林，1-天王，2-唐门，3-五毒，4-峨嵋，5-翠烟，6-丐帮，7-天忍，8-武当，9-昆仑，10-华山"
	CurFaction          byte     "当前加入门派：-1(255)-未加入门派，0-少林，1-天王，2-唐门，3-五毒，4-峨嵋，5-翠烟，6-丐帮，7-天忍，8-武当，9-昆仑，10-华山"
	FightMode           byte     "战斗状态：0-非战斗状态，1-战斗状态"
	UseRevive           byte     "是否使用复活点复活"
	IsExchanged         byte     "是否处于跨服状态"
	PkStatus            byte     "PK状态"
	AddFactionTimes     int32    "加入门总次数"
	SectRole            int32    "未知含义"
	GroupCode           int32    "当前未使用"
	GroupRole           int32    "当前未使用"
	RevivalID           int32    "重生点地图ID"
	RevivalX            int32    "重生点ID"
	RevivalY            int32    "0"
	SubWorldID          int32    "上次登出地图ID"
	SubWorldMpsX        int32    "上次登出地图X坐标"
	SubWorldMpsY        int32    "上次登出地图Y坐标"
	PrimaryKey          [32]byte "角色唯一标识，MD5"
	BoxMoney            int32    "储物箱金钱"
	BagMoney            int32    "身上金钱"
	FiveElement         int32    "五行"
	Camp                int32    "阵营"
	RoleLevel           uint16   "角色等级"
	ExpHigh             int16    "经验高位"
	ExpLow              int32    "经验低位"
	LeadLevel           int32    "统帅力等级"
	LeadExp             int32    "统帅力经验"
	LiveExp             int32    "当前未使用"
	Strength            int32    "力量"
	Dexterity           int32    "身法"
	Vitality            int32    "外功"
	Energy              int32    "内功"
	Luck                int32    "幸运值"
	LifeMax             int32    "最大生命值"
	StaminaMax          int32    "最大体力值"
	ManaMax             int32    "最大内力值"
	CurLife             int32    "当前生命值"
	CurStamina          int32    "当前体力值"
	CurMana             int32    "当前内力值"
	PkValue             int32    "当前PK值"
	LeftPropPoint       int32    "剩余潜能点"
	LeftSkillPoint      int32    "剩余技能点"
	LeftLife            int32    "当前未使用"
	PlayGameTime        int32    "角色游戏时间"
	ArmorRes            int16    "当前未使用"
	Weaponres           int16    "当前未使用"
	HeadImage           int16    "头像编号"
	SectStat            int32    "未知含义"
	WorldStat           int32    "未知含义"
	KillPeopleNumber    int32    "未知含义"
	BitFlag             int32    "未知含义"
	TongID              uint32   "帮会ID"
	Repute              int32    "当前未使用"
	VotePoint           int32    "当前未使用"
	LastLogoutTime      uint32   "上次登出时间"
	PhysicsRes          int16    "当前未使用"
	ColdRes             int16    "当前未使用"
	PoisonRes           int16    "当前未使用"
	LightingRes         int16    "当前未使用"
	FireRes             int16    "当前未使用"
	ReLiveTime          int16    "当前未使用"
	ExtBox              byte     "扩展箱状态：0x01 box1; 0x04 box2; 0x10 box3"
	BoxPasswordParam    byte     "储物箱密码参数"
	Reserved13          byte     "当前未使用"
	Reserved14          byte     "当前未使用"
	BoxPassword         uint32   "储物箱密码"
	CatchTimeForAntiBot uint32   "使用外挂被抓时间"
	RefuseLoginCount    byte     "已拒绝使用外挂的角色登录的次数"
	HaveRefuseLogin     byte     "已拒绝状态中"
	IsExchangeServer    byte     "是否处于跨服务器中"
	RefuseLoginRe2      byte     "当前未使用"
	MapCopyIndex        int32    "当前未使用"
	RoleCreateTime      uint32   "角色创建时间"
	DataTransMark       byte     "数据转换标记"
	LastTransLifeLevel  byte     "上次转生等级"
	Reserved72          uint16   "当前未使用"
	ExtBuffOffset       uint32   "新扩充数据在RoleData的偏移"
	Reserved9           uint32   "当前未使用"
	Reserved0           uint32   "当前未使用"
}

// RoleBaseData : 角色基础数据，包含l"eBaseInfo和其他数据区块的偏移信息
type RoleBaseData struct {
	Version         uint32 "角色数据版本"
	RoleBaseInfo           // 角色基本数据：匿名字段，全部展开
	BaseNeedUpdate  byte   "通知是否需要更新"
	FightSkillCount int16  "战斗技能数量"
	LiveSkillCount  int16  "生活技能数量"
	TaskCount       byte   "该字段废弃"
	ItemCount       int16  "物品数量"
	StateCount      int16  "未知含义"
	TaskOffset      uint32 "任务变量数据偏移"
	LSkillOffset    uint32 "生活技能数据偏移"
	FSkillOffset    uint32 "战斗技能数据偏移"
	ItemOffset      uint32 "物品数据偏移"
	StateOffset     uint32 "未知含义"
	DataLen         uint32 "数据长度"
}

// SkillData : a data struct of skill data
type SkillData struct {
	SkillID  int16  // 技能ID
	SkillLv  int16  // 技能等级
	SkillExp uint32 // 技能经验
}

// TaskData : a data struct of task data
type TaskData struct {
	TaskID    int32
	TaskValue int32
}

// DataHead : a data struct of data header
type DataHead struct {
	DataType  int32 // 数据类型 : 高2字节-数据块大类型，低2字节-数据块小类型
	DataCount int16 // 此类型数据个数
	DataLen   int32 // 此数据结构字节数 : DataHead大小 + 此类型数据结构大小 * 此类型数据个数
}

// ItemDataStd : a data struct of item standard data
type ItemDataStd struct {
	ExParam1                  byte   // 物品扩展参数1
	ExParam2                  byte   // 物品扩展参数2
	ExParam3                  uint16 // 物品扩展参数3
	ClassCode                 int32  // 高4位 : 物品品质(Quality)，低4位 : 物品类型(Genre)
	Place                     int32  // 物品存储空间
	PosX                      byte   // 物品存储空间X坐标
	Feature1                  byte   // 换装外观字节1
	Reserved                  uint16 // 保留字段
	PosY                      byte   // 物品存储空间Y坐标
	Feature2                  byte   // 换装外观字节2
	Feature3                  byte   // 换装外观字节3
	Feature4                  byte   // 换装外观字节4
	GenTime                   int32  // 装备生成时间
	DetailType                int32  // Item(G, D, P)中的D
	ParticularType            int32  // Item(G, D, P)中的P
	Level                     byte   // 物品等级
	BindFlag                  byte   // 绑定标志 : 1-绑定中，0-解除绑定时间高位
	DeBindTime                uint16 // 解除绑定时间低位（离2000年1月1日的小时数）
	Series                    int32  // 五行
	Version                   int32  // 版本
	RandSeed                  int32  // 随机数种子
	Param2                    int32  // 物品扩展参数2
	Param3                    int32  // 物品扩展参数3
	Param5                    int32  // 物品扩展参数5
	Param4                    int32  // 物品扩展参数4
	Param6                    int32  // 物品扩展参数6
	Param1                    int32  // 物品扩展参数1
	Lucky                     int32  // 生成时角色的幸运值
	MaxDurability             int32  // 最大耐久度
	DurabilityOrLeftUsageTime int32  // 耐久度或剩余使用时间
}

// ItemDataLockSoul : a data struct of item lock soul data
type ItemDataLockSoul struct {
	Owner             [32]byte // 物品归属人
	State             byte     // 锁魂状态
	UnLockExpiredTime uint32   // 解魂到期时间
	ItemGUID          int64    // 物品GUID
	OwnerGUID         int64    // 归属人GUID
}

// ItemDataBill : a data struct of item bill data, if player buys item, the item will have this data
type ItemDataBill struct {
	ExpiredTime   uint32 //
	CurrencyType  uint16 //
	ComeFromPlace uint16 //
	GoodsPrice    int32  //
	ItemGUID      int64  // 物品GUID
}

// ItemDataExtend : a data struct of item extend data
type ItemDataExtend struct {
	FusionP         [6]uint16 // 熔炼的纹钢的P
	FusionMagicSeed [6]int32  // 熔炼的纹钢的魔法属性随机种子
	CurStarLevel    uint16    // 装备当前星级
	StarStoneP      [5]uint16 // 镶嵌的星辰石的P
	StarStoneLevel  [5]uint16 // 装备上对应镶孔的等级
	CurWishValue    uint16    // 装备当前幸运值
	LastBreakTime   uint32    // 装备上次突破时间
	OwnerName       [32]byte  // 装备所有者名字-十万VIP奖励
	Reserved        [4]byte   // 预留空间
}

// ItemData : a data struct of item data
type ItemData struct {
	HasStandard bool             // 是否有标准数据
	Standard    ItemDataStd      // 标准数据
	HasLockSoul bool             // 是否有锁魂数据
	LockSoul    ItemDataLockSoul // 锁魂数据
	HasBill     bool             // 是否有账单数据
	Bill        ItemDataBill     // 账单数据
	HasExtend   bool             // 是否有扩展数据
	Extend      ItemDataExtend   // 扩展数据
}

// enum state type
const (
	NoneStateType = iota
	SkillStateType
	SkillCDType
	FeatureInfoType
	PlayerEventInfoType
	PlayerTitleType
	PlayerMaxSkillLevelType
	CustomStructType
)

// SkillState :
type SkillState struct {
	StateSkillID   int32
	StateType      int32
	StateLevel     int32
	Time           uint32
	NoClearOnDeath byte
	Reserved2      byte
	Reserved3      byte
	Reserved4      byte
}

// Partner :
type Partner struct {
	TemplateID int32
	Series     byte
	Level      byte
	CurLife    int32
	MapX       int32
	MapY       int32
}

// FeatureInfo :
type FeatureInfo struct {
	FeaturePriority   byte
	FeatureState      byte
	AvailableTimeType byte
	AvailableTime     uint32
	NpcSettingIdx     uint32
	HelmType          uint16
	ArmorType         uint16
	WeaponType        uint16
	HorseType         uint16
}

// PlayerEventData :
type PlayerEventData struct {
	Data uint16
	ID   uint16
}

// PlayerEvent :
type PlayerEvent struct {
	EventCount uint16
	Reserved   uint16
	EventData  [7]PlayerEventData
}

// 角色称号类型
const (
	RoleTitleTypeOfNormal   = iota // 普通称号，存档不记录
	RoleTitleTypeOfGameTime        // 游戏帧数称号，称号存在截止时间以玩家实际游戏时间为准，存档
	RoleTitleTypeOfTrueTime        // 实际时间称号，称号存在存在截止时间以真实时间为准，存档
	RoleTitleTypeCount             // 角色称号类型总数
)

// RoleTitleTime : 角色称号时间
type RoleTitleTime struct {
	Type     int32 // 称号时间类型，跟称号类型定义一致
	Time     int32 // 称号时间，剩余游戏帧数
	TrueTime int32 // 称号时间，剩余时间秒数
}

// RoleTitle : 角色称号数据
type RoleTitle struct {
	TitleTime       RoleTitleTime // 角色称号时间
	TitleID         uint32        // 角色称号ID
	IsActiveTitleID byte          // 是否当前激活的称号ID
}

// MaxSkillLevelInfoItem :
type MaxSkillLevelInfoItem struct {
	SkillID       uint16
	SkillMaxLevel byte
	Reserved      byte
}

// MaxSkillLevelInfo :
type MaxSkillLevelInfo struct {
	Data [7]MaxSkillLevelInfoItem
}

// enum custom struct type
const (
	CustomDataTypeOfPartner = iota
	CustomStructTypeCount
)

// CustomDataHeader : 角色状态数据中的自定义类型数据
type CustomDataHeader struct {
	Type     byte   // 自定义状态数据类型
	Size     uint32 // 自定义状态数据类型大小，size = header大小 + 数据大小
	Reserved [11]byte
}

// CustomDataOfPartnerHeader : 角色状态自定义数据的同伴数据的数据头
type CustomDataOfPartnerHeader struct {
	CustomDataHeader
	CurPartnerIDX          byte // 当前选择的同伴在管理同伴数组中的索引
	IsCurPartnerCalledOut  byte // 当前选择的同伴是否召唤状态
	IsCurPartnerFollowOnly byte // 当前选择的同伴是否只是跟随状态
	PartnerCount           byte // 管理同伴数组大小
}

// SkillCDData :
type SkillCDData struct {
	SkillID     int32
	DelaytFrame int32
}

// SkillCD :
type SkillCD struct {
	LastTime int32
	Data     [3]SkillCDData
	Reserved int32
}

// StateData :
type StateData struct {
	Type byte
	Data [32]byte
}

// RoleExtDataOfBase ：
type RoleExtDataOfBase struct {
	RoleNameGUID        int64  "角色GUID"
	Password            uint32 "绑定或锁魂操作的密码（两者共一个密码）"
	PasswordExpiredTime uint32 "绑定或锁魂操作解除密码日期"
	PasswordTimeOrTimes uint32 "输入密码错误次数或密码错误解禁日期"
	HavePassword        byte   "绑定或锁魂是否设置了密码"
}

// RoleExtDataOfLingLongLockParam ;
type RoleExtDataOfLingLongLockParam struct {
	CardHash uint32 "网卡硬件地址哈希码"
	DiskHash uint32 "硬盘序列号哈希码"
}

// RoleExtDataOfLingLongLock :
type RoleExtDataOfLingLongLock struct {
	RoleExtDataOfLingLongLockParam
	Password uint32 "玲珑锁密码"
	Timeout  uint32 "到期时间，非0表示玲珑锁有效"
	Locked   byte   "是否锁定（跨服时此状态保持不变，登录时重新设置）"
}

// RoleExtDataOfHangerOnData :
type RoleExtDataOfHangerOnData struct {
	CurTaskType     byte  // 门客当前任务类型
	CurTaskNum      byte  // 门客当前任务数量
	CurTaskRestTime int32 // 门客当前任务剩余时间(游戏逻辑帧数,0表示任务完成可领奖,-1表示当前无任务)
	ExpiredTime     int32
}

// RoleExtDataOfHangerOn :
type RoleExtDataOfHangerOn struct {
	PermanentHangerOn RoleExtDataOfHangerOnData     // 永久门客
	TemporaryHangerOn [10]RoleExtDataOfHangerOnData // 临时门客
}

// RoleExtDataOfTransNimbus :
type RoleExtDataOfTransNimbus struct {
	TransNimbusExpHigh int16 "转灵经验(高位)"
	TransNimbusExpLow  int32 "转灵经验(低位)"
}

// RoleExtDataOfBreak ：
type RoleExtDataOfBreak struct {
	HasBreak byte "200突破，0表示没突破"
}

// RoleExtDataOfEquipCompose ：
type RoleExtDataOfEquipCompose struct {
	ComposeLv    uint32 "合成等级"
	ComposeExp   uint32 "合成经验"
	DecomposeLv  uint32 "分解等级"
	DecomposeExp uint32 "分解经验"
}

// RoleExtData :
type RoleExtData struct {
	HasBase         bool
	Base            RoleExtDataOfBase
	HasLingLongLock bool
	LingLongLock    RoleExtDataOfLingLongLock
	HasHangerOn     bool
	HangerOn        RoleExtDataOfHangerOn
	HasTransNimbus  bool
	TransNimbus     RoleExtDataOfTransNimbus
	HasBreak        bool
	Break           RoleExtDataOfBreak
	HasEquipCompose bool
	EquipCompose    RoleExtDataOfEquipCompose
}

func (r RoleExtData) PrintEquipComposeData() {
	fmt.Println("===============================[EQUIP COMPOSE]================================")
	fmt.Printf("Compose Level : %d\t Decompose Level : %d\n", r.EquipCompose.ComposeLv, r.EquipCompose.DecomposeLv)
	fmt.Printf("Compose exp   : %d\t Decompose exp   : %d\n", r.EquipCompose.ComposeExp, r.EquipCompose.DecomposeExp)
}
