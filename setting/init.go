package setting

import (
	"blog.ffeeii.com/lib"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"time"
)

var Langs []string

// 初始化
func Initialize() {
	Initlog()         //初始化日志
	Initdb()          //初始化db
	InitAddViewFunc() //初始化注册到view函数
	settingLocales()  //本地化多语言
}

//初始化db
func Initdb() {
	var dns string
	db_type := beego.AppConfig.String("db_type")
	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")

	dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
	//beego.Debug(dns)

	orm.RegisterDriver("mysql", orm.DRMySQL)
	//orm.RegisterDataBase("default", "mysql", "root:root@localhost/ydllocalhost?charset=utf8", 30)
	orm.RegisterDataBase("default", db_type, dns, 200)
	//orm.RunSyncdb("default", false, true)
	//
	orm.Debug = true //设置 Debug 为 true 打印查询的语句,线上关闭
}

//数据库连接
// func GetLink() beedb.Model {
// 	//beedb.OnDebug = true
// 	db, err := sql.Open("mysql", db_user+":"+db_pass+"@tcp("+db_host+":"+db_port+")/"+db_name)
// 	if err != nil {
// 		panic(err)
// 	}

// 	orm := beedb.New(db)
// 	return orm
// }

//初始化日志，目前日志是否按照时间生成？
func Initlog() {
	beego.Info("loglog")
	//设置日志 http://beego.me/docs/mvc/controller/logs.md
	now := time.Now().Format("2006_01_02")
	fn := "/tmp/blog.ffeeii.com." + now + ".log"
	beego.SetLogger("file", `{"filename":"`+fn+`"}`)
}

/**
 * [InitAddViewFunc 初始化注册到view函数]
 */
func InitAddViewFunc() {
	beego.AddFuncMap("Strtomd5", lib.Strtomd5) //lib 注册函数到view模板的办法，可以继续增加
	beego.AddFuncMap("i18n", i18n.Tr)          //https://beego.me/docs/module/i18n.md
}

/**
 * 本地化多语言
 * @return {[type]} [description]
 */
func settingLocales() {
	// load locales with locale_LANG.ini files
	langs := "en-US|zh-CN"
	for _, lang := range strings.Split(langs, "|") {
		lang = strings.TrimSpace(lang)
		files := []string{"conf/" + "locale_" + lang + ".ini"}
		if fh, err := os.Open(files[0]); err == nil {
			fh.Close()
		} else {
			files = nil
		}
		beego.Debug(lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini", files...); err != nil {
			beego.Error("Fail to set message file: " + err.Error())
			os.Exit(2)
		}
	}
	Langs = i18n.ListLangs()
}
