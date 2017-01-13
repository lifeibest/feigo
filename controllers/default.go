package controllers

// import (
// 	"github.com/astaxie/beego"
// )

type MainController struct {
	//beego.Controller
	FrontendbaseController
}

func (this *MainController) Get() {

	this.Data["Website"] = "go.ffeeii.com"
	this.Data["Email"] = "123178784@qq.com"
	this.Data["IsHome"] = 1
	this.Data["PageTitle"] = "Home - ffeeii"
	this.Data["PageDes"] = "ffeei's blog"

	this.Layout = "layout/default.html"
	this.TplName = "index/index.html"
}
