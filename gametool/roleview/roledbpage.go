package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/heartchord/jxonline/gameencoder"
	"github.com/henrylee2cn/mahonia"
	"github.com/lxn/walk"
	dcl "github.com/lxn/walk/declarative"
)

// RoleDBData :
type RoleDBData struct {
	ID           string
	RoleName     string
	Account      string
	RoleData     []byte
	LastModified string
}

// DBInfoBindData :
type DBInfoBindData struct {
	IP       string
	Port     string
	Username string
	Password string
	DBName   string
	Charset  string
	Rolename string
}

// RoleDbPage :
type RoleDbPage struct {
	*walk.TabPage
	dbInfoBindData          *DBInfoBindData
	roleBaseDataModel       *DataModel1
	roleBaseDataTV          *walk.TableView
	dbIPText                *walk.LineEdit
	dbPortText              *walk.LineEdit
	dbUsernameText          *walk.LineEdit
	dbPasswordText          *walk.LineEdit
	dbNameText              *walk.LineEdit
	dbCharsetText           *walk.LineEdit
	dbRolenameText          *walk.LineEdit
	dbQueryIDText           *walk.LineEdit
	dbQueryUserNameText     *walk.LineEdit
	dbQueryAccountText      *walk.LineEdit
	dbQueryLastModifiedText *walk.LineEdit
	dbProcessLogText        *walk.TextEdit
	roleExtDataDlg          *RoleExtDataDialog
	roleSkillDlg            *RoleSkillDialog
	roleTaskDlg             *RoleTaskDialog
	encoder                 *gameencoder.RoleEncoder
	decodeProcessFinished   bool
}

// Create is
func (pg *RoleDbPage) Create() *dcl.TabPage {

	pg.roleBaseDataModel = NewDataModel1("../../gameresource/img/right-arrow2.ico")
	pg.dbInfoBindData = new(DBInfoBindData)
	pg.decodeProcessFinished = true

	pg.roleExtDataDlg = new(RoleExtDataDialog)
	pg.roleSkillDlg = new(RoleSkillDialog)
	pg.roleTaskDlg = new(RoleTaskDialog)

	pg.encoder = new(gameencoder.RoleEncoder)
	pg.encoder.Init()
	pg.encoder.SetLogger(pg.WriteLog)

	return pg.createPage()
}

// SetPageDefaultSettings :
func (pg *RoleDbPage) SetPageDefaultSettings() {
	pg.dbIPText.SetText("127.0.0.1")
	pg.dbPortText.SetText("3306")
	pg.dbUsernameText.SetText("root")
	pg.dbPasswordText.SetText("liuyubin")
	pg.dbNameText.SetText("jxib")
	pg.dbCharsetText.SetText("gbk")
	pg.dbRolenameText.SetText("无毒教教")
}

func (pg *RoleDbPage) onQueryRoleDBData() {
	if !pg.decodeProcessFinished {
		pg.WriteLog("Info - Last decode process hasn't finished, please wait...")
		return
	}

	pg.decodeProcessFinished = false
	pg.dbProcessLogText.SetText("")
	go pg.queryRoleDBDataRoutineFunction()
}

func (pg *RoleDbPage) queryRoleDBDataRoutineFunction() {
	defer func() {
		pg.decodeProcessFinished = true
	}()

	var data RoleDBData

	mencoder := mahonia.NewEncoder("GBK")
	mdecoder := mahonia.NewDecoder("GBK")
	//fmt.Printf("IP = %s\n", pg.dbInfoBindData.IP)
	//fmt.Printf("Port = %s\n", pg.dbInfoBindData.Port)
	//fmt.Printf("Username = %s\n", pg.dbInfoBindData.Username)
	//fmt.Printf("Password = %s\n", pg.dbInfoBindData.Password)
	//fmt.Printf("DBName = %s\n", pg.dbInfoBindData.DBName)
	//fmt.Printf("Charset = %s\n", pg.dbInfoBindData.Charset)
	//fmt.Printf("Rolename = %s\n", pg.dbInfoBindData.Rolename)
	// 连接数据库
	openString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		pg.dbInfoBindData.Username, pg.dbInfoBindData.Password,
		pg.dbInfoBindData.IP, pg.dbInfoBindData.Port,
		pg.dbInfoBindData.DBName, pg.dbInfoBindData.Charset)
	pg.WriteLog("开始连接数据库")
	pg.WriteLog(">> Info : %s", openString)

	var db *sql.DB
	var err error
	db, err = sql.Open("mysql", openString)
	if err != nil {
		pg.WriteLog(">> Error : %s", err.Error())
		return
	}
	pg.WriteLog(">> Info : 连接数据库成功")

	defer func() {
		// 关闭数据库
		pg.WriteLog(">> Info : 开始关闭数据库")
		err = db.Close()
		if err != nil {
			pg.WriteLog(">> Error : %s", err.Error())
			return
		}
		pg.WriteLog(">> Info : 关闭数据库成功")
	}()

	// 查询数据
	pg.WriteLog("开始查询角色数据")

	name, ok := mencoder.ConvertStringOK(pg.dbInfoBindData.Rolename)
	if !ok {
		pg.WriteLog(">> Error : 角色名转换编码[UTF8 -> GBK]失败")
		return
	}

	sqlString := fmt.Sprintf("select * from role where RoleName = '%s'", name)
	pg.WriteLog(">> Info : %s", sqlString)

	var rows *sql.Rows
	rows, err = db.Query(sqlString)
	if err != nil {
		pg.WriteLog(">> Error : %s", err.Error())
		return
	}

	defer func() {
		err = rows.Close()
		if err != nil {
			pg.WriteLog(">> Error : %s", err.Error())
			return
		}
	}()

	var fields []string
	fields, err = rows.Columns()
	if err != nil {
		return
	}
	if len(fields) <= 0 {
		pg.WriteLog(">> Info : 未查询到玩家[%s]的角色数据!", pg.dbInfoBindData.Rolename)
		return
	}

	pg.WriteLog("查询角色数据成功")

	for rows.Next() {
		var id int32
		var username string
		var account string
		var datetime time.Time
		var roledata []byte

		err = rows.Scan(&id, &username, &account, &roledata, &datetime)
		if err != nil {
			pg.WriteLog(">> Error : %s", err.Error())
			return
		}

		data.ID = fmt.Sprintf("%d", id)
		data.RoleName = mdecoder.ConvertString(username)
		data.Account = mdecoder.ConvertString(account)
		data.RoleData = roledata
		data.LastModified = datetime.Format("2006-01-02 15:04:05")
	}

	pg.dbQueryIDText.SetText(data.ID)
	pg.dbQueryUserNameText.SetText(data.RoleName)
	pg.dbQueryAccountText.SetText(data.Account)
	pg.dbQueryLastModifiedText.SetText(data.LastModified)

	pg.WriteLog("开始解析角色数据")
	pg.encoder.Decode(data.RoleData)

	pg.roleBaseDataModel.ResetRows(pg.encoder.RoleBaseData)
}

func (pg *RoleDbPage) onShowRoleSkillDialog() {
	if !pg.roleSkillDlg.CreateInstance(mw) {
		return
	}

	if pg.encoder.FSkillData != nil {
		pg.roleSkillDlg.RoleFSkillDataModel.ResetRows(pg.encoder.FSkillData)
		pg.roleSkillDlg.RoleLSkillDataModel.ResetRows(pg.encoder.LSkillData)
	}

	pg.roleSkillDlg.Run()
}

func (pg *RoleDbPage) onShowRoleTaskDialog() {
	if !pg.roleTaskDlg.CreateInstance(mw) {
		return
	}

	if pg.encoder.TaskData != nil {
		pg.roleTaskDlg.RoleTaskDataModel.ResetRows(pg.encoder.TaskData)
	}

	pg.roleTaskDlg.Run()
}

func (pg *RoleDbPage) onShowRoleExtDataDialog() {
	if !pg.roleExtDataDlg.CreateInstance(mw) {
		return
	}

	if pg.encoder.RoleExtData.HasBase {
		pg.roleExtDataDlg.LockSoulDataModel.ResetRows(pg.encoder.RoleExtData.Base)
	} else {
		pg.roleExtDataDlg.LockSoulDataModel.ResetRows(nil)
	}

	if pg.encoder.RoleExtData.HasBreak {
		pg.roleExtDataDlg.RoleBreakDataModel.ResetRows(pg.encoder.RoleExtData.Break)
	} else {
		pg.roleExtDataDlg.RoleBreakDataModel.ResetRows(nil)
	}

	if pg.encoder.RoleExtData.HasTransNimbus {
		pg.roleExtDataDlg.TransNimbusDataModel.ResetRows(pg.encoder.RoleExtData.TransNimbus)
	} else {
		pg.roleExtDataDlg.TransNimbusDataModel.ResetRows(nil)
	}

	if pg.encoder.RoleExtData.HasLingLongLock {
		pg.roleExtDataDlg.LingLongLockDataModel.ResetRows(pg.encoder.RoleExtData.LingLongLock)
	} else {
		pg.roleExtDataDlg.LingLongLockDataModel.ResetRows(nil)
	}

	if pg.encoder.RoleExtData.HasEquipCompose {
		pg.roleExtDataDlg.EquipComposeDataModel.ResetRows(pg.encoder.RoleExtData.EquipCompose)
	} else {
		pg.roleExtDataDlg.EquipComposeDataModel.ResetRows(nil)
	}

	pg.roleExtDataDlg.Run()
}

// WriteLog :
func (pg *RoleDbPage) WriteLog(format string, a ...interface{}) (n int, err error) {
	ts := time.Now().Unix()
	tm := time.Unix(ts, 0)
	t := tm.Format("2006-01-02 15:04:05")

	format = t + " : " + format + "\r\n"
	log := fmt.Sprintf(format, a...)
	pg.dbProcessLogText.AppendText(log)
	return pg.dbProcessLogText.TextLength(), nil
}

func (pg *RoleDbPage) createPage() *dcl.TabPage {
	var ep walk.ErrorPresenter

	return &dcl.TabPage{
		AssignTo: &pg.TabPage,
		Title:    "Role DB",
		Layout:   dcl.HBox{},
		DataBinder: dcl.DataBinder{
			DataSource:     pg.dbInfoBindData,
			AutoSubmit:     true,
			ErrorPresenter: dcl.ErrorPresenterRef{ErrorPresenter: &ep},
		},
		Children: []dcl.Widget{
			dcl.HSplitter{
				Children: []dcl.Widget{
					dcl.Label{
						ColumnSpan: 1,
						Text:       "Bak文件路径：",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
					},
					dcl.Label{
						ColumnSpan: 1,
						Text:       "Bak文件路径：",
						Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
					},
				},
			},
			dcl.Composite{ // 这里重新布局
				MinSize: dcl.Size{Width: 500, Height: 0},
				Font:    dcl.Font{Family: "微软雅黑", PointSize: 10},
				Layout:  dcl.Grid{Columns: 1},
				Children: []dcl.Widget{

					dcl.Composite{ // 这里重新布局
						Layout: dcl.Grid{Columns: 8},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 1,
								Text:       "IP地址：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbIPText,
								ColumnSpan: 1,
								Text:       dcl.Bind("IP"),
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "端口：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbPortText,
								ColumnSpan: 1,
								Text:       dcl.Bind("Port"),
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "用户名：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbUsernameText,
								ColumnSpan: 1,
								Text:       dcl.Bind("Username"),
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "密码：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbPasswordText,
								ColumnSpan: 1,
								Text:       dcl.Bind("Password"),
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "数据库名：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbNameText,
								ColumnSpan: 1,
								Text:       dcl.Bind("DBName"),
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "字符集：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbCharsetText,
								ColumnSpan: 1,
								Text:       dcl.Bind("Charset"),
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色名：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbRolenameText,
								ColumnSpan: 1,
								Text:       dcl.Bind("Rolename"),
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.HSpacer{
								ColumnSpan: 1,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "查询",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
								OnClicked:  pg.onQueryRoleDBData,
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						Layout: dcl.HBox{},
						Children: []dcl.Widget{
							dcl.TextEdit{
								AssignTo:   &pg.dbProcessLogText,
								ColumnSpan: 8,
								MinSize:    dcl.Size{Width: 100, Height: 100},
								Text:       "",
								ReadOnly:   true,
								OnSizeChanged: func() {
								},
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						Font:   dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout: dcl.Grid{Columns: 8, Spacing: 1},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 1,
								Text:       "ID：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbQueryIDText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "角色：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbQueryUserNameText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "帐号：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								AssignTo:   &pg.dbQueryAccountText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
								MinSize:    dcl.Size{Width: 100, Height: 0},
							},
							dcl.Label{
								ColumnSpan: 1,
								Text:       "上次存档时间：",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10, Bold: true},
							},
							dcl.LineEdit{
								MinSize:    dcl.Size{Width: 100, Height: 0},
								AssignTo:   &pg.dbQueryLastModifiedText,
								ColumnSpan: 1,
								Text:       "",
								ReadOnly:   true,
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						MinSize: dcl.Size{Width: 0, Height: 400},
						Font:    dcl.Font{Family: "微软雅黑", PointSize: 10},
						Layout:  dcl.Grid{Columns: 1, Spacing: 10},
						Children: []dcl.Widget{
							dcl.Label{
								ColumnSpan: 1,
								Text:       "【角色基础数据信息】",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 11, Bold: true},
							},
							dcl.Composite{ // 这里重新布局
								//AssignTo:   &pg.roleBaseDataComposite,
								ColumnSpan: 1,
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 10},
								Layout:     dcl.Grid{Columns: 1, Spacing: 10},
								//ContextMenuItems: []dcl.MenuItem{
								//	dcl.Action{
								//		Text:        "时间戳转换",
								//		OnTriggered: pg.timeStamp2timeActionHandler,
								//	},
								//	dcl.Action{
								//		Text:        "十六进制显示",
								//		OnTriggered: pg.notifyIconOpenActionHandler,
								//	},
								//	dcl.Action{
								//		Text:        "数据还原显示",
								//		OnTriggered: pg.restoreContentActionHandler,
								//	},
								//},
								Children: []dcl.Widget{
									dcl.TableView{
										AssignTo:         &pg.roleBaseDataTV,
										ColumnSpan:       1,
										CheckBoxes:       true,
										ColumnsOrderable: true,
										MultiSelection:   true,
										Columns: []dcl.TableViewColumn{
											{Title: "数据索引"},
											{Title: "数据名称"},
											{Title: "数据内容"},
											{Title: "数据说明"},
										},
										Model: pg.roleBaseDataModel,
										OnSelectedIndexesChanged: func() {
											fmt.Printf("OnSelectedIndexesChanged: %v\n", pg.roleBaseDataTV.SelectedIndexes())
										},
										OnItemActivated: func() {
											//idx := pg.roleBaseDataTV.CurrentIndex()
											//pg.roleBaseDataModel.SwitchRowCheckedState(idx)
										},
										OnMouseDown: func(x, y int, button walk.MouseButton) {
											// OnMouseDown函数会比OnSelectedIndexesChanged先执行，
											// 所以CurrentIndex()会不准确
											if button != walk.RightButton {
												return
											}

											//idx := pg.roleBaseDataTV.CurrentIndex()
											//fmt.Printf("OnMouseDown: %d\n", idx)
										},
									},
								},
							},
						},
					},

					dcl.Composite{ // 这里重新布局
						Font:   dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
						Layout: dcl.Grid{Columns: 5, Spacing: 1},
						Children: []dcl.Widget{
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "角色技能",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								OnClicked:  pg.onShowRoleSkillDialog,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "任务变量",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								OnClicked:  pg.onShowRoleTaskDialog,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "角色物品",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								//OnClicked:  pg.onDecodeROleBakData,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "角色状态",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								//OnClicked:  pg.onDecodeROleBakData,
							},
							dcl.PushButton{
								ColumnSpan: 1,
								Text:       "扩展数据",
								Font:       dcl.Font{Family: "微软雅黑", PointSize: 9, Bold: true},
								OnClicked:  pg.onShowRoleExtDataDialog,
							},
						},
					},
				},
			},
		},
	}
}
