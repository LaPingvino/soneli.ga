package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Email"] = "ikojba+soneli.ga@gmail.com"
	this.TplNames = "index.tpl"
}
