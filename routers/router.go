package routers

import (
	"DDN_XS/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("*", &controllers.MainController{})
}
