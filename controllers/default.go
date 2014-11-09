package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "http://wiki.lapingvino.nl/Kial publikigi mian muzikon per soneli.ga"
	this.Data["Email"] = "ikojba+soneli.ga@gmail.com"
	this.TplNames = "index.tpl"
}
