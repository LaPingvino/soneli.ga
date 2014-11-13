package controllers

import (
	"io/ioutil"
	"net/http"
	"html/template"

	"github.com/astaxie/beego"
	md "github.com/russross/blackfriday"
)

type MainController struct {
	beego.Controller
}

type EkPaĝo struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "http://wiki.lapingvino.nl/Kial publikigi mian muzikon per soneli.ga"
	this.Data["Email"] = "ikojba+soneli.ga@gmail.com"
	this.TplNames = "index.tpl"
}

func (this *EkPaĝo) Get() {
	this.TplNames = "ek.tpl"
	this.Data["Title"] = "Testo de la ekpaĝo"
	site, err := http.Get("http://wiki.lapingvino.nl/_showraw/Kial%20publikigi%20mian%20muzikon%20per%20soneli.ga")
	if err != nil {
		this.Data["Contents"] = "An error occured: " + err.Error()
		beego.Error(err.Error())
	}
	defer site.Body.Close()
	body, err := ioutil.ReadAll(site.Body)
	if err != nil {
		beego.Error(err.Error())
	}
	this.Data["Contents"] = template.HTML(md.MarkdownBasic(body))
}
