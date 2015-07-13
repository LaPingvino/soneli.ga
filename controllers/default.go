package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/astaxie/beegae"
	md "github.com/russross/blackfriday"
)

type MainController struct {
	beegae.Controller
}

type EkPaĝo struct {
	beegae.Controller
}

type MailReceiver struct {
	beegae.Controller
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
		beegae.Error(err.Error())
	}
	defer site.Body.Close()
	body, err := ioutil.ReadAll(site.Body)
	if err != nil {
		beegae.Error(err.Error())
	}
	this.Data["Contents"] = template.HTML(md.MarkdownBasic(body))
}

func (this *MailReceiver) Get() {
	this.TplNames = "ek.tpl"
	this.Data["Title"] = ""
	this.Data["Contents"] = ""
}

type MandrillJSON []struct {
	Msg struct {
		Raw_Msg string
	}
}

func (this *MailReceiver) Post() {
	this.TplNames = "ek.tpl"
	getjson := this.GetString("mandrill_events")
	this.Data["Contents"] = getjson
	if len(getjson) < 1 {
		return
	}
	var mj MandrillJSON
	err := json.Unmarshal([]byte(getjson), &mj)
	if err != nil {
		beegae.Error(err.Error())
	}
	beegae.Info("mandrill_events arrived")
	beegae.Info("Contents: " + getjson)
	beegae.Info("Raw message: " + mj[0].Msg.Raw_Msg)
	auth := smtp.PlainAuth("", beegae.AppConfig.String("mailuser"),
		beegae.AppConfig.String("mailauth"),
		strings.Split(beegae.AppConfig.String("mailserver"), ":")[0])
	beegae.Info("Auth: " + fmt.Sprintln(auth))
	err = smtp.SendMail(beegae.AppConfig.String("mailserver"),
		auth, "forward@soneli.ga",
		strings.Split(beegae.AppConfig.String("mailto"), ";"), []byte(mj[0].Msg.Raw_Msg))
	if err != nil {
		beegae.Error(err.Error())
	}
}
