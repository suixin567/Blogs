package main

import (
	"DDN_XS/models/class"
	_ "DDN_XS/routers"
	"encoding/gob"

	"github.com/astaxie/beego"
)

func init() {
	gob.Register(class.User{})
}
func main() {
	//	beego.SetStaticPath("img", "./views/img")
	//	beego.SetStaticPath("css", "./views/css")
	//	beego.SetStaticPath("js", "./views/js")
	beego.Run()
}
