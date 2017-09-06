package models

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beedb"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/devuser/beego-admin-demo/lib"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sort"
	"strings"
	"time"
	// _ "github.com/lib/pq"
	"math/rand"
	// _ "github.com/mattn/go-sqlite3"
)

const (
	cstLongdatefmt = "2006-01-02"
)

var o orm.Ormer

func Syncdb() {
	createdb()
	Connect()
	o = orm.NewOrm()
	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	insertUser()
	insertGroup()
	insertRole()
	insertNodes()
	//////// 文档清单 //////////////
	insertDocname()
	insertDocOrg()
	insertBaseurl()
	//////// 报表授权组系列信息 //////////////
	insertDocGroup()
	insertDocaccess()
	/////////////////////////////

	////////学习案例///////////////////////////
	insertCity()
	insertPerson()
	insertHotel()
	////////ETL系列信息//////////////
	insertBaseBizdate()
	insertDocDBConnection()
	//////////////////////////////////
	insertBaseCurrency()
	insertBaseExchangeRate()
	insertBaseDateUnit()
	insertBaseCodeDesc()
	//////////////////////////////////
	insertDanBao()
	insertDPSCklx()
	insertZhengZhaoZhongLei()

	////////////////////////////////////////////////////////////////////
	insertDianPiaoProductType()
	insertDianPiaoDetails()
	//如下为真实项目的构建数据，对于本Case来说没有意义，备注掉
	///////////////// SAP /////////////////
	//insertSAPOrgmap()
	//insertSAPAccountTitleCode()
	//insertSAPDayBalFlag()
	//insertSAPDayBal()

	//insertExportFileTracs()
	//
	///
	////////////////////////////////////////////////////////////////////
	insertDocSysvalue()
	insertDocPasswdRuleSuite()
	insertDocPasswdRule()
	//
	func() {
		db, _ := orm.GetDB("default")
		orm := beedb.New(db)
		orm.Exec(`insert into user_roles(user_id,role_id) values(1,1);`)
		orm.Exec(`insert into node_roles(node_id,role_id) select id,1 from node`)

	}()
	/////////////////////////////
	fmt.Println("database init is complete.\nPlease restart the application")

}

//数据库连接
func Connect() {
	var dns string
	db_type := beego.AppConfig.String("db_default_type")
	dbHost := beego.AppConfig.String("db_default_host")
	dbPort := beego.AppConfig.String("db_default_port")
	dbUser := beego.AppConfig.String("db_default_user")
	dbPass := beego.AppConfig.String("db_default_pass")
	dbName := beego.AppConfig.String("db_default_name")
	// db_path := beego.AppConfig.String("db_path")
	// db_sslmode := beego.AppConfig.String("db_sslmode")

	// db_path := beego.AppConfig.String("db_path")
	// db_sslmode := beego.AppConfig.String("db_sslmode")
	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DRMySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)
		break
	// case "postgres":
	// 	orm.RegisterDriver("postgres", orm.DRPostgres)
	// 	dns = fmt.Sprintf("dbname=%s host=%s  user=%s  password=%s  port=%s  sslmode=%s", dbName, dbHost, dbUser, dbPass, dbPort, db_sslmode)
	// case "sqlite3":
	// 	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	// 	if db_path == "" {
	// 		db_path = "./"
	// 	}
	// 	dns = fmt.Sprintf("%s%s.db", db_path, dbName)
	// 	break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	orm.RegisterDataBase("default", db_type, dns)
}

//创建数据库
func createdb() {

	db_type := beego.AppConfig.String("db_default_type")
	dbHost := beego.AppConfig.String("db_default_host")
	dbPort := beego.AppConfig.String("db_default_port")
	dbUser := beego.AppConfig.String("db_default_user")
	dbPass := beego.AppConfig.String("db_default_pass")
	dbName := beego.AppConfig.String("db_default_name")
	// db_path := beego.AppConfig.String("db_path")
	// db_sslmode := beego.AppConfig.String("db_sslmode")

	var dns string
	var sqlstring string
	beego.Critical("Checking :", db_type)
	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", dbUser, dbPass, dbHost, dbPort)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", dbName)
		break
	// case "postgres":
	// 	dns = fmt.Sprintf("host=%s  user=%s  password=%s  port=%s  sslmode=%s", dbHost, dbUser, dbPass, dbPort, db_sslmode)
	// 	sqlstring = fmt.Sprintf("CREATE DATABASE %s", dbName)
	// 	break
	// case "sqlite3":
	// 	if db_path == "" {
	// 		db_path = "./"
	// 	}
	// 	dns = fmt.Sprintf("%s%s.db", db_path, dbName)
	// 	os.Remove(dns)
	// 	sqlstring = "create table init (n varchar(32));drop table init;"
	// 	break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	db, err := sql.Open(db_type, dns)
	if err != nil {
		panic(err.Error())
	}
	//@TODO update by devuser
	//for oracle
	sqlstring = strings.ToUpper(sqlstring)
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", dbName, " created")
	}
	defer db.Close()

}

// mysql> desc user;
// +-----------------+--------------+------+-----+---------+----------------+
// | Field           | Type         | Null | Key | Default | Extra          |
// +-----------------+--------------+------+-----+---------+----------------+
// | id              | bigint(20)   | NO   | PRI | NULL    | auto_increment |
// | username        | varchar(32)  | NO   | UNI |         |                |
// | password        | varchar(32)  | NO   |     |         |                |
// | nickname        | varchar(32)  | NO   |     |         |                |
// | realname        | varchar(32)  | NO   |     |         |                |
// | title           | varchar(32)  | NO   |     |         |                |
// | delta           | varchar(32)  | NO   |     |         |                |
// | phone           | varchar(32)  | NO   |     |         |                |
// | mobel_phone     | varchar(32)  | NO   | UNI |         |                |
// | email           | varchar(255) | NO   | UNI |         |                |
// | remark          | varchar(200) | YES  |     | NULL    |                |
// | status          | int(11)      | NO   |     | 2       |                |
// | last_login_time | datetime     | YES  |     | NULL    |                |
// | login_time      | datetime     | YES  |     | NULL    |                |
// | last_login_ip   | datetime     | YES  |     | NULL    |                |
// | login_ip        | datetime     | YES  |     | NULL    |                |
// | online          | tinyint(1)   | YES  |     | NULL    |                |
// | creator         | bigint(20)   | NO   |     | 0       |                |
// | create_at       | datetime     | YES  |     | NULL    |                |
// | update_at       | datetime     | YES  |     | NULL    |                |
// | updater         | bigint(20)   | YES  |     | NULL    |                |
// | message         | varchar(255) | YES  |     | NULL    |                |
// | doc_org_id      | bigint(20)   | NO   |     | NULL    |                |
// +-----------------+--------------+------+-----+---------+----------------+
// 23 rows in set (0.00 sec)
//

func insertUser() {
	fmt.Println("insert user ...")
	boyosoftDocOrg := &DocOrg{Id: 1}
	huabeiDocOrg := &DocOrg{Id: 2}
	hebeiDocOrg := &DocOrg{Id: 7}
	users := []User{
		{Username: "admin", Nickname: "ClownFish", DocOrg: boyosoftDocOrg},
		{Username: "testuser1", Nickname: "testuser1", DocOrg: boyosoftDocOrg},
		{Username: "testuser2", Nickname: "testuser2", DocOrg: boyosoftDocOrg},
		{Username: "testuser3", Nickname: "testuser3", DocOrg: boyosoftDocOrg},
		{Username: "testuser4", Nickname: "testuser4", DocOrg: boyosoftDocOrg},

		{Username: "testuser5", Nickname: "testuser5", DocOrg: huabeiDocOrg},
		{Username: "testuser6", Nickname: "testuser6", DocOrg: huabeiDocOrg},
		{Username: "testuser7", Nickname: "testuser7", DocOrg: hebeiDocOrg},
		{Username: "testuser8", Nickname: "testuser8", DocOrg: hebeiDocOrg},
		{Username: "testuser9", Nickname: "testuser9", DocOrg: hebeiDocOrg},
		{Username: "testuser10", Nickname: "testuser10", DocOrg: hebeiDocOrg},
	}
	orm.Debug = true
	for index, v := range users {
		//fmt.Println("insert user with username as" + v.Username + "...")
		u := new(User)
		u.Id = int64(index + 1)
		u.Username = v.Username
		u.Nickname = v.Nickname
		u.Realname = v.Username
		u.Title = v.Title
		u.Delta = "5"
		u.Phone = fmt.Sprintf("6277000%d", index+1)
		u.MobelPhone = fmt.Sprintf("1391122000%d", index+1)
		u.Password = Pwdhash(v.Username)
		u.Email = fmt.Sprintf("%s@boyosoft.com.cn", v.Username)
		u.Remark = fmt.Sprintf("I'm %s", v.Username)
		u.Status = 2
		u.DocOrg = v.DocOrg

		u.LastLoginTime = getDay(-1)
		u.LoginTime = time.Now()
		u.LastLoginIP = "127.0.0.1"
		u.LoginIP = "127.0.0.1"

		o = orm.NewOrm()
		o.Insert(u)

	}
	orm.Debug = false
	fmt.Println("insert user end")
}

func insertDocname() {
	fmt.Println("insert docname ...")
	docnames := []Docname{
		{Docname: "DPS0001", Description: "自营存款来源月报", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},
		{Docname: "DPS0002", Description: "自营存款期限月报", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},
		{Docname: "DPS0005", Description: "流动性比例监测表", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},

		{Docname: "DPS0006", Description: "存款明细报表", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},
		{Docname: "DPS0007", Description: "存款排名表", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},
		{Docname: "DPS0008", Description: "存款剩余期限报表", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},
		{Docname: "DPS0009", Description: "流动性期限缺口统计表", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},
		{Docname: "DPS0010", Description: "存款客户组合查询", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},

		{Docname: "DPS0011", Description: "存款账户组合查询", Cabinet: "BOYOSOFT", Folder: "存款类", Paramdesc: "", FileType: "HTML"},

		{Docname: "FUD0001", Description: "资金日报汇总表", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},
		{Docname: "FUD0002", Description: "银行存款时点余额", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},
		{Docname: "FUD0003", Description: "银行存款时点余额（银行维度）", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},

		{Docname: "FUD0004", Description: "银行存款平均余额", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},

		{Docname: "FUD0005", Description: "银行存款平均余额（银行维度）", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},

		{Docname: "FUD0006", Description: "银行存款统计月报", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},
		{Docname: "FUD0007", Description: "结算类银行存款账户明细表", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},
		{Docname: "FUD0008", Description: "结算类银行存款账户平均余额表", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},
		{Docname: "FUD0009", Description: "时点备付率（银行维度）", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},
		{Docname: "FUD0010", Description: "平均备付率（银行维度）", Cabinet: "BOYOSOFT", Folder: "资金类", Paramdesc: "", FileType: "HTML"},

		{Docname: "CRD0001", Description: "信贷日报-汇总表（所有币种）", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		{Docname: "CRD0002", Description: "信贷日报—当月累计自营贷款发放明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		{Docname: "CRD0003", Description: "信贷日报—当月累计自营贷款回收明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0004	信贷类
		{Docname: "CRD0004", Description: "信贷日报—当月累计委托贷款发放明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0005	信贷类
		{Docname: "CRD0005", Description: "信贷日报—当月累计委托贷款回收明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0006	信贷类
		{Docname: "CRD0006", Description: "信贷业务收入统计表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0007	信贷类
		{Docname: "CRD0007", Description: "自营贷款余额明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},

		// CRD0009	信贷类	委托贷款余额明细表
		{Docname: "CRD0009", Description: "委托贷款余额明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0010	信贷类
		{Docname: "CRD0010", Description: "贷款统计表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0011	信贷类	贷款结构余额表
		{Docname: "CRD0011A", Description: "贷款结构余额表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		{Docname: "CRD0011B", Description: "贷款结构余额表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		{Docname: "CRD0011C", Description: "贷款结构余额表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0012	信贷类	贷款类业务汇总表
		{Docname: "CRD0012", Description: "贷款类业务汇总表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0013	信贷类	贷款期限日终表
		{Docname: "CRD0013", Description: "贷款期限日终表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0014	信贷类	担保汇总表
		{Docname: "CRD0014", Description: "担保汇总表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0015	信贷类	对外担保业务明细表
		{Docname: "CRD0015", Description: "对外担保业务明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0016	信贷类	贷款承诺业务汇总表
		{Docname: "CRD0016", Description: "贷款承诺业务汇总表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},
		// CRD0017	信贷类	贷款承诺明细表
		{Docname: "CRD0017", Description: "贷款承诺明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},

		{Docname: "CRD0018", Description: "贷款承诺明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},

		{Docname: "CRD0019", Description: "贷款承诺明细表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},

		{Docname: "CRD0020", Description: "各类贷款规模及收益统计表", Cabinet: "BOYOSOFT", Folder: "信贷类", Paramdesc: "", FileType: "HTML"},

		{Docname: "STM0001", Description: "营业结算月度报表", Cabinet: "BOYOSOFT", Folder: "结算类", Paramdesc: "", FileType: "HTML"},

		{Docname: "STM0002", Description: "客户账户月度报表", Cabinet: "BOYOSOFT", Folder: "结算类", Paramdesc: "", FileType: "HTML"},

		{Docname: "STM0003", Description: "营业类结算账户月度报表", Cabinet: "BOYOSOFT", Folder: "结算类", Paramdesc: "", FileType: "HTML"},
	}
	for _, v := range docnames {
		u := new(Docname)
		u.Docname = v.Docname
		u.Description = v.Description
		u.Cabinet = v.Cabinet
		u.Folder = v.Folder
		u.Paramdesc = v.Paramdesc
		u.FileType = v.FileType

		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert docname end")
}

// func insertUser() {
// 	fmt.Println("insert user ...")
// 	users := [5]User{
// 		{Username: "admin", Nickname: "ClownFish"},
// 		{Username: "testuser1", Nickname: "testuser1"},
// 		{Username: "testuser2", Nickname: "testuser2"},
// 		{Username: "testuser3", Nickname: "testuser3"},
// 		{Username: "testuser4", Nickname: "testuser4"},
// 	}
// 	orm.Debug = true
// 	for index, v := range users {
// 		fmt.Println("insert user with username as" + v.Username + "...")
// 		u := new(User)
// 		u.Id = int64(index + 1)
// 		u.Username = v.Username
// 		u.Nickname = v.Nickname
// 		u.Realname = v.Username
// 		u.Title = v.Title
// 		u.Delta = "5"
// 		u.Phone = fmt.Sprintf("6277000%d", index+1)
// 		u.MobelPhone = fmt.Sprintf("1391122000%d", index+1)
// 		u.Password = Pwdhash(v.Username)
// 		u.Email = fmt.Sprintf("%s@boyosoft.com.cn", v.Username)
// 		u.Remark = fmt.Sprintf("I'm %s", v.Username)
// 		u.Status = 2
// 		u.DocOrg = &DocOrg{Id: 1}
// 		o = orm.NewOrm()
// 		o.Insert(u)
//
// 	}
// 	orm.Debug = false
// 	fmt.Println("insert user end")
// }

func insertDocDBConnection() {
	fmt.Println("insert DocDBConnection ...")
	users := [5]DocDBConnection{
		{StartAt: time.Now(), Used: true, DataSourceName: "newkernel"},
		{StartAt: time.Now(), Used: true, DataSourceName: "DianPiao"},
		{StartAt: time.Now(), Used: true, DataSourceName: "xindai"},
		{StartAt: time.Now(), Used: true, DataSourceName: "zijin"},
		{StartAt: time.Now(), Used: true, DataSourceName: "qudao"},
	}
	orm.Debug = true
	for index, v := range users {
		fmt.Println("insert DocDBConnection")
		u := new(DocDBConnection)
		u.Id = int64(index + 1)
		u.StartAt = v.StartAt
		u.Used = v.Used
		u.DataSourceName = v.DataSourceName

		o = orm.NewOrm()
		o.Insert(u)

	}
	orm.Debug = false
	fmt.Println("insert DocDBConnection end")
}
func insertBaseurl() {
	fmt.Println("insert DocBaseurl ...")
	itemList := make([]*DocBaseurl, 0, 43)
	for index := 0; index < 43; index++ {
		fooItem := &DocBaseurl{DocnameId: int64(index + 1)}
		fooItem.Baseurl = "http://127.0.0.1:8080/reportdemo"
		itemList = append(itemList, fooItem)
	}

	orm.Debug = true
	for _, fooItem := range itemList {
		fmt.Println("insert DocBaseurl")
		u := new(DocBaseurl)
		u.Id = fooItem.Id
		u.DocnameId = fooItem.DocnameId
		u.Baseurl = fooItem.Baseurl

		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"
		o = orm.NewOrm()
		o.Insert(u)

	}
	orm.Debug = false
	fmt.Println("insert DocBaseurl end")
}

func insertDocGroup() {
	fmt.Println("insert DocGroup ...")

	users := []DocGroup{
		{Groupname: "Admin Group", Description: "admin", UserList: []*User{&User{Id: 1}}},
		{Groupname: "测试组", Description: "测试组", UserList: []*User{&User{Id: 2}}},
		{Groupname: "存款组", Description: "存款组", UserList: []*User{&User{Id: 5}}},
		{Groupname: "信贷组", Description: "信贷组", UserList: []*User{&User{Id: 4}}},
		{Groupname: "资金组", Description: "资金组", UserList: []*User{&User{Id: 5}}},
		{Groupname: "结算组", Description: "结算组", UserList: []*User{&User{Id: 6}}},
	}
	orm.Debug = true
	for index, v := range users {
		fmt.Println("insert DocGroup with Groupname as" + v.Groupname + "...")
		u := new(DocGroup)
		u.Id = int64(index + 1)
		u.Groupname = v.Groupname
		u.Description = v.Description
		//u.UserList
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)

	}
	func() {

		groupuserList := []struct {
			doc_groups_id int
			user_id       int
		}{
			{1, 1},
			{2, 1},
			{3, 2},
			{4, 2},
			{4, 5},
			{5, 4},
			{6, 5},
		}
		for _, itmX := range groupuserList {
			add := make(map[string]interface{})
			add["doc_groups_id"] = itmX.doc_groups_id
			add["user_id"] = itmX.user_id
			db, _ := orm.GetDB("default")
			orm := beedb.New(db)
			orm.SetTable("doc_groups_users").Insert(add)
		}
	}()

	orm.Debug = false
	fmt.Println("insert DocGroup end")
}

func insertDocOrg() {
	fmt.Println("insert doc_orgs ...")
	const (
		cstBenbu = "本部"
	)
	mapDocOrgs := make(map[string]string, 100)
	docorgs := []DocOrg{
		{Orgcode: "000000", StdOrgcode: "000000", Orgname: "BOYOSOFT"},
		{Orgcode: "000001", StdOrgcode: "000001", Orgname: "直属营业部"},

		{Orgcode: "001100", StdOrgcode: "001100", Orgname: "东北分公司"},
		{Orgcode: "001101", StdOrgcode: "001101", Orgname: cstBenbu, FullOrgname: "东北分公司本部"},
		{Orgcode: "001102", StdOrgcode: "001102", Orgname: "吉林"},
		{Orgcode: "001103", StdOrgcode: "001103", Orgname: "黑龙江"},
		{Orgcode: "001104", StdOrgcode: "001104", Orgname: "内蒙"},

		{Orgcode: "001200", StdOrgcode: "001200", Orgname: "西北分公司"},
		{Orgcode: "001201", StdOrgcode: "001201", Orgname: cstBenbu, FullOrgname: "西北分公司本部"},
		{Orgcode: "001202", StdOrgcode: "001202", Orgname: "甘肃"},
		{Orgcode: "001203", StdOrgcode: "001203", Orgname: "宁夏"},
		{Orgcode: "001204", StdOrgcode: "001204", Orgname: "青海"},
		{Orgcode: "001205", StdOrgcode: "001205", Orgname: "新疆"},

		{Orgcode: "001300", StdOrgcode: "001300", Orgname: "华中分公司"},
		{Orgcode: "001301", StdOrgcode: "001301", Orgname: cstBenbu, FullOrgname: "华中分公司本部"},
		{Orgcode: "001302", StdOrgcode: "001302", Orgname: "湖南"},
		{Orgcode: "001303", StdOrgcode: "001303", Orgname: "河南"},
		{Orgcode: "001304", StdOrgcode: "001304", Orgname: "江西"},
		{Orgcode: "001305", StdOrgcode: "001305", Orgname: "四川"},
		{Orgcode: "001306", StdOrgcode: "001306", Orgname: "重庆"},

		{Orgcode: "001400", StdOrgcode: "001400", Orgname: "华东分公司"},
		{Orgcode: "001401", StdOrgcode: "001401", Orgname: cstBenbu, FullOrgname: "华东公司本部"},
		{Orgcode: "001402", StdOrgcode: "001402", Orgname: "江苏"},
		{Orgcode: "001403", StdOrgcode: "001403", Orgname: "浙江"},
		{Orgcode: "001404", StdOrgcode: "001404", Orgname: "安徽"},
		{Orgcode: "001405", StdOrgcode: "001405", Orgname: "福建"},

		{Orgcode: "001500", StdOrgcode: "001500", Orgname: "华北分公司"},
		{Orgcode: "001501", StdOrgcode: "001501", Orgname: cstBenbu, FullOrgname: "华北公司本部"},
		{Orgcode: "001502", StdOrgcode: "001502", Orgname: "天津"},
		{Orgcode: "001503", StdOrgcode: "001503", Orgname: "河北"},
		{Orgcode: "001504", StdOrgcode: "001504", Orgname: "山西"},
		{Orgcode: "001505", StdOrgcode: "001505", Orgname: "山东"},
	}

	for _, v := range docorgs {
		mapDocOrgs[v.Orgcode] = v.Orgname
	}

	for _, v := range docorgs {
		u := new(DocOrg)
		u.Orgcode = v.Orgcode
		u.StdOrgcode = v.StdOrgcode
		u.Orgname = v.Orgname
		if u.Orgname == cstBenbu {
			baseOrgcode := fmt.Sprintf("%s00", u.Orgcode[:4])
			if orgName, found := mapDocOrgs[baseOrgcode]; found {
				suggestFullOrgname := fmt.Sprintf("%s%s", orgName, cstBenbu)
				u.FullOrgname = suggestFullOrgname
			}
		} else {
			u.FullOrgname = u.Orgname
		}
		// if u.FullOrgname == "" {
		// 	u.FullOrgname = v.Orgname
		// }

		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		fmt.Printf("%q\n", u)
		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert doc_orgs end")
}

func checkErr(err error) {

	if err != nil {
		log.Printf("checkErr %q\n", err)
		panic(err)
	} //end of if
} //end of checkErr

func insertDocaccess() {
	fmt.Println("insert Docaccess ...")
	docaccessList := []*Docaccess{
		&Docaccess{Docname: &Docname{Id: 1}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 2}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 3}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 4}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 5}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 6}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 7}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 8}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 9}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 10}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 11}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},
		&Docaccess{Docname: &Docname{Id: 12}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"},

		&Docaccess{Docname: &Docname{Id: 1}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},
		&Docaccess{Docname: &Docname{Id: 2}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},
		&Docaccess{Docname: &Docname{Id: 3}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},
		&Docaccess{Docname: &Docname{Id: 4}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},

		&Docaccess{Docname: &Docname{Id: 5}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},
		&Docaccess{Docname: &Docname{Id: 6}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},
		&Docaccess{Docname: &Docname{Id: 7}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},
		&Docaccess{Docname: &Docname{Id: 8}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "G"},
	}
	//append of testuser10

	testuser10UserId := func() (userId int64) {
		userId = -1
		if one := GetUserByUsername("testuser10"); one.Id != 0 {
			userId = one.Id
		}
		return userId
	}()

	fmt.Printf("testuser10 userId: %s\n", testuser10UserId)
	extDocaccessList := func() (extDocaccessList []*Docaccess) {

		db, _ := orm.GetDB("default")
		orm := beedb.New(db)
		var docnameList []struct {
			Id int64 `orm:"column(id);auto;pk"`
		}
		err := orm.SetTable(beego.AppConfig.String("doc_docname_table")).Select("Id").OrderBy("Id").FindAll(&docnameList)
		checkErr(err)
		extDocaccessList = make([]*Docaccess, 0, len(docnameList))
		fmt.Printf("docnameList Len:%d\n", len(docnameList))
		for _, docnameX := range docnameList {
			itmX := &Docaccess{Docname: &Docname{Id: docnameX.Id}, StartAt: getDay(-100), OverAt: getDay(+100), DocAccessType: "U"}
			docaccessList = append(docaccessList, itmX)
			extDocaccessList = append(extDocaccessList, itmX)
		}
		return extDocaccessList
	}()

	for _, v := range docaccessList {
		u := new(Docaccess)
		u.Docname = v.Docname
		u.StartAt = v.StartAt
		u.OverAt = v.OverAt
		u.DocAccessType = v.DocAccessType

		o = orm.NewOrm()
		insertId, _ := o.Insert(u)
		v.Id = insertId
	}

	func() {

		groupuserList := []struct {
			doc_docaccess_id int64
			doc_groups_id    int64
		}{
			{13, 3},
			{14, 3},
			{15, 3},
			{16, 3},
			{17, 3},
			{18, 3},
			{19, 3},
			{20, 3},

			{13, 4},
			{14, 4},
			{15, 4},
			{16, 4},
			{17, 4},
			{18, 4},
			{19, 4},
			{20, 4},

			{13, 5},
			{14, 5},
			{15, 5},
			{16, 5},
			{17, 5},
			{18, 5},
			{19, 5},
			{20, 5},

			{13, 6},
			{14, 6},
			{15, 6},
			{16, 6},
			{17, 6},
			{18, 6},
			{19, 6},
			{20, 6},
		}
		for _, itmX := range groupuserList {
			add := make(map[string]interface{})
			add["doc_docaccess_id"] = itmX.doc_docaccess_id
			add["doc_groups_id"] = itmX.doc_groups_id
			db, _ := orm.GetDB("default")
			orm := beedb.New(db)
			orm.SetTable("doc_docaccess_doc_groupss").Insert(add)
		}
	}()
	func() {

		groupuserList := []struct {
			doc_docaccess_id int64
			user_id          int64
		}{
			{1, 2},
			{2, 2},
			{3, 2},
			{4, 2},
			{4, 2},
			{5, 2},
			{6, 2},

			{7, 2},
			{8, 2},
			{9, 2},
			{10, 2},
			{11, 2},
			{12, 2},

			{1, 5},
			{2, 5},
			{3, 5},
			{4, 5},
			{4, 5},
			{5, 5},
			{6, 5},

			{1, 4},
			{2, 4},
			{3, 4},
			{4, 4},
			{4, 4},
			{5, 4},
			{6, 4},

			{1, 5},
			{2, 5},
			{3, 5},
			{4, 5},

			{1, 6},
			{2, 6},

			{1, 7},
			{2, 7},

			{1, 8},
		}
		for _, itmX := range groupuserList {
			add := make(map[string]interface{})
			add["doc_docaccess_id"] = itmX.doc_docaccess_id
			add["user_id"] = itmX.user_id
			db, _ := orm.GetDB("default")
			orm := beedb.New(db)
			orm.SetTable("doc_docaccess_users").Insert(add)
		}
	}()
	func() {
		type FOO struct {
			doc_docaccess_id int64
			user_id          int64
		}
		groupuserList := []*FOO{}
		for _, itmX := range extDocaccessList {
			groupuserList = append(groupuserList, &FOO{itmX.Id, testuser10UserId})
		}

		for _, itmX := range groupuserList {
			add := make(map[string]interface{})
			add["doc_docaccess_id"] = itmX.doc_docaccess_id
			add["user_id"] = itmX.user_id
			db, _ := orm.GetDB("default")
			orm := beedb.New(db)
			orm.SetTable("doc_docaccess_users").Insert(add)
		}
	}()
	fmt.Println("insert Docaccess end")
}

func getDay(dayCount int) time.Time {
	// d, _ := time.ParseDuration("-24h")
	// dt := time.Now().Add(d * dayCount)
	// timeStr := time.Now().Format(cstLongdatefmt)

	// t, _ := time.Parse(cstLongdatefmt, timeStr)
	dtX := time.Now().AddDate(0, 0, dayCount)
	timeStr := dtX.Format(cstLongdatefmt)
	t, _ := time.Parse(cstLongdatefmt, timeStr)
	return t
}

//insertDianPiaoProductType 构造电票产品类型数据
//用以区分如下类型
// 1个月（含1月）
// 1-3个月贴现（含3月）
// 3-12月（含12月）贴现
func insertDianPiaoProductType() {
	fmt.Println("insert DianPiaoProductType ...")
	const (
		cstStateDONE = "DONE"
	)
	docorgs := []DianPiaoProductType{
		{
			Id:              1,
			ProductTypeCode: "01",
			Description:     "1个月（含1月）"},
		{
			Id:              2,
			ProductTypeCode: "03",
			Description:     "1-3个月贴现（含3月）"},
		{
			Id:              1,
			ProductTypeCode: "12",
			Description:     "3-12月（含12月）贴现"},
	}
	for _, v := range docorgs {
		u := new(DianPiaoProductType)
		u.Id = v.Id
		u.ProductTypeCode = v.ProductTypeCode
		u.Description = v.Description

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert DianPiaoProductType end")
}

//insertDianPiaoDetails 插入电票交易明细
func insertDianPiaoDetails() {
	fmt.Println("insert DianPiaoDetails ...")
	const (
		cstStateDONE = "DONE"
	)
	docorgs := []DianPiaoDetails{
		{
			Id:             1,
			DianPiaoCode:   "0000001",
			Orgcode:        cstFirstOrgcode,
			Bizdate:        getDay(-1).Format(cstStdDataFmt),
			DebitOrCredit:  cstDEBT,
			ProductType:    &DianPiaoProductType{Id: 1, ProductTypeCode: "01"},
			Balance:        100.0,
			Interest:       100.0,
			InsidePlanFlag: false,
			StartAt:        getDay(-100).Format(cstStdDataFmt),
			OverAt:         getDay(-50).Format(cstStdDataFmt),
		},
	}
	for _, v := range docorgs {
		u := new(DianPiaoDetails)
		u.Id = v.Id
		u.DianPiaoCode = v.DianPiaoCode
		u.Orgcode = v.Orgcode

		u.Bizdate = v.Bizdate
		u.DebitOrCredit = v.DebitOrCredit
		u.Balance = v.Balance
		u.Interest = v.Interest
		u.InsidePlanFlag = v.InsidePlanFlag

		u.StartAt = v.StartAt
		u.OverAt = v.OverAt

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert DianPiaoDetails end")
}

func insertExportFileTracs() {
	fmt.Println("insert ExportFileTrac ...")
	docorgs := make([]*ExportFileTrac, 0, 1000)
	o := orm.NewOrm()
	node := new(Docname)
	qs := o.QueryTable(node)
	var nodes []orm.Params
	// var nodes []Docname
	qs.OrderBy("docname").Values(&nodes, "Id")
	exportFileTypes := []string{cstExportFileTypePDF, cstExportFileTypeWPS, cstExportFileTypeXLS}

	var index int64
	for _, m := range nodes {
		docnameId := orm.ToInt64(m["Id"])
		exportFileType := exportFileTypes[rand.Intn(len(exportFileTypes))]
		index += 1
		t := getDay(rand.Intn(10))
		strDay := t.Format("20060102")
		xlsTrac := &ExportFileTrac{
			Id:             index,
			ExportFileName: fmt.Sprintf("%s-%08d.%s", strDay, index, exportFileType),
			State:          cstStateDONE,
			ExportFileType: exportFileType,
			TransferAt:     time.Now(),
			Docname:        &Docname{Id: docnameId},
		}
		docorgs = append(docorgs, xlsTrac)
	}
	for _, u := range docorgs {

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		log.Printf("%q\n", u)
		o = orm.NewOrm()
		_, err := o.Insert(u)
		checkErr(err)
	}
	fmt.Println("insert ExportFileTrac end")
}

func insertBaseBizdate() {
	fmt.Println("insert BaseBizdate ...")
	const (
		cstStateDONE = "DONE"
	)
	docorgs := []BaseBizdate{
		{Id: 1, Bizdate: getDay(0).Format("20060102"), State: cstStateDONE},
		{Id: 2, Bizdate: getDay(-1).Format("20060102"), State: cstStateDONE},
		{Id: 3, Bizdate: getDay(-2).Format("20060102"), State: cstStateDONE},
		{Id: 4, Bizdate: getDay(-3).Format("20060102"), State: cstStateDONE},
		{Id: 5, Bizdate: getDay(-4).Format("20060102"), State: cstStateDONE},
		{Id: 6, Bizdate: getDay(-5).Format("20060102"), State: cstStateDONE},
	}
	for _, v := range docorgs {
		u := new(BaseBizdate)
		u.Id = v.Id
		u.Bizdate = v.Bizdate
		u.State = v.State
		u.CreateAt = time.Now()
		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert BaseBizdate end")
}

func insertZhengZhaoZhongLei() {
	fmt.Println("insert base_common_zhengzhaozhonglei_table ...")

	docorgs := []ZhengZhaoZhongLei{
		{Id: 1, ZhengZhaoZhongLeiCode: "0001", Description: "营业执照"},
		{Id: 2, ZhengZhaoZhongLeiCode: "0002", Description: "组织机构代码证"},
	}
	for _, v := range docorgs {
		u := new(ZhengZhaoZhongLei)
		u.Id = v.Id
		u.ZhengZhaoZhongLeiCode = v.ZhengZhaoZhongLeiCode
		u.Description = v.Description
		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert base_common_zhengzhaozhonglei_table end")
}

func insertBaseCodeDesc() {
	fmt.Println("insert base_common_codedesc_table ...")

	docorgs := []BaseCodeDesc{
		{Id: 1, Kvgroup: "KEHU_STATUS", Code: "0", Description: "正常"},
		{Id: 1, Kvgroup: "KEHU_STATUS", Code: "1", Description: "归并"},
		{Id: 1, Kvgroup: "KEHU_STATUS", Code: "2", Description: "封存"},

		{Id: 1, Kvgroup: "ZHANGHU_STATUS", Code: "ZHS01", Description: "正常"},
		{Id: 1, Kvgroup: "ZHANGHU_STATUS", Code: "ZHS02", Description: "冻结"},
		{Id: 1, Kvgroup: "ZHANGHU_STATUS", Code: "ZHS03", Description: "久悬"},
		{Id: 1, Kvgroup: "ZHANGHU_STATUS", Code: "ZHS04", Description: "销户"},

		{Id: 3, Kvgroup: "HOSTKvgroup", Code: "G102", Description: "南网集团"},
		{Id: 1, Kvgroup: "HOSTKvgroup", Code: "G100", Description: "电网"},
		{Id: 2, Kvgroup: "HOSTKvgroup", Code: "G101", Description: "国网集团"},
		{Id: 3, Kvgroup: "HOSTKvgroup", Code: "G102", Description: "南网集团"},
		{Id: 4, Kvgroup: "HOSTKvgroup", Code: "G103", Description: "内蒙古电网集团"},
		{Id: 5, Kvgroup: "HOSTKvgroup", Code: "G104", Description: "其他"},
		{Id: 6, Kvgroup: "HOSTKvgroup", Code: "G200", Description: "发电"},
		{Id: 7, Kvgroup: "HOSTKvgroup", Code: "G201", Description: "大唐集团"},
		{Id: 8, Kvgroup: "HOSTKvgroup", Code: "G202", Description: "华电集团"},
		{Id: 9, Kvgroup: "HOSTKvgroup", Code: "G203", Description: "华能集团"},
		{Id: 10, Kvgroup: "HOSTKvgroup", Code: "G204", Description: "中电投集团"},
		{Id: 11, Kvgroup: "HOSTKvgroup", Code: "G205", Description: "国电"},
		{Id: 12, Kvgroup: "HOSTKvgroup", Code: "G206", Description: "其他"},
		{Id: 13, Kvgroup: "HOSTKvgroup", Code: "G300", Description: "其他"},

		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10000", Description: "电网公司"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10100", Description: "国网公司总部"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10200", Description: "国网公司分部"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10300", Description: "省（自治区、直辖市）电力公司"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10400", Description: "地市级电力公司"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10500", Description: "县级电力公司"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10001", Description: "全资"},

		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10002", Description: "控股"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10003", Description: "参股"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10004", Description: "代管"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "10005", Description: "其他"},

		{Id: 13, Kvgroup: "KEHULEIXING", Code: "20000", Description: "发电公司"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "30000", Description: "产业公司"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "40000", Description: "科研教培单位"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "50000", Description: "专业公司"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "60000", Description: "其他"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "60100", Description: "集体企业（国网集团-非并表）"},
		{Id: 13, Kvgroup: "KEHULEIXING", Code: "60200", Description: "其它"},
	}

	currentGroup := ""
	var currentSortFactor uint64 = 0

	for i, v := range docorgs {
		u := new(BaseCodeDesc)
		u.Id = int64(i + 1)
		u.Source = "SELF"
		u.Kvgroup = v.Kvgroup
		u.Code = v.Code
		u.Description = v.Description

		if v.Kvgroup == currentGroup {
			currentSortFactor += 1
		} else {
			currentGroup = v.Kvgroup
			currentSortFactor = 1
		}
		u.SortFactor = currentSortFactor

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert base_common_codedesc_table end")

}

func insertDocSysvalue() {
	fmt.Println("insert DocSysvalue ...")

	docorgs := []DocSysvalue{
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_DAYS",
			Value:      "PASSWORD_DAYS",
			SortFactor: 1,
			Label:      "PASSWORD_DAYS"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_CHANGERIGHTNOW",
			Value:      "PASSWORD_CHANGERIGHTNOW",
			SortFactor: 1,
			Label:      "PASSWORD_CHANGERIGHTNOW"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_INITPATTEN",
			Value:      "PASSWORD_INITPATTEN",
			SortFactor: 1,
			Label:      "PASSWORD_INITPATTEN"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_INITFIX",
			Value:      "PASSWORD_INITFIX",
			SortFactor: 1,
			Label:      "PASSWORD_INITFIX"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_INITRANDOM",
			Value:      "PASSWORD_INITRANDOM",
			SortFactor: 1,
			Label:      "PASSWORD_INITRANDOM"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_INITRANDOMLEN",
			Value:      "PASSWORD_INITRANDOMLEN",
			SortFactor: 1,
			Label:      "PASSWORD_INITRANDOMLEN"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_LEN_MIN",
			Value:      "PASSWORD_LEN_MIN",
			SortFactor: 1,
			Label:      "PASSWORD_LEN_MIN"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_LEN_MIN",
			Value:      "PASSWORD_LEN_MIN",
			SortFactor: 1,
			Label:      "PASSWORD_LEN_MIN"},

		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_LEN_MAX",
			Value:      "PASSWORD_LEN_MAX",
			SortFactor: 1,
			Label:      "PASSWORD_LEN_MAX"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_LEN_MAX",
			Value:      "PASSWORD_LEN_MAX",
			SortFactor: 1,
			Label:      "PASSWORD_LEN_MAX"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "LOGINTRY_COUNT",
			Value:      "LOGINTRY_COUNT",
			SortFactor: 1,
			Label:      "LOGINTRY_COUNT"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "LOGINTRY_RESETBYDAY",
			Value:      "LOGINTRY_RESETBYDAY",
			SortFactor: 1,
			Label:      "LOGINTRY_RESETBYDAY"},

		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_CHARS",
			Value:      "PASSWORD_CHARS",
			SortFactor: 1,
			Label:      "PASSWORD_CHARS"},
		{
			Id:         1,
			Kvgroup:    "PASSWD_RULE",
			Key:        "PASSWORD_DUPCHARS",
			Value:      "PASSWORD_DUPCHARS",
			SortFactor: 1,
			Label:      "PASSWORD_DUPCHARS"}}
	for index, v := range docorgs {
		u := new(DocSysvalue)
		u.Id = int64(index + 1)
		u.Kvgroup = v.Kvgroup
		u.Key = v.Key
		u.Value = v.Value
		u.SortFactor = uint64(index + 1)
		u.Label = v.Label

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert DocSysvalue end")

}

///////////////////////////////////
// 裴金胤查询的四种类型的存款
// 定期、协议、通知、保证金存款
// 上述四种基本都分同业和对公
///////////////////////////////////

func insertDPSCklx() {
	fmt.Println("insert base_dps_cklx_table ...")

	docorgs := []DPSCklx{
		{Id: 1, CKLXCode: "0001", Description: "活期"},
		{Id: 2, CKLXCode: "0002", Description: "定期"},
		{Id: 3, CKLXCode: "0003", Description: "通知存款"},
		{Id: 4, CKLXCode: "0004", Description: "协定存款"},
		{Id: 5, CKLXCode: "0005", Description: "协议存款"},
		{Id: 6, CKLXCode: "0006", Description: "保证金存款"},
		{Id: 7, CKLXCode: "0007", Description: "同业存款"}}
	for _, v := range docorgs {
		u := new(DPSCklx)
		u.Id = v.Id
		u.CKLXCode = v.CKLXCode
		u.Description = v.Description
		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert base_dps_cklx_table end")

}

// Id        int64  `orm:"column(id);pk;auto"`
// SuiteName string `orm:"column(rule_name);size(255)"`
// //开始时间、结束时间
// StartAt time.Time `orm:"column(start_at);type(datetime);null"`
// OverAt  time.Time `orm:"column(over_at);type(datetime);null"`
//
// DocPasswdRuleList []*DocPasswdRule `orm:"reverse(many)"`
// //管控部分
// Status   uint64    `orm:"column(status);default(2)" form:"Status"`
// Creator  *User     `orm:"column(creator);null;rel(fk)"`
// CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
// Updater  *User     `orm:"column(updater);null;rel(fk)"`
// UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
// Message  string    `orm:"column(message);size(255);null"`
const (
	cstStdDataFmt = "20060102"
)

func insertDocPasswdRuleSuite() {
	fmt.Println("insert DocPasswdRuleSuite ...")

	dt20160701, _ := time.Parse(cstStdDataFmt, "20160701")
	dt20161231, _ := time.Parse(cstStdDataFmt, "20161231")

	dt20170101, _ := time.Parse(cstStdDataFmt, "20170101")
	dt20170630, _ := time.Parse(cstStdDataFmt, "20170630")

	dt20170701, _ := time.Parse(cstStdDataFmt, "20170701")
	dt20171231, _ := time.Parse(cstStdDataFmt, "20171231")

	ruleSuites := []DocPasswdRuleSuite{
		{
			Id:        1,
			SuiteName: "20160701测试阶段",
			StartAt:   dt20160701,
			OverAt:    dt20161231,
			Current:   true,
		},
		{
			Id:        2,
			SuiteName: "20170101系统投产初始化",
			StartAt:   dt20170101,
			OverAt:    dt20170630,
			Current:   false,
		},
		{
			Id:        3,
			SuiteName: "20170701投产半年后",
			StartAt:   dt20170701,
			OverAt:    dt20171231,
			Current:   false,
		},
	}
	for _, v := range ruleSuites {
		u := new(DocPasswdRuleSuite)
		u.Id = v.Id
		u.SuiteName = v.SuiteName
		u.StartAt = v.StartAt
		u.OverAt = v.OverAt
		u.Current = v.Current

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert DocPasswdRuleSuite end")

}

// Id          int64  `orm:"column(id);pk;auto"`
// RuleName    string `orm:"column(rule_name);size(255)"`
// RuleDesc    string `orm:"column(rule_desc);size(255)"`
// RuleContent string `orm:"column(rule_content);size(255)"`
//
// State string `orm:"default(TODO)" form:"State"  valid:"Required"`
//
// DocPasswdRuleSuite *DocPasswdRuleSuite `orm:"rel(fk)"`

func insertDocPasswdRule() {
	fmt.Println("insert DocPasswdRule ...")

	rules := []DocPasswdRule{
		{
			Id:          1,
			RuleName:    "PASSWORD_DAYS",
			RuleDesc:    "密码有效期",
			RuleType:    "RT_INT",
			RuleContent: "90"},
		{
			Id:          1,
			RuleName:    "PASSWORD_CHANGERIGHTNOW",
			RuleDesc:    "首次登录后立即修改密码",
			RuleType:    "RT_INT",
			RuleContent: "1"},
		{
			Id:          1,
			RuleName:    "PASSWORD_INITPATTEN",
			RuleDesc:    "初始化密码方式",
			RuleType:    "RT_STRING",
			RuleContent: "PASSWORD_INITFIX"},
		{
			Id:          1,
			RuleName:    "PASSWORD_INITFIX",
			RuleDesc:    "固定密码",
			RuleType:    "RT_STRING",
			RuleContent: "123456"},

		{
			Id:          1,
			RuleName:    "PASSWORD_INITRANDOM",
			RuleDesc:    "随机密码候选字符",
			RuleType:    "RT_STRING",
			RuleContent: "1234567890"},

		{
			Id:          1,
			RuleName:    "PASSWORD_INITRANDOMLEN",
			RuleDesc:    "初始化密码长度",
			RuleType:    "RT_INT",
			RuleContent: "8"},

		{
			Id:          1,
			RuleName:    "PASSWORD_LEN_MIN",
			RuleDesc:    "密码长度限制不得短于",
			RuleType:    "RT_INT",
			RuleContent: "4"},
		{
			Id:          2,
			RuleName:    "PASSWORD_LEN_MAX",
			RuleDesc:    "密码长度限制不得超过",
			RuleType:    "RT_INT",
			RuleContent: "16"},
		{
			Id:          3,
			RuleName:    "LOGINTRY_COUNT",
			RuleDesc:    "密码尝试次数",
			RuleType:    "RT_INT",
			RuleContent: "3"},
		{
			Id:          4,
			RuleName:    "LOGINTRY_RESETBYDAY",
			RuleDesc:    "密码尝试次数重制限制",
			RuleType:    "RT_INT",
			RuleContent: "1"},
		{
			Id:          5,
			RuleName:    "PASSWORD_CHARS",
			RuleDesc:    "密码字符组成",
			RuleType:    "RT_STRING",
			RuleContent: "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ!@#&"},
		{
			Id:          6,
			RuleName:    "PASSWORD_DUPCHARS",
			RuleDesc:    "密码连续字符出现次数",
			RuleType:    "RT_INT",
			RuleContent: "3"},
	}
	suiteIds := []int64{1, 2, 3}
	for _, suiteId := range suiteIds {

		for _, v := range rules {

			u := new(DocPasswdRule)
			u.Id = v.Id
			u.RuleName = v.RuleName
			u.RuleDesc = v.RuleDesc

			u.RuleContent = v.RuleContent
			u.RuleType = v.RuleType
			u.DocPasswdRuleSuite = &DocPasswdRuleSuite{Id: suiteId}

			u.Status = 2
			u.Creator = &User{Id: 1}
			u.Updater = &User{Id: 1}
			u.CreateAt = time.Now()
			u.UpdateAt = time.Now()
			u.Message = "init the data"

			o = orm.NewOrm()
			o.Insert(u)
		}
	}
	fmt.Println("insert DocPasswdRule end")

}

func insertDanBao() {
	fmt.Println("insert base_crd_danbao_table ...")
	const (
		cstStateDONE = "DONE"
	)
	docorgs := []CRDDanBao{
		{Id: 1, DanBaoXingShi: "0001", Description: "保证金"},
		{Id: 2, DanBaoXingShi: "0002", Description: "抵押"},
		{Id: 3, DanBaoXingShi: "0003", Description: "质押"},
		{Id: 4, DanBaoXingShi: "0004", Description: "保证人"},
		{Id: 5, DanBaoXingShi: "0005", Description: "信用"}}
	for _, v := range docorgs {
		u := new(CRDDanBao)
		u.Id = v.Id
		u.DanBaoXingShi = v.DanBaoXingShi
		u.Description = v.Description
		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert base_crd_danbao_table end")
}

//insertBaseDateUnit 初始化期限单位
//
// <restrictionType id="E_QIXIANDW" longname="期限单位" base="BaseType.U_LEIXIN01" tags="AP">
// 	<enumeration id="None" value="N" longname="无关"/>
// 	<enumeration id="Day" value="D" longname="日"/>
// 	<enumeration id="Week" value="W" longname="周"/>
// 	<enumeration id="Xun" value="S" longname="旬"/>
// 	<enumeration id="Month" value="M" longname="月"/>
// 	<enumeration id="Quart" value="Q" longname="季"/>
// 	<enumeration id="HalfYear" value="H" longname="半年"/>
// 	<enumeration id="Year" value="Y" longname="年"/>
// 	<subenums>
func insertBaseDateUnit() {
	fmt.Println("insert base_common_dateunit_table ...")

	docorgs := []BaseDateUnit{
		{Id: 1, DateUnit: "天"},
		{Id: 2, DateUnit: "周"},
		{Id: 3, DateUnit: "月"},
		{Id: 4, DateUnit: "年"},
	}
	for _, v := range docorgs {
		u := new(BaseDateUnit)
		u.Id = v.Id
		u.DateUnit = v.DateUnit

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert base_common_dateunit_table end")
}

// StartAt time.Time `orm:"unique;column(start_at);type(datetime);null"`
// OverAt  time.Time `orm:"unique;column(over_at);type(datetime);null"`
// //源货币代号
// SrcHuobdaih string `orm:"column(srchuobdaih);unique;size(2);valid(required)"`
// //目标货币代号
// DstHuobdaih string `orm:"column(dsthuobdaih);unique;size(2);valid(required)"`
// //汇率
// ExchangeRate float64 `orm:"column(exchangerate);type(float64-decimal);decimals(6);digits(23);null"`
//
// //============= 与标签有有关结束 =============
// Status int `orm:"column(status);default(2)" form:"Status"`
// //============= 管控部分开始 =============
// Creator  int64    `orm:"column(creator);size(20);null"`
// UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
// Updater  *User     `orm:"column(updater);null;rel(fk)"`
// Message  string    `orm:"column(message);size(255);null"`
// CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`

func insertBaseExchangeRate() {
	fmt.Println("insert base_common_exchangerate ...")

	docorgs := []BaseExchangeRate{
		{Id: 1, StartAt: getDay(-400).Format(cstStdDataFmt), OverAt: getDay(-301).Format(cstStdDataFmt), SrcHuobdaih: "01", DstHuobdaih: "02", ExchangeRate: 6.03},
		{Id: 2, StartAt: getDay(-300).Format(cstStdDataFmt), OverAt: getDay(-201).Format(cstStdDataFmt), SrcHuobdaih: "01", DstHuobdaih: "03", ExchangeRate: 7.03},
		{Id: 3, StartAt: getDay(-200).Format(cstStdDataFmt), OverAt: getDay(-101).Format(cstStdDataFmt), SrcHuobdaih: "01", DstHuobdaih: "04", ExchangeRate: 8.03},
		{Id: 4, StartAt: getDay(-100).Format(cstStdDataFmt), OverAt: getDay(-1).Format(cstStdDataFmt), SrcHuobdaih: "01", DstHuobdaih: "05", ExchangeRate: 9.03},
		{Id: 5, StartAt: getDay(0).Format(cstStdDataFmt), OverAt: getDay(100).Format(cstStdDataFmt), SrcHuobdaih: "01", DstHuobdaih: "06", ExchangeRate: 10.03},

		{Id: 6, StartAt: getDay(-300).Format(cstStdDataFmt), OverAt: getDay(-201).Format(cstStdDataFmt), SrcHuobdaih: "02", DstHuobdaih: "03", ExchangeRate: 12.03},
		{Id: 7, StartAt: getDay(-200).Format(cstStdDataFmt), OverAt: getDay(-101).Format(cstStdDataFmt), SrcHuobdaih: "02", DstHuobdaih: "04", ExchangeRate: 13.03},
		{Id: 8, StartAt: getDay(-100).Format(cstStdDataFmt), OverAt: getDay(-1).Format(cstStdDataFmt), SrcHuobdaih: "02", DstHuobdaih: "05", ExchangeRate: 14.03},
		{Id: 9, StartAt: getDay(0).Format(cstStdDataFmt), OverAt: getDay(100).Format(cstStdDataFmt), SrcHuobdaih: "02", DstHuobdaih: "06", ExchangeRate: 15.03},
	}
	for _, v := range docorgs {
		u := new(BaseExchangeRate)
		u.Id = v.Id
		u.StartAt = v.StartAt
		u.OverAt = v.OverAt
		u.SrcHuobdaih = v.SrcHuobdaih
		u.DstHuobdaih = v.DstHuobdaih
		u.ExchangeRate = v.ExchangeRate

		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert base_common_exchangerate end")
}

//insertBaseCurrency
// 01--人民币
// 12--英镑
// 13--港币
// 14--美元
// 15--瑞士法郎
// 21--瑞典克朗
// 22--丹麦克朗
// 23--挪威克朗
// 27--日元
// 28--加拿大元
// 29--澳大利亚元
// 18--新加坡元
// 38--欧元
// 43--韩元
// 81--澳门元
// 82--新台币
// 83--津巴布韦币
// 87--新西兰元
// 99--所有币种
// 98--所有外币
// 1	9999	RMB	01	人民币	CNY	C	CHN	1	8	2	2	2	0	0.00	1	0	01	0		0	0	0	0	0
// 2	9999	GBP	12	英镑	GBP	G	GBR	1	8	2	2	2	1	0.00	3	0	12	0		0	0	0	0	0
// 3	9999	HKD	13	港币	HKD	H	HKG	1	8	2	2	2	1	0.00	3	0	13	0		0	0	0	0	0
// 4	9999	USD	14	美元	USD	C	USD	1	8	2	2	2	1	0.00	2	0	14	0		0	0	0	0	0
// 5	9999	CHF	15	瑞士法郎	CHF	S	CHE	1	8	2	2	2	10	0.00	3	0	15	0		0	0	0	0	0
// 6	9999	SGD	18	新加坡元	SGD	S	SGP	1	8	2	2	2	1	0.00	3	0	18	0		0	0	0	0	0
// 7	9999	SEK	21	瑞典克郎	SEK	S	SWE	1	8	2	2	2	1	0.00	3	0	21	0		0	0	0	0	0
// 8	9999	DKK	22	丹麦克朗	DKK	D	DNK	1	8	2	2	2	1	0.00	3	0	22	0		0	0	0	0	0
// 9	9999	NOK	23	挪威克朗	NOK	N	NOR	1	8	2	2	2	1	0.00	3	0	23	0		0	0	0	0	0
// 10	9999	JPY	27	日元	JPY	J	JPN	1	8	0	2	0	1000	0.00	3	0	27	0		0	0	0	0	0
// 11	9999	CAD	28	加拿大元	CAD	C	CAN	1	8	2	2	2	1	0.00	3	0	28	0		0	0	0	0	0
// 12	9999	AUD	29	澳大利亚元	AUD	A	AUS	1	8	2	2	2	1	0.00	3	0	29	0		0	0	0	0	0
// 13	9999	EUR	38	欧元	EUR	E	EUR	1	8	2	2	2	10	0.00	4	0	38	0		0	0	0	0	0
// 14	9999	NZD	87	新西兰元	NZD	S	NZD	1	8	2	2	2	1	0.00	3	0	87	0		0	0	0	0	0

func insertBaseCurrency() {
	fmt.Println("insert doc_currency_table ...")
	const (
		cstStateDONE = "DONE"
	)
	docorgs := []BaseCurrency{
		{Id: 1, Huobdaih: "01", Hbmcheng: "人民币", Bzbiemng: "CNY"},
		{Id: 2, Huobdaih: "12", Hbmcheng: "英镑", Bzbiemng: "GBP"},
		{Id: 3, Huobdaih: "13", Hbmcheng: "港币", Bzbiemng: "HKD"},
		{Id: 4, Huobdaih: "14", Hbmcheng: "美元", Bzbiemng: "USD"},
		{Id: 5, Huobdaih: "15", Hbmcheng: "瑞士法郎", Bzbiemng: "CHF"},
		{Id: 6, Huobdaih: "16", Hbmcheng: "新加坡元", Bzbiemng: "SGD"},
	}
	for _, v := range docorgs {
		u := new(BaseCurrency)
		u.Id = v.Id
		u.Huobdaih = v.Huobdaih

		u.Hbmcheng = v.Hbmcheng
		u.Bzbiemng = v.Bzbiemng

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert doc_currency_table end")
}

func insertPerson() {
	fmt.Println("insert person ...")
	persons := [2]Person{
		{Firstname: "xie", Lastname: "huanang"},
		{Firstname: "smart", Lastname: "devuser"}}
	for _, v := range persons {
		u := new(Person)
		u.Firstname = v.Firstname
		u.Lastname = v.Lastname

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert person end")
}

//getSAPTsList 返回前后1000天的日期字符串
func getSAPTsList() (tsList []string) {
	intDayCount := 3
	cstSuffix := " 00:00:00 000"
	cstlayout := cstLongdatefmt
	tsList = make([]string, 0, intDayCount+intDayCount)
	for intX := 0; intX < intDayCount; intX++ {
		tsLast := getDay(-1 * intX)
		tsCurrent := getDay(1 * intX)

		strTsLast := tsLast.Format(cstlayout) + cstSuffix
		strTsCurrent := tsCurrent.Format(cstlayout) + cstSuffix
		tsList = append(tsList, strTsLast)
		tsList = append(tsList, strTsCurrent)
	}
	sort.Strings(tsList)
	return
}

func insertSAPDayBalFlag() {
	fmt.Println("insert SAPDayBalFlag ...")
	tsList := getSAPTsList()
	flags := make([]SAPDayBalFlag, 0, len(tsList))
	for intX, ts := range tsList {
		foo := SAPDayBalFlag{Id: int64(intX + 1),
			State: "1",
			Ts:    ts}
		flags = append(flags, foo)
	}

	for _, v := range flags {
		u := new(SAPDayBalFlag)
		u.Id = v.Id
		u.State = v.State
		u.Ts = v.Ts

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert SAPDayBalFlag end")
}

// Id int64 `orm:"column(id);pk;auto"`
// //核心组织机构编码
// Orgcode string `orm:"column(org_code);unique;size(32);valid(required)"`
// //是否末端机构
// LeafOrg bool `orm:"column(leaf_org);default(1)"`
// //SAP机构编码
// SAPOrgcode    string `orm:"column(sap_org_code);size(32);default()"`
// ParentOrgcode string `orm:"column(parent_org_code);size(32);"`

func insertSAPOrgmap() {

	fmt.Println("insert SAPOrgmap ...")
	orgmapList := []SAPOrgmap{
		{Orgcode: "10099", LeafOrg: true, SAPOrgcode: "30020", ParentOrgcode: ""},
	}
	for index, v := range orgmapList {
		u := new(SAPOrgmap)
		u.Id = int64(index + 1)
		u.Orgcode = v.Orgcode
		u.LeafOrg = v.LeafOrg
		u.SAPOrgcode = v.SAPOrgcode
		u.ParentOrgcode = v.ParentOrgcode

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert SAPOrgmap end")
}

// SAPAccountTitleCode struct {
// 	Id int64 `orm:"column(id);auto;pk"`
//
// 	//科目代号
// 	AccountTitleCode string `orm:"column(accounttitle_code);unique;size(32);"`
// 	//科目名称
// 	AccountTitleName string `orm:"column(accounttitle_name);unique;size(32);"`
// 	//银行名称
// 	Bankname string `orm:"column(bankname);size(32);"`
//
// 	//排序因子
// 	SortFactor uint64 `orm:"column(sort_factor);size(20);default(0)"`

func insertSAPAccountTitleCode() {

	fmt.Println("insert SAPAccountTitleCode ...")
	orgmapList := []SAPAccountTitleCode{
		{
			SAPAccountTitleCode:     "100201",
			SunlineAccountTitleCode: "100201",
			AccountTitleName:        "银行存款-工商银行",
			Bankname:                "工商银行",
		},
		{
			SAPAccountTitleCode:     "100202",
			SunlineAccountTitleCode: "100202",
			AccountTitleName:        "银行存款-建设银行",
			Bankname:                "建设银行",
		},
		{
			SAPAccountTitleCode:     "100203",
			SunlineAccountTitleCode: "100203",
			AccountTitleName:        "银行存款-中国银行",
			Bankname:                "中国银行",
		},
		{
			SAPAccountTitleCode:     "100204",
			SunlineAccountTitleCode: "100204",
			AccountTitleName:        "银行存款-农业银行",
			Bankname:                "农业银行",
		},
		{
			SAPAccountTitleCode:     "100205",
			SunlineAccountTitleCode: "100205",
			AccountTitleName:        "银行存款-交通银行",
			Bankname:                "交通银行",
		},
		{
			SAPAccountTitleCode:     "100206",
			SunlineAccountTitleCode: "100206",
			AccountTitleName:        "银行存款-民生银行",
			Bankname:                "民生银行",
		},
		{
			SAPAccountTitleCode:     "100207",
			SunlineAccountTitleCode: "100207",
			AccountTitleName:        "银行存款-浦发银行",
			Bankname:                "浦发银行",
		},
		{
			SAPAccountTitleCode:     "100208",
			SunlineAccountTitleCode: "100208",
			AccountTitleName:        "银行存款-兴业银行",
		},
		{
			SAPAccountTitleCode:     "100209",
			SunlineAccountTitleCode: "100209",
			AccountTitleName:        "银行存款-华夏银行",
		},
		{
			SAPAccountTitleCode:     "100210",
			SunlineAccountTitleCode: "100210",
			AccountTitleName:        "银行存款-招商银行",
		},
		{
			SAPAccountTitleCode:     "100211",
			SunlineAccountTitleCode: "100211",
			AccountTitleName:        "银行存款-中信银行",
		},
		{
			SAPAccountTitleCode:     "100212",
			SunlineAccountTitleCode: "100212",
			AccountTitleName:        "银行存款-光大银行",
		},
		{
			SAPAccountTitleCode:     "100213",
			SunlineAccountTitleCode: "100213",
			AccountTitleName:        "银行存款-深圳发展银行",
		},
		{
			SAPAccountTitleCode:     "100214",
			SunlineAccountTitleCode: "100214",
			AccountTitleName:        "银行存款-广东发展银行",
		},

		{
			SAPAccountTitleCode:     "100215",
			SunlineAccountTitleCode: "100215",
			AccountTitleName:        "银行存款-国家开发银行",
		},
		{
			SAPAccountTitleCode:     "100216",
			SunlineAccountTitleCode: "100216",
			AccountTitleName:        "银行存款-城市商业银行",
		},

		{
			SAPAccountTitleCode:     "100217",
			SunlineAccountTitleCode: "100217",
			AccountTitleName:        "银行存款-农村信用合作社",
		},
		{
			SAPAccountTitleCode:     "100220",
			SunlineAccountTitleCode: "100220",
			AccountTitleName:        "银行存款-上海银行",
		},

		{
			SAPAccountTitleCode:     "100221",
			SunlineAccountTitleCode: "100221",
			AccountTitleName:        "银行存款-渣打银行",
		},
		{
			SAPAccountTitleCode:     "100222",
			SunlineAccountTitleCode: "100222",
			AccountTitleName:        "银行存款-邮政储蓄银行",
		},
		{
			SAPAccountTitleCode:     "100218",
			SunlineAccountTitleCode: "100218",
			AccountTitleName:        "银行存款-住房资金中心",
		},
		{
			SAPAccountTitleCode:     "100298",
			SunlineAccountTitleCode: "100298",
			AccountTitleName:        "银行存款-Gopher",
		},
		{
			SAPAccountTitleCode:     "100299",
			SunlineAccountTitleCode: "100299",
			AccountTitleName:        "银行存款-其他银行",
		},
	}
	for index, v := range orgmapList {
		u := new(SAPAccountTitleCode)
		u.Id = int64(index + 1)

		u.SAPAccountTitleCode = v.SAPAccountTitleCode
		u.SunlineAccountTitleCode = v.SunlineAccountTitleCode
		u.AccountTitleName = v.AccountTitleName
		arr := strings.Split(v.AccountTitleName, "-")
		if len(arr) > 1 {
			u.Bankname = arr[1]
		} else {
			u.Bankname = "未定义"
		}

		u.SortFactor = uint64(u.Id)

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert SAPAccountTitleCode end")
}

// BIZ_DATE       VARCHAR2(10),
// ORG_CODE       VARCHAR2(10),
// COIN_TYPE_CODE VARCHAR2(10),
// ACCOUNT_SETS   CHAR(1),
// BAL            NUMBER(20,2)
const (
	//财务帐套
	cstFinancialAccountSet = "1"
	//业务帐套
	cstBizAccountSet = "0"
)

type SAPParam struct {
	SAPOrgcode       string
	AccountTitleCode string
	CoinTypeCode     string
	AccountSets      string
}

type CoinTypeCodeDesc struct {
	CoinTypeCode string
	Desc         string
}

var (
	coinTypeCodeList     []CoinTypeCodeDesc
	accountTitleCodeList []string
	sapOrgcodeList       []string
	accountSetsList      []string
)

func init() {
	sapOrgcodeList = []string{"3200",
		"3201",
		"3210",
		"3211",
		"3212",
		"3220",
		"3221",
		"3222",
		"3223",
		"3224",
		"3230",
		"3231",
		"3232",
		"3233",
		"3234",
		"3235",
		"3236",
		"3240",
		"3241",
		"3242",
		"3243",
		"3244",
		"3250",
		"3251",
		"3252",
		"3253",
		"3254",
		"3255",
		"3261",
		"3262",
		"3263",
		"3264",
		"3265"}

	accountTitleCodeList = []string{
		"1001010000",
		"1003000000",
		"1003000001",
		"1003000002",
		"1011980000",
		"1011980100",
		"1111010000",
		"1111020000",
		"1111030000",
		"1125080000",
		"1125090000",
		"1125980000",
		"1125980001",
		"1125980002",
		"1125980003",
		"1125980004",
		"1125980005",
		"1125980006",
		"1125980007",
		"1125980008",
		"1125980009",
		"1125980010",
		"1132010000",
		"1132030200",
		"1132030400",
		"1132039800",
		"1132039801",
		"1132039802",
		"1132980101",
		"1132980104",
		"1132980105",
		"1132980106",
		"1132989800",
		"1221010100",
		"1221019800",
		"1221020000",
		"1221030100",
		"1221030200",
		"1221030300",
		"1221989801",
		"1221989802",
		"1221989803",
		"1221989804",
		"1221989805",
		"1221989806",
		"1231030000",
		"1231050000",
		"1231090000",
		"1231100000",
		"1231110000",
		"1301010000",
		"1301010001",
		"1301020000",
		"1301020001",
		"1302010000",
		"1302980000",
		"1305010000",
		"1305010000",
		"1305010001",
		"1305010001",
		"1305010002",
		"1305010002",
		"1305010003",
		"1305010003",
		"1305010004",
		"1305010004",
		"1305010005",
		"1305010005",
		"1305010006",
		"1305020000",
		"1305020001",
		"1305020002",
		"1305020003",
		"1305020004",
		"1305020005",
		"1305030000",
		"1306010000",
		"1306020000",
		"1306030000",
		"1306040000",
		"1321010000",
		"1321010100",
		"1321010200",
		"1441010000",
		"1441020000",
		"1441980000",
		"1442010000",
		"1442020000",
		"1442980000",
		"1471000000",
		"1471000001",
		"1471000002",
		"1531020000",
		"1531020001",
		"1531020002",
		"1531020003",
		"1531020004",
		"1531020005",
		"1531020006",
		"1531020007",
		"1531020008",
		"1531980000",
		"1532000000",
		"1611000000",
		"1612000000",
		"2002040100",
		"2002040101",
		"2002040200",
		"2002980100",
		"2002980100",
		"2002980100",
		"2002980101",
		"2002980101",
		"2002980101",
		"2002980200",
		"2002980200",
		"2002980200",
		"2003010000",
		"2003029800",
		"2011010100",
		"2011010100",
		"2011010500",
		"2011020100",
		"2011020101",
		"2011020200",
		"2011020201",
		"2011020202",
		"2011020203",
		"2011020204",
		"2011020205",
		"2011020206",
		"2011020207",
		"2011020208",
		"2011020209",
		"2011020210",
		"2011020211",
		"2011020212",
		"2011020213",
		"2011030100",
		"2011030101",
		"2011030102",
		"2011030103",
		"2011030104",
		"2011030105",
		"2011030106",
		"2011030107",
		"2011030108",
		"2011030109",
		"2011030110",
		"2011030111",
		"2011030112",
		"2011030113",
		"2011030114",
		"2011030115",
		"2012010000",
		"2012020000",
		"2012030000",
		"2012040000",
		"2012980000",
		"2111010000",
		"2111020000",
		"2111030000",
		"2206980000",
		"2206980001",
		"2221010100",
		"2221010100",
		"2221010100",
		"2221010200",
		"2221010300",
		"2221010400",
		"2221010400",
		"2221010500",
		"2221011900",
		"2221019800",
		"2221019800",
		"2221019900",
		"2221019901",
		"2221019902",
		"2221019903",
		"2221019904",
		"2221020000",
		"2231080100",
		"2231080200",
		"2231080300",
		"2231080400",
		"2231080500",
		"2231080600",
		"2231980198",
		"2231980199",
		"2231980200",
		"2231980300",
		"2231980400",
		"2231989800",
		"2241080100",
		"2241089800",
		"2241090100",
		"2241090200",
		"2241090300",
		"2241090400",
		"2241099800",
		"2241320102",
		"2241320298",
		"2241320299",
		"2241980000",
		"2241980001",
		"2241980002",
		"2241980003",
		"2241980004",
		"2243050100",
		"2243050200",
		"2243060100",
		"2243060200",
		"2243070100",
		"2243070200",
		"2243980100",
		"2243980200",
		"2243980300",
		"2243989800",
		"2314010000",
		"2314980100",
		"2701980000",
		"2701980001",
		"2701980002",
		"3002000000",
		"3002010100",
		"3002010200",
		"3002010300",
		"3002020100",
		"3002020200",
		"3002020300",
		"3002030000",
		"4103000000",
		"6011010100",
		"6011010100",
		"6011010101",
		"6011010101",
		"6011010102",
		"6011010200",
		"6011010201",
		"6011010202",
		"6011020000",
		"6011020000",
		"6011030400",
		"6011030401",
		"6011040200",
		"6011050300",
		"6011050400",
		"6011059800",
		"6011980000",
		"6011990000",
		"6021010000",
		"6021020000",
		"6021040000",
		"6021140000",
		"6021989800",
		"6021989801",
		"6021989802",
		"6021989803",
		"6021989804",
		"6021989805",
		"6021989806",
		"6021989807",
		"6021989808",
		"6021989809",
		"6021990000",
		"6041010000",
		"6041010001",
		"6041990000",
		"6061010100",
		"6061010200",
		"6061990000",
		"6301980000",
		"6301980001",
		"6301990000",
		"6411011101",
		"6411012100",
		"6411012101",
		"6411012200",
		"6411012201",
		"6411012202",
		"6411012203",
		"6411012204",
		"6411012205",
		"6411012206",
		"6411012207",
		"6411012208",
		"6411012209",
		"6411012210",
		"6411012211",
		"6411012212",
		"6411013100",
		"6411013101",
		"6411013102",
		"6411013103",
		"6411013104",
		"6411013105",
		"6411013106",
		"6411013107",
		"6411013108",
		"6411013109",
		"6411013110",
		"6411013111",
		"6411013112",
		"6411013113",
		"6411013114",
	}

	coinTypeCodeList = []CoinTypeCodeDesc{
		{"RMB", "人民币"},
		{"USD", "美元（原币）"},
		{"GBP", "英镑（原币）"},
		{"EUR", "欧元（原币）"},
		{"JPY", "日元（原币）"},
		{"HKD", "港币（原币）"},
		// {"Z00", "所有外币（本位币）"},
		// {"Z01", "美元（本位币）"},
		// {"Z02", "英镑（本位币）"},
		// {"Z03", "欧元（本位币）"},
		// {"Z04", "日元（本位币）"},
		// {"Z05", "港币（本位币）"},
		// {"ZZZ", "本外币合计（本位币）"},
	}
	accountSetsList = append(accountSetsList, cstFinancialAccountSet)
	accountSetsList = append(accountSetsList, cstBizAccountSet)
}

// flatSAPParam 提供缺省的参数集合
func getFlatSAPParamList() (sapParamList []SAPParam) {
	sapParamList = make([]SAPParam, 0, 1000)
	log.Printf("%d %d %d %d\n",
		len(sapOrgcodeList),
		len(accountTitleCodeList),
		len(coinTypeCodeList),
		len(accountSetsList))
	//panic(len(sapOrgcodeList) * len(accountTitleCodeList) * len(coinTypeCodeList) * len(accountSetsList))
	for _, sapOrgcode := range sapOrgcodeList {
		for _, accountTitleCode := range accountTitleCodeList {
			for _, coinTypeCode := range coinTypeCodeList {
				for _, accountSets := range accountSetsList {
					foo := SAPParam{SAPOrgcode: sapOrgcode,
						AccountTitleCode: accountTitleCode,
						CoinTypeCode:     coinTypeCode.CoinTypeCode,
						AccountSets:      accountSets}
					sapParamList = append(sapParamList, foo)
				}

			}

		}
	}
	return
}

func insertSAPDayBal() {
	dtNow := time.Now()
	fmt.Println("insert SAPDayBal ...")
	tsList := getSAPTsList()
	paramList := getFlatSAPParamList()
	//log.Printf("tsList %q\n", tsList)
	//panic("tsList")
	bizdateList := make([]string, 0, len(tsList))
	for _, ts := range tsList {
		bizdate := TsToShortdate(ts)
		bizdateList = append(bizdateList, bizdate)
	}
	//panic(strings.Join(bizdateList, "\n"))
	var index int64 = 0
	for _, bizdate := range bizdateList {
		rowList := make([]*SAPDayBal, 0, len(paramList))
		for _, param := range paramList {
			//log.Printf("%d %s %s", index,u.Bizdate,u.SAPOrgcode)
			index += 1
			dblBalance := float64(1.0 * rand.Intn(10000))
			u := new(SAPDayBal)
			u.Id = index
			u.Bizdate = bizdate

			u.SAPOrgcode = param.SAPOrgcode
			u.AccountTitleCode = param.AccountTitleCode
			u.CoinTypeCode = param.CoinTypeCode
			u.AccountSets = param.AccountSets
			u.Balance = dblBalance

			u.Status = 2
			u.Creator = &User{Id: 1}
			u.Updater = &User{Id: 1}
			u.CreateAt = dtNow
			u.UpdateAt = dtNow
			u.Message = "init the data"
			o = orm.NewOrm()
			o.Insert(u)
			rowList = append(rowList, u)
		}

		//o = orm.NewOrm()
		//o.InsertMulti(len(rowList), rowList)
		log.Printf("%d %s", len(rowList), bizdate)
	}

	fmt.Println("insert SAPDayBal end")
}

func insertCity() {
	fmt.Println("insert city ...")

	cityes := [1]FooCity{
		{Name: "Los Angeles", State: "California", Country: "USA"},
	}

	for _, v := range cityes {
		u := new(FooCity)
		u.Name = v.Name
		u.State = v.State
		u.Country = v.Country

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert hotel end")
}

func insertHotel() {
	fmt.Println("insert hotel ...")

	hotels := [4]FooHotel{
		{Name: "Los Angeles International Airport", Address: "Los Angeles International Airport", Zip: "100010"},
		{Name: "Crest Villa Mansion", Address: "Pyin Oo Lwin", Zip: "100010"},
		{Name: "Royal Green Hotel", Address: "Pyin Oo Lwin", Zip: "100020"},
		{Name: "Hotel Pyin Oo Lwin", Address: "Maymyo", Zip: "100020"},
	}

	fooCity := &FooCity{}
	fooCity.Id = 1
	for _, v := range hotels {
		u := new(FooHotel)
		u.Name = v.Name
		u.Address = v.Address
		u.FooCity = fooCity

		u.Status = 2
		u.Creator = &User{Id: 1}
		u.Updater = &User{Id: 1}
		u.CreateAt = time.Now()
		u.UpdateAt = time.Now()
		u.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(u)
	}
	fmt.Println("insert hotel end")
}

func insertGroup() {
	fmt.Println("insert group ...")
	g := new(Group)
	g.Name = "APP"
	g.Title = "System"
	g.Sort = 1
	g.Status = 2
	o.Insert(g)
	fmt.Println("insert group end")
}

func insertRole() {
	fmt.Println("insert role ...")
	roles := []Role{
		{
			Id:         1,
			Name:       "AdminRole",
			Title:      "Admin role",
			IsAdmin:    true,
			IsAduditor: true,
			Remark:     "I'm a admin role"},
		{
			Id:         2,
			Name:       "AuditorRole",
			Title:      "Auditor role",
			IsAdmin:    false,
			IsAduditor: true,
			Remark:     "I'm a auditor role"},
	}
	for _, role := range roles {
		r := new(Role)
		r.Id = role.Id
		r.Name = role.Name
		r.Title = role.Title
		r.IsAdmin = role.IsAdmin
		r.IsAduditor = role.IsAduditor
		r.Remark = role.Remark

		r.Status = 2
		r.Creator = &User{Id: 1}
		r.Updater = &User{Id: 1}
		r.CreateAt = time.Now()
		r.UpdateAt = time.Now()
		r.Message = "init the data"

		o = orm.NewOrm()
		o.Insert(r)
	}
	fmt.Println("insert role end")
}

//insertNodes
func insertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	//nodes := make([20]Node)
	nodes := []Node{
		{Name: "rbac", Title: "RBAC", Remark: "", Level: 1, Pid: 0, Status: 2, Group: g},
		{Name: "node/index", Title: "Node", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "node list", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "AddAndEdit", Title: "add or edit", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "DelNode", Title: "del node", Remark: "", Level: 3, Pid: 2, Status: 2, Group: g},
		{Name: "user/index", Title: "User", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "Index", Title: "user list", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "AddUser", Title: "add user", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "UpdateUser", Title: "update user", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "DelUser", Title: "del user", Remark: "", Level: 3, Pid: 6, Status: 2, Group: g},
		{Name: "group/index", Title: "Group", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "group list", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "AddGroup", Title: "add group", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "UpdateGroup", Title: "update group", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "DelGroup", Title: "del group", Remark: "", Level: 3, Pid: 11, Status: 2, Group: g},
		{Name: "role/index", Title: "Role", Remark: "", Level: 2, Pid: 1, Status: 2, Group: g},
		{Name: "index", Title: "role list", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddAndEdit", Title: "add or edit", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "DelRole", Title: "del role", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "Getlist", Title: "get roles", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AccessToNode", Title: "show access", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddAccess", Title: "add accsee", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "RoleToUserList", Title: "show role to userlist", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		{Name: "AddRoleToUser", Title: "add role to user", Remark: "", Level: 3, Pid: 16, Status: 2, Group: g},
		//
		{Name: "rbac", Title: "Configure", Remark: "", Level: 1, Pid: 0, Status: 2, Group: g},

		{Name: "docname/index", Title: "文档列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "docname list", Remark: "", Level: 3, Pid: 26, Status: 2, Group: g},
		{Name: "AddDocname", Title: "add docname", Remark: "", Level: 3, Pid: 26, Status: 2, Group: g},
		{Name: "UpdateDocname", Title: "update docname", Remark: "", Level: 3, Pid: 26, Status: 2, Group: g},
		{Name: "DelDocname", Title: "del docname", Remark: "", Level: 3, Pid: 26, Status: 2, Group: g},

		{Name: "docgroup/index", Title: "报表组列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "docgroup list", Remark: "", Level: 3, Pid: 31, Status: 2, Group: g},
		{Name: "AddDocgroup", Title: "add docgroup", Remark: "", Level: 3, Pid: 31, Status: 2, Group: g},
		{Name: "UpdateDocgroup", Title: "update docgroup", Remark: "", Level: 3, Pid: 31, Status: 2, Group: g},
		{Name: "DelDocgroup", Title: "del docgroup", Remark: "", Level: 3, Pid: 31, Status: 2, Group: g},

		{Name: "basebizdate/index", Title: "业务日期列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "basebizdate list", Remark: "", Level: 3, Pid: 36, Status: 2, Group: g},
		{Name: "AddBaseBizdate", Title: "add basebizdate", Remark: "", Level: 3, Pid: 36, Status: 2, Group: g},
		{Name: "UpdateBaseBizdate", Title: "update basebizdate", Remark: "", Level: 3, Pid: 36, Status: 2, Group: g},
		{Name: "DelBaseBizdate", Title: "del basebizdate", Remark: "", Level: 3, Pid: 36, Status: 2, Group: g},

		{Name: "basecurrency/index", Title: "货币代码表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "doccurrency list", Remark: "", Level: 3, Pid: 41, Status: 2, Group: g},
		{Name: "AddBasecurrency", Title: "add basecurrency", Remark: "", Level: 3, Pid: 41, Status: 2, Group: g},
		{Name: "UpdateBasecurrency", Title: "update basecurrency", Remark: "", Level: 3, Pid: 41, Status: 2, Group: g},
		{Name: "DelBasecurrency", Title: "del basecurrency", Remark: "", Level: 3, Pid: 41, Status: 2, Group: g},

		{Name: "dpsdetail/index", Title: "存款账户交易明细", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "dpsdetail list", Remark: "", Level: 3, Pid: 46, Status: 2, Group: g},
		{Name: "AddDPSDetail", Title: "add dpsdetail", Remark: "", Level: 3, Pid: 46, Status: 2, Group: g},
		{Name: "UpdateDPSDetail", Title: "update dpsdetail", Remark: "", Level: 3, Pid: 46, Status: 2, Group: g},
		{Name: "DelDPSDetail", Title: "del dpsdetail", Remark: "", Level: 3, Pid: 46, Status: 2, Group: g},

		{Name: "basekehu/index", Title: "客户列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "basekehu list", Remark: "", Level: 3, Pid: 51, Status: 2, Group: g},
		{Name: "AddBasekehu", Title: "add basekehu", Remark: "", Level: 3, Pid: 51, Status: 2, Group: g},
		{Name: "UpdateBasekehu", Title: "update basekehu", Remark: "", Level: 3, Pid: 51, Status: 2, Group: g},
		{Name: "DelBasekehu", Title: "del basekehu", Remark: "", Level: 3, Pid: 51, Status: 2, Group: g},

		{Name: "dpsckzh/index", Title: "存款账户列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "dpsckzh list", Remark: "", Level: 3, Pid: 56, Status: 2, Group: g},
		{Name: "AddDPSckzh", Title: "add dpsckzh", Remark: "", Level: 3, Pid: 56, Status: 2, Group: g},
		{Name: "UpdateDPSckzh", Title: "update dpsckzh", Remark: "", Level: 3, Pid: 56, Status: 2, Group: g},
		{Name: "DelDpsckzh", Title: "del dpsckzh", Remark: "", Level: 3, Pid: 56, Status: 2, Group: g},

		{Name: "dpscklx/index", Title: "存款账户类型列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "dpscklx list", Remark: "", Level: 3, Pid: 61, Status: 2, Group: g},
		{Name: "AddDPScklx", Title: "add dpscklx", Remark: "", Level: 3, Pid: 61, Status: 2, Group: g},
		{Name: "UpdateDPScklx", Title: "update dpscklx", Remark: "", Level: 3, Pid: 61, Status: 2, Group: g},
		{Name: "DelDpscklx", Title: "del dpscklx", Remark: "", Level: 3, Pid: 61, Status: 2, Group: g},

		{Name: "dpsckzhzbx/index", Title: "存款账户指标项列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "dpsckzhzbx list", Remark: "", Level: 3, Pid: 66, Status: 2, Group: g},
		{Name: "AddDPSckzhzbx", Title: "add dpsckzhzbx", Remark: "", Level: 3, Pid: 66, Status: 2, Group: g},
		{Name: "UpdateDPSckzhzbx", Title: "update dpsckzhzbx", Remark: "", Level: 3, Pid: 66, Status: 2, Group: g},
		{Name: "DelDpsckzhzbx", Title: "del dpsckzhzbx", Remark: "", Level: 3, Pid: 66, Status: 2, Group: g},

		{Name: "dpsorgzbx/index", Title: "机构维度指标项列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "dpsDpsorgzbx list", Remark: "", Level: 3, Pid: 71, Status: 2, Group: g},
		{Name: "AddDPSorgzbx", Title: "add dpsorgzbx", Remark: "", Level: 3, Pid: 71, Status: 2, Group: g},
		{Name: "UpdateDPSorgzbx", Title: "update dpsorgzbx", Remark: "", Level: 3, Pid: 71, Status: 2, Group: g},
		{Name: "DelDpsorgzbx", Title: "del dpsorgzbx", Remark: "", Level: 3, Pid: 71, Status: 2, Group: g},

		{Name: "baseexchangerate/index", Title: "汇率列表", Remark: "", Level: 2, Pid: 25, Status: 2, Group: g},
		{Name: "index", Title: "baseexchangerate list", Remark: "", Level: 3, Pid: 76, Status: 2, Group: g},
		{Name: "AddBaseExchangeRate", Title: "add baseexchangerate", Remark: "", Level: 3, Pid: 76, Status: 2, Group: g},
		{Name: "UpdateBaseExchangeRate", Title: "update baseexchangerate", Remark: "", Level: 3, Pid: 76, Status: 2, Group: g},
		{Name: "DelBaseExchangeRate", Title: "del baseexchangerate", Remark: "", Level: 3, Pid: 76, Status: 2, Group: g},
	}
	for _, v := range nodes {
		n := new(Node)
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		o.Insert(n)
	}
	fmt.Println("insert node end")
}
