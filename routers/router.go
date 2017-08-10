package routers

import (
	"xing/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//指定方法
	beego.Router("/getConfig", &controllers.MainController{},"*:GetConfig") //后面第三个参数可以指明对应的方法

	//获取表单
	beego.Router("/getForm", &controllers.MainController{},"*:GetForm")

	//返回json
	beego.Router("/getJson",&controllers.MainController{},"*:ReturnJson")

	//自动注册
	// 	beego.AutoRouter(&controllers.ObjectController{}) //当访问  /object/login   调用 ObjectController 中的 Login 方法

	//静态文件处理
	beego.SetStaticPath("/down1", "download1")//down1是url前缀  download1是目录  http://localhost:8080/down1/test.html


	//使用命名空间
	//APIS
//	ns :=
//		beego.NewNamespace("/api",
//			//此处正式版时改为验证加密请求
//			beego.NSCond(func(ctx *context.Context) bool {
//				if ua := ctx.Input.Request.UserAgent(); ua != "" {
//					return true
//				}
//				return false
//			}),
//			beego.NSNamespace("/ios",
//				//CRUD Create(创建)、Read(读取)、Update(更新)和Delete(删除)
//				beego.NSNamespace("/create",
//					// /api/ios/create/node/
//					beego.NSRouter("/node", &apis.CreateNodeHandler{}),
//					// /api/ios/create/topic/
//					beego.NSRouter("/topic", &apis.CreateTopicHandler{}),
//				),
//				beego.NSNamespace("/read",
//					beego.NSRouter("/node", &apis.ReadNodeHandler{}),
//					beego.NSRouter("/topic", &apis.ReadTopicHandler{}),
//				),
//				beego.NSNamespace("/update",
//					beego.NSRouter("/node", &apis.UpdateNodeHandler{}),
//					beego.NSRouter("/topic", &apis.UpdateTopicHandler{}),
//				),
//				beego.NSNamespace("/delete",
//					beego.NSRouter("/node", &apis.DeleteNodeHandler{}),
//					beego.NSRouter("/topic", &apis.DeleteTopicHandler{}),
//				)
//	),
//)
//
//	beego.AddNamespace(ns)

}