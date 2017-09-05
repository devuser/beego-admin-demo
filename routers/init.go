package routers

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/compress"
	"github.com/beego/i18n"
	"github.com/devuser/beego-admin-demo/rbac"
	"github.com/howeyc/fsnotify"
)

var (
	CompressConfPath = "conf/compress.json"
)

func initLocales() {
	// Initialized language type list.
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	langTypes = make([]*langType, 0, len(langs))
	for i, v := range langs {
		langTypes = append(langTypes, &langType{
			Lang: v,
			Name: names[i],
		})
	}

	for _, lang := range langs {
		beego.Trace("Loading language: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}

func settingCompress() {
	setting, err := compress.LoadJsonConf(CompressConfPath, IsPro, "/")
	if err != nil {
		beego.Error(err)
		return
	}

	setting.RunCommand()

	if IsPro {
		setting.RunCompress(true, false, true)
	}

	beego.AddFuncMap("compress_js", setting.Js.CompressJs)
	beego.AddFuncMap("compress_css", setting.Css.CompressCss)
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func loadtimes(t time.Time) int {
	return int(time.Now().Sub(t).Nanoseconds() / 1e6)
}

func initTemplates() {
	beego.AddFuncMap("dict", dict)
	beego.AddFuncMap("loadtimes", loadtimes)
}

func InitApp() {
	initTemplates()
	initLocales()
	settingCompress()

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic("Failed start app watcher: " + err.Error())
	}

	go func() {
		for {
			select {
			case event := <-watcher.Event:
				switch filepath.Ext(event.Name) {
				case ".ini":
					beego.Info(event)

					if err := i18n.ReloadLangs(); err != nil {
						beego.Error("Conf Reload: ", err)
					}

					beego.Info("Config Reloaded")

				case ".json":
					if event.Name == CompressConfPath {
						settingCompress()
						beego.Info("Beego Compress Reloaded")
					}
				}
			}
		}
	}()

	if err := watcher.WatchFlags("conf", fsnotify.FSN_MODIFY); err != nil {
		beego.Error(err)
	}
	//////////////////////////////////update by devuser
	//注册rbac路由
	rbac_router()

	//注册RESTful Controller 路由
	//RESTful 是一种目前 API 开发中广泛采用的形式，beego 默认就是支持这样的请求方法，
	//也就是用户 Get 请求就执行 Get 方法，Post 请求就执行 Post 方法。
	//因此默认的路由是这样 RESTful 的请求方式。
	const (
		apiprefix = "/api/v1"
	)
	// beego.Router(/download/ceshi/*“, &rbac.RController{})
	beego.Router(apiprefix+"/docaccess/:uId([0-9]+)/", &rbac.DocaccessController{})
	ns := beego.NewNamespace(apiprefix,

		beego.NSNamespace("/area",
			beego.NSInclude(
				&rbac.AreasController{},
			),
		),

		beego.NSNamespace("/docname",
			beego.NSInclude(
				&rbac.DocnameController{},
			),
		),

		beego.NSNamespace("/downloadtracs",
			beego.NSInclude(
				&rbac.DownloadtracController{},
			),
		),

		beego.NSNamespace("/eventbybps",
			beego.NSInclude(
				&rbac.DocLogController{},
			),
		),

		beego.NSNamespace("/groups",
			beego.NSInclude(
				&rbac.DocgroupController{},
			),
		),

		beego.NSNamespace("/pdftracs",
			beego.NSInclude(
				&rbac.PdftracController{},
			),
		),

		beego.NSNamespace("/roles",
			beego.NSInclude(
				&rbac.RolesController{},
			),
		),
	)
	beego.AddNamespace(ns)

}
