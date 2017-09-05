package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

}
