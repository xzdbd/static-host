package routers

import (
	"github.com/astaxie/beego/context"

	"github.com/astaxie/beego"
	"github.com/xzdbd/static-host/controllers"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeStatic, func(ctx *context.Context) {
		ctx.Output.Header("Access-Control-Allow-Origin", "*")
	})

	beego.SetStaticPath("/arcgis", "static/arcgis")

	beego.Router("/", &controllers.MainController{})
}
