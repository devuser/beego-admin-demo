package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil,
			nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:AreasController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseBizdateController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseCurrencyController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseExchangeRateController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:BaseKehuController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCklxController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSCkzhzbxController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSDetailController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DPSOrgzbxController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocLogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocaccessController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocgroupController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DocnameController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:DownloadtracController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:PdftracController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/rbac:RolesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

}
