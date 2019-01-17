package main

import (
	_ "migo/routers"
	_"migo/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"

)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "data.db")
	orm.RunSyncdb("default", false, true) //程序启动时自动生成数据表
}


func main() {
	beego.Run()
}

