package routers

//命名空间的路由使用
////初始化 namespace
//ns :=
//beego.NewNamespace("/v1",
//beego.NSCond(func(ctx *context.Context) bool {
//	if ctx.Input.Domain() == "api.beego.me" {
//		return true
//	}
//	return false
//}),
//beego.NSBefore(auth),
//beego.NSGet("/notallowed", func(ctx *context.Context) {
//	ctx.Output.Body([]byte("notAllowed"))
//}),
//beego.NSRouter("/version", &AdminController{}, "get:ShowAPIVersion"),
//beego.NSRouter("/changepassword", &UserController{}),
//beego.NSNamespace("/shop",
//beego.NSBefore(sentry),
//beego.NSGet("/:id", func(ctx *context.Context) {
//	ctx.Output.Body([]byte("notAllowed"))
//}),
//),
//beego.NSNamespace("/cms",
//beego.NSInclude(
//&controllers.MainController{},
//&controllers.CMSController{},
//&controllers.BlockController{},
//),
//),
//)
////注册 namespace
//beego.AddNamespace(ns)