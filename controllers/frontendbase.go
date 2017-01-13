/**
 * 基础Controller，方便继续
 */
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lifeibest/feigo/lib"
	"time"
)

type FrontendbaseController struct {
	beego.Controller
	Testsomething int
}

/**
 * 分页
 */
func (this *FrontendbaseController) SetPaginator(per int, nums int64) *lib.Paginator {
	p := lib.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func init() {

	beego.Debug("init")
}

// Prepare implemented Prepare method for FrontendbaseController.
// 这部分数据可以共用，比如缓存，共用组件在这里处理
func (this *FrontendbaseController) Prepare() {
	// Setting properties.
	beego.Debug("Prepare")
	this.Data["Date_y"] = time.Now().Format("2006")

}
