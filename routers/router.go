package routers

import (
	"github.com/lapingvino/soneli.ga/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
