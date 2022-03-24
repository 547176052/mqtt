package mqtt

import (
	"github.com/547176052/mqtt/controller"
	"github.com/547176052/mqtt/guard"
	language2 "github.com/547176052/mqtt/modules/language"
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/modules/service"
	"github.com/GoAdminGroup/go-admin/modules/utils"
	"github.com/GoAdminGroup/go-admin/plugins"
	form2 "github.com/GoAdminGroup/go-admin/plugins/admin/modules/form"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
)

type Mqtt struct {
	*plugins.Base

	handler *controller.Handler
	guard   *guard.Guardian

	// ...
}

func init() {
	plugins.Add(&Mqtt{
		Base: &plugins.Base{PlugName: "mqtt", URLPrefix: "mqtt"},
		// ....
	})
}

func NewMqtt( /*...*/ ) *Mqtt {
	return &Mqtt{
		Base: &plugins.Base{PlugName: "mqtt", URLPrefix: "mqtt"},
		// ...
	}
}

func (plug *Mqtt) IsInstalled() bool {
	// ... implement it
	return true
}

func (plug *Mqtt) GetIndexURL() string {
	return config.Url("/mqtt/example?param=helloworld")
}

func (plug *Mqtt) InitPlugin(srv service.List) {

	// DO NOT DELETE
	plug.InitBase(srv, "mqtt")

	plug.handler = controller.NewHandler( /*...*/ )
	plug.guard = guard.New( /*...*/ )
	plug.App = plug.initRouter(srv)
	plug.handler.HTML = plug.HTML
	plug.handler.HTMLMenu = plug.HTMLMenu

	language.Lang[language.CN].Combine(language2.CN)
	language.Lang[language.EN].Combine(language2.EN)

	plug.SetInfo(info)
}

var info = plugins.Info{
	Website:     "",
	Title:       "Mqtt",
	Description: "",
	Version:     "",
	Author:      "",
	Url:         "",
	Cover:       "",
	Agreement:   "",
	Uuid:        "",
	Name:        "",
	ModulePath:  "",
	CreateDate:  utils.ParseTime("2000-01-11"),
	UpdateDate:  utils.ParseTime("2000-01-11"),
}

func (plug *Mqtt) GetSettingPage() table.Generator {
	return func(ctx *context.Context) (mqttConfiguration table.Table) {

		cfg := table.DefaultConfigWithDriver(config.GetDatabases().GetDefault().Driver)

		if !plug.IsInstalled() {
			cfg = cfg.SetOnlyNewForm()
		} else {
			cfg = cfg.SetOnlyUpdateForm()
		}

		mqttConfiguration = table.NewDefaultTable(cfg)

		formList := mqttConfiguration.GetForm().
			AddXssJsFilter().
			HideBackButton().
			HideContinueNewCheckBox().
			HideResetButton()

		// formList.AddField(...)

		formList.SetInsertFn(func(values form2.Values) error {
			// TODO: finish the installation
			return nil
		})

		formList.EnableAjaxData(types.AjaxData{
			SuccessTitle:   language2.Get("install success"),
			ErrorTitle:     language2.Get("install fail"),
			SuccessJumpURL: "...",
		}).SetFormNewTitle(language2.GetHTML("mqtt installation")).
			SetTitle(language2.Get("mqtt installation")).
			SetFormNewBtnWord(language2.GetHTML("install"))

		return
	}
}
