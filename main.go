package main

import (
	_ "server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:wszdwlhw51868@tcp(182.254.206.244:3306)/make_money?charset=utf8")
}
func main() {
	beego.Run()
}
