package main

import (
	"DDN_XS/models/class"
	_ "DDN_XS/routers"
	"encoding/gob"
	"strings"

	"github.com/astaxie/beego"
)

func init() {
	gob.Register(class.User{})
	beego.AddFuncMap("split", SplitHobby)

}
func main() {
	//	beego.SetStaticPath("img", "./views/img")
	//	beego.SetStaticPath("css", "./views/css")
	//	beego.SetStaticPath("js", "./views/js")
	beego.Run()
}

/*	Template Function	*/

func SplitHobby(s string, sep string) []string {
	return strings.Split(s, sep)
}
