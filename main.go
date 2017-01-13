package main

import (
	"github.com/astaxie/beego"
	_ "github.com/lifeibest/feigo/routers"
	"github.com/lifeibest/feigo/setting" //前面点号，引用的时候就省略了前面的名字
)

//常量定义
const (
	DOMAIN  = "go.ffeeii.cm" //域名
	VERSION = "1.0"          //版本号

)

func init() {
	//初始化
	setting.Initialize()

}

func main() {
	beego.Run()
}
