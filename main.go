package main

import (
	_ "DDN_XS/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("img", "./views/img")
	beego.SetStaticPath("css", "./views/css")
	beego.SetStaticPath("js", "./views/js")
	beego.Run()
}
