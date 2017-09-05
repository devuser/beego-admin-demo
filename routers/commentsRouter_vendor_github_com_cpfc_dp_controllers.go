package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:AreasController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocgroupController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocLogController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocaccessController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DocnameController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:DownloadtracController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:PdftracController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"],
		beego.ControllerComments{
			"Put",
			`/:id`,
			[]string{"put"},
			nil, nil})

	beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"] = append(beego.GlobalControllerRouter["vendor/github.com/devuser/beego-admin-demo/controllers:RolesController"],
		beego.ControllerComments{
			"Delete",
			`/:id`,
			[]string{"delete"},
			nil, nil})

}
