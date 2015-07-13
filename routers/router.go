package routers

import (
	"github.com/astaxie/beegae"
	"github.com/lapingvino/soneli.ga/controllers"
)

func init() {
	beegae.Router("/", &controllers.MainController{})
	beegae.Router("/ek", &controllers.EkPaĝo{})
	beegae.Router("/mandrill/"+beegae.AppConfig.String("mailroute"), &controllers.MailReceiver{})
}
