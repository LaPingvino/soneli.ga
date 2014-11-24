package controllers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/astaxie/beego"
	md "github.com/russross/blackfriday"
)

type MainController struct {
	beego.Controller
}

type EkPaĝo struct {
	beego.Controller
}

type MailReceiver struct {
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

func (this *MailReceiver) Get() {
	this.TplNames = "ek.tpl"
	this.Data["Title"] = ""
	this.Data["Contents"] = ""
}

func (this *MailReceiver) Post() {
	this.TplNames = "ek.tpl"
	json := this.GetString("mandrill_events")
	this.Data["Contents"] = json
	if len(json) < 1 {
		return
	}
	beego.Info("mandrill_events arrived")
	beego.Info("Contents: " + json)
	auth := smtp.PlainAuth("", beego.AppConfig.String("mailuser"),
		beego.AppConfig.String("mailauth"),
		strings.Split(beego.AppConfig.String("mailserver"), ":")[0])
	beego.Info("Auth: " + fmt.Sprintln(auth))
	err := smtp.SendMail(beego.AppConfig.String("mailserver"),
		auth, "forward@soneli.ga",
		strings.Split(beego.AppConfig.String("mailto"), ";"), []byte(json))
	if err != nil {
		beego.Error(err.Error())
	}
}
