package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["xing/controllers:CMSController"] = append(beego.GlobalControllerRouter["xing/controllers:CMSController"],
		beego.ControllerComments{
			Method: "StaticBlock",
			Router: `/staticblock/:key`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["xing/controllers:CMSController"] = append(beego.GlobalControllerRouter["xing/controllers:CMSController"],
		beego.ControllerComments{
			Method: "Product",
			Router: `/products`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
