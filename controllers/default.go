package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
	"xing/models/users"
)

type MainController struct {
	beego.Controller
}


/**
  crontroller  中的方法执行顺序
   （1）Init(ct *context.Context, childName string, app interface{})
   （2）Prepare() 在get等http方法之前执行，用于初始化一些东西
   （3）Get()
     。。。。。。
    （N）Finish() 这个函数是在执行完相应的 HTTP Method 方法之后执行的
    （last）Render() error 这个函数主要用来实现渲染模板，如果 beego.AutoRender 为 true 的情况下才会执行。


 */




/**
   会在所有定义方法之前执行 验证可以写在这里
 */
func  (c *MainController) Prepare()  {
      fmt.Println("begin  xing")  //会打印在cmd界面
}
/**
如果用户请求的 HTTP Method 是 GET，那么就执行该函数，默认是 405，用户继承的子 struct 中可以实现了该方法以处理 Get 请求。
 */
func (c *MainController) Get() {

	o := orm.NewOrm()

	user := users.Users{Id:219857}
	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Title)
	}

	/**
	 模板渲染技巧

	 mp["name"]="astaxie"
	 mp["nickname"] = "haha"
	 this.Data["m"]=mp

	the username is {{.m.name}}
	the username is {{.m.nickname}}
	 */

	c.Data["Website"] = "beego.me"//拿来存储赋值到模板的值
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"//文件 文件夹必须小写  默认支持 tpl 和 html后缀

}

/**
 用来获取配置信息
 */
func (c *MainController) GetConfig() {
    fmt.Println(beego.AppConfig.String("mysqlurls"))//获取普通的字符
	/**
	[dev]
    httpport = 8080
	 */
	fmt.Println(beego.AppConfig.String("dev::httpport"))//获[注解下的值] 8080

	//系统默认参数
	//beego.BConfig保存了系统的所有默认参数

	//加载自定义配置文件
	beego.LoadAppConfig("ini", "conf/app2.conf") //为何第一个参数是ini，尚且不清楚
	fmt.Println(beego.AppConfig.String("APP2"))//获取app2.conf下的值 121312sdfa3

		c.Ctx.WriteString("hello world")

}

/**
  获取表单的值
 */

//试验失败
type user struct {
	Id    int         `form:"-"`
	Name  interface{} `form:"username"`
	Age   int         `form:"age"`
	Email string
}

func(c *MainController) GetForm(){

	u := user{}
	fmt.Println(c.Input().Get("username")) //这个GET方法不分get还是post都是用这个获取

	if err := c.ParseForm(&u); err != nil {
		//handle error
		fmt.Println("\r\n")
		fmt.Println(u)
	}
	c.TplName = "test_form.tpl"

}


/**
 json输出
 */

func(this *MainController) ReturnJson(){
	mystruct := user{123,"1321212",123,"adhahsd@qq.com" }
	this.Data["json"] = &mystruct
	this.ServeJSON()
/**
{
    "Id": 123,
    "Name": "1321212",
    "Age": 123,
    "Email": "adhahsd@qq.com"
}
 */
}