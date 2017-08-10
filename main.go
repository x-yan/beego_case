package main

import (
	_ "xing/routers"
	//"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"fmt"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:123456@/test?charset=utf8")
}

func main() {
		/**
		设置静态资源访问目录
		 */
	beego.SetStaticPath("/down1", "download1")
	/**
	获取app.conf配置文件信息  AppConfig其他方法请查看文档
	 */
	mysqldb := beego.AppConfig.String("mysqldb") //AppConfig  结构体  string方法
	fmt.Println(mysqldb)

	beego.Run()
}

