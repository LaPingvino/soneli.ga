package routers

import (
	"github.com/astaxie/beego"
	"github.com/lapingvino/soneli.ga/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/ek", &controllers.EkPaĝo{})
	beego.Router("/mandrill/"+beego.AppConfig.String("mailroute"), &controllers.MailReceiver{})
}
