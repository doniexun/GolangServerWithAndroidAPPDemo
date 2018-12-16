package main

import (
	_ "github.com/doniexun/GolangServerWithAndroidAPPDemo/ServerDemo/routers"
	"github.com/doniexun/GolangServerWithAndroidAPPDemo/ServerDemo/models"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 注册数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	
	// 连接 MySQL 数据库
	orm.RegisterDataBase("default", "mysql", "root:qrs123456@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	
	// 注册 model
	orm.RegisterModel(new(models.User))
	
	// 建表
	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true
	beego.Run()
}

