package routers

import (
	"github.com/astaxie/beego"
	"github.com/devuser/beego-admin-demo/rbac"
)

func rbac_router() {
	beego.Router("/admin", &rbac.MainController{}, "*:Index")
	beego.Router("/public/index", &rbac.MainController{}, "*:Index")
	beego.Router("/public/login", &rbac.MainController{}, "*:Login")
	beego.Router("/public/logout", &rbac.MainController{}, "*:Logout")
	beego.Router("/public/changepwd", &rbac.MainController{}, "*:Changepwd")

	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	beego.Router("/rbac/user/index", &rbac.UserController{}, "*:Index")

	beego.Router("/rbac/node/AddAndEdit", &rbac.NodeController{}, "*:AddAndEdit")
	beego.Router("/rbac/node/DelNode", &rbac.NodeController{}, "*:DelNode")
	beego.Router("/rbac/node/index", &rbac.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group/index", &rbac.GroupController{}, "*:Index")

	beego.Router("/rbac/role/AddAndEdit", &rbac.RoleController{}, "*:AddAndEdit")
	beego.Router("/rbac/role/DelRole", &rbac.RoleController{}, "*:DelRole")
	beego.Router("/rbac/role/AccessToNode", &rbac.RoleController{}, "*:AccessToNode")
	beego.Router("/rbac/role/AddAccess", &rbac.RoleController{}, "*:AddAccess")
	beego.Router("/rbac/role/RoleToUserList", &rbac.RoleController{}, "*:RoleToUserList")
	beego.Router("/rbac/role/AddRoleToUser", &rbac.RoleController{}, "*:AddRoleToUser")
	beego.Router("/rbac/role/Getlist", &rbac.RoleController{}, "*:Getlist")
	beego.Router("/rbac/role/index", &rbac.RoleController{}, "*:Index")

	beego.Router("/rbac/docname/AddDocname", &rbac.DocnameController{}, "*:Post")
	beego.Router("/rbac/docname/UpdateDocname", &rbac.DocnameController{}, "*:Put")
	beego.Router("/rbac/docname/DelDocname", &rbac.DocnameController{}, "*:Delete")
	beego.Router("/rbac/docname/index", &rbac.DocnameController{}, "*:Index")

	beego.Router("/rbac/docgroup/AddDocgroup", &rbac.DocgroupController{}, "*:Post")
	beego.Router("/rbac/docgroup/UpdateDocgroup", &rbac.DocgroupController{}, "*:Put")
	beego.Router("/rbac/docgroup/DelDocgroup", &rbac.DocgroupController{}, "*:Delete")
	beego.Router("/rbac/docgroup/index", &rbac.DocgroupController{}, "*:Index")

	beego.Router("/rbac/basecurrency/AddBasecurrency", &rbac.BaseCurrencyController{}, "*:Post")
	beego.Router("/rbac/basecurrency/UpdateBasecurrency", &rbac.BaseCurrencyController{}, "*:Put")
	beego.Router("/rbac/basecurrency/DelBasecurrency", &rbac.BaseCurrencyController{}, "*:Delete")
	beego.Router("/rbac/basecurrency/index", &rbac.BaseCurrencyController{}, "*:Index")

	beego.Router("/rbac/basebizdate/AddBaseBizdate", &rbac.BaseBizdateController{}, "*:Post")
	beego.Router("/rbac/basebizdate/UpdateBaseBizdate", &rbac.BaseBizdateController{}, "*:Put")
	beego.Router("/rbac/basebizdate/DelBaseBizdate", &rbac.BaseBizdateController{}, "*:Delete")
	beego.Router("/rbac/basebizdate/index", &rbac.BaseBizdateController{}, "*:Index")

	beego.Router("/rbac/dpsdetail/AddDPSDetail", &rbac.DPSDetailController{}, "*:Post")
	beego.Router("/rbac/dpsdetail/UpdateDPSDetail", &rbac.DPSDetailController{}, "*:Put")
	beego.Router("/rbac/dpsdetail/DelDPSDetail", &rbac.DPSDetailController{}, "*:Delete")
	beego.Router("/rbac/dpsdetail/index", &rbac.DPSDetailController{}, "*:Index")

	beego.Router("/rbac/basekehu/AddBasekehu", &rbac.BaseKehuController{}, "*:Post")
	beego.Router("/rbac/basekehu/UpdateBasekehu", &rbac.BaseKehuController{}, "*:Put")
	beego.Router("/rbac/basekehu/DelBasekehu", &rbac.BaseKehuController{}, "*:Delete")
	beego.Router("/rbac/basekehu/index", &rbac.BaseKehuController{}, "*:Index")

	beego.Router("/rbac/dpscklx/AddDPSCklx", &rbac.DPSCklxController{}, "*:Post")
	beego.Router("/rbac/dpscklx/UpdateDPSCklx", &rbac.DPSCklxController{}, "*:Put")
	beego.Router("/rbac/dpscklx/DelDPSCklx", &rbac.DPSCklxController{}, "*:Delete")
	beego.Router("/rbac/dpscklx/index", &rbac.DPSCklxController{}, "*:Index")

	beego.Router("/rbac/dpsckzhzbx/AddDPSckzhzbx", &rbac.DPSCkzhzbxController{}, "*:Post")
	beego.Router("/rbac/dpsckzhzbx/UpdateDPSckzhzbx", &rbac.DPSCkzhzbxController{}, "*:Put")
	beego.Router("/rbac/dpsckzhzbx/DelDpsckzhzbx", &rbac.DPSCkzhzbxController{}, "*:Delete")
	beego.Router("/rbac/dpsckzhzbx/index", &rbac.DPSCkzhzbxController{}, "*:Index")

	beego.Router("/rbac/dpsckzh/AddDPSckzh", &rbac.DPSCkzhController{}, "*:Post")
	beego.Router("/rbac/dpsckzh/UpdateDPSckzh", &rbac.DPSCkzhController{}, "*:Put")
	beego.Router("/rbac/dpsckzh/DelDpsckzh", &rbac.DPSCkzhController{}, "*:Delete")
	beego.Router("/rbac/dpsckzh/index", &rbac.DPSCkzhController{}, "*:Index")

	beego.Router("/rbac/dpsorgzbx/AddDPSorgzbx", &rbac.DPSOrgzbxController{}, "*:Post")
	beego.Router("/rbac/dpsorgzbx/UpdateDPSorgzbx", &rbac.DPSOrgzbxController{}, "*:Put")
	beego.Router("/rbac/dpsorgzbx/DelDpsorgzbx", &rbac.DPSOrgzbxController{}, "*:Delete")
	beego.Router("/rbac/dpsorgzbx/index", &rbac.DPSOrgzbxController{}, "*:Index")

	beego.Router("/rbac/baseexchangerate/AddBaseExchangeRate", &rbac.BaseExchangeRateController{}, "*:Post")
	beego.Router("/rbac/baseexchangerate/UpdateBaseExchangeRate", &rbac.BaseExchangeRateController{}, "*:Put")
	beego.Router("/rbac/baseexchangerate/DelBaseExchangeRate", &rbac.BaseExchangeRateController{}, "*:Delete")
	beego.Router("/rbac/baseexchangerate/index", &rbac.BaseExchangeRateController{}, "*:Index")

}
