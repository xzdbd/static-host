package routers

import (
	"github.com/astaxie/beego"
	"github.com/xzdbd/static-host/controllers"
)

func init() {
	beego.SetStaticPath("/arcgis", "static/arcgis")

	beego.Router("/", &controllers.MainController{})
}
