package routers

import (
	"github.com/astaxie/beego"
	"github.com/lifeibest/feigo/controllers"
)

func init() {

	beego.Router("/", &controllers.MainController{})

	beego.Router("/post", &controllers.PostController{}, "get:GetAll")
	beego.Router("/post/:id", &controllers.PostController{}, "get:GetOne")
}
