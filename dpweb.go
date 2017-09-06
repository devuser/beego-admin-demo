// Copyright 2013 Beego Web authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

// An open source project for official documentation and blog website of beego app framework.
package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"os"

	"github.com/devuser/beego-admin-demo/models"

	"github.com/devuser/beego-admin-demo/routers"

	"fmt"
	_ "github.com/astaxie/beego/session/mysql"
	"github.com/devuser/beego-admin-demo/lib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"mime"
)

const (
	VERSION = "0.1.1"
	APP_VER = "1.0.0"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/dp")
}

// We have to call a initialize function manully
// because we use `bee bale` to pack static resources
// and we cannot make sure that which init() execute first.
func initialize() {
	models.InitModels()

	routers.IsPro = beego.BConfig.RunMode == "prod"
	if routers.IsPro {
		beego.SetLevel(beego.LevelInformational)
		os.Mkdir("./log", os.ModePerm)
		beego.BeeLogger.SetLogger("file", `{"filename": "log/log"}`)
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	///////////////////////////////////////////////////
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "mysql"

	beego.BConfig.WebConfig.Session.SessionProviderConfig = func() string {
		// db_type := beego.AppConfig.String("db_default_type")
		db_host := beego.AppConfig.String("db_default_host")
		db_port := beego.AppConfig.String("db_default_port")
		db_user := beego.AppConfig.String("db_default_user")
		db_pass := beego.AppConfig.String("db_default_pass")
		db_name := beego.AppConfig.String("db_default_name")
		dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		return dns
	}()

	beego.BConfig.WebConfig.Session.SessionOn, _ = beego.AppConfig.Bool("sessionon")
	beego.BConfig.WebConfig.Session.SessionProvider = beego.AppConfig.String("session_provider")
	//beego.BConfig.WebConfig.Session.SessionSavePath = Cfg.MustValue("session", "session_path", "sessions")
	sessionNameUUID := uuid.NewV2(byte(0X08))
	beego.BConfig.WebConfig.Session.SessionName = sessionNameUUID.String()
	// beego.AppConfig.String("session_name")
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 0
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime, _ = beego.AppConfig.Int("session_life_time")
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 86400
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime, _ = beego.AppConfig.Int64("session_life_time")

	// xsrf token expire time
	//过期时间，默认60秒
	beego.BConfig.WebConfig.EnableXSRF = false
	beego.BConfig.WebConfig.XSRFKey = "61oETzKXQAGaYdkL5gEmGeJJFuYh7EQnp2XdTP1o"
	beego.BConfig.WebConfig.XSRFExpire = 86400 * 365
	///////////////////////////////////////////////////
	routers.InitApp()

	mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	initArgs()

	models.Connect()

	beego.AddFuncMap("stringsToJson", lib.StringsToJson)

}

func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			//@todo
			models.Syncdb()
			os.Exit(0)
		}
	}
}

func main() {

	initialize()

	beego.Info(beego.BConfig.AppName, APP_VER)

	beego.InsertFilter("/docs/images/:all", beego.BeforeRouter, routers.DocsStatic)

	if !routers.IsPro {
		beego.SetStaticPath("/static_source", "static_source")
		beego.BConfig.WebConfig.DirectoryIndex = true
	}

	beego.SetStaticPath("/products/images", "products/images/")

	// Register routers.
	beego.Router("/", &routers.HomeRouter{})
	beego.Router("/community", &routers.CommunityRouter{})
	beego.Router("/quickstart", &routers.QuickStartRouter{})
	beego.Router("/video", &routers.VideoRouter{})
	beego.Router("/products", &routers.ProductsRouter{})
	beego.Router("/team", &routers.PageRouter{})
	beego.Router("/about", &routers.AboutRouter{})
	beego.Router("/donate", &routers.DonateRouter{})
	beego.Router("/docs/", &routers.DocsRouter{})
	beego.Router("/docs/*", &routers.DocsRouter{})
	beego.Router("/blog", &routers.BlogRouter{})
	beego.Router("/blog/*", &routers.BlogRouter{})

	// // beego.Router("/", &rbac.MainController{})
	// beego.Router("/admin", &rbac.MainController{}, "*:Index")
	// beego.Router("/public/index", &rbac.MainController{}, "*:Index")
	// beego.Router("/public/login", &rbac.MainController{}, "*:Login")
	// beego.Router("/public/logout", &rbac.MainController{}, "*:Logout")
	// beego.Router("/public/changepwd", &rbac.MainController{}, "*:Changepwd")
	//
	// beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	// beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	// beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	// beego.Router("/rbac/user/index", &rbac.UserController{}, "*:Index")
	//
	// beego.Router("/rbac/node/AddAndEdit", &rbac.NodeController{}, "*:AddAndEdit")
	// beego.Router("/rbac/node/DelNode", &rbac.NodeController{}, "*:DelNode")
	// beego.Router("/rbac/node/index", &rbac.NodeController{}, "*:Index")
	//
	// beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	// beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	// beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	// beego.Router("/rbac/group/index", &rbac.GroupController{}, "*:Index")
	//
	// beego.Router("/rbac/role/AddAndEdit", &rbac.RoleController{}, "*:AddAndEdit")
	// beego.Router("/rbac/role/DelRole", &rbac.RoleController{}, "*:DelRole")
	// beego.Router("/rbac/role/AccessToNode", &rbac.RoleController{}, "*:AccessToNode")
	// beego.Router("/rbac/role/AddAccess", &rbac.RoleController{}, "*:AddAccess")
	// beego.Router("/rbac/role/RoleToUserList", &rbac.RoleController{}, "*:RoleToUserList")
	// beego.Router("/rbac/role/AddRoleToUser", &rbac.RoleController{}, "*:AddRoleToUser")
	// beego.Router("/rbac/role/Getlist", &rbac.RoleController{}, "*:Getlist")
	// beego.Router("/rbac/role/index", &rbac.RoleController{}, "*:Index")
	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run()
}
