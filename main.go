package main

import (
	"github.com/astaxie/beego"
	_ "github.com/xzdbd/static-host/routers"
)

func main() {
	beego.Run()
}
