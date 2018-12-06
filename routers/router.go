package routers

import (
	"DDN_XS/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/join", &controllers.UserController{}, `get:PageJoin`)
	beego.Router("/login", &controllers.UserController{}, `post:Login`)
	beego.Router("/register", &controllers.UserController{}, `post:Register`)
	beego.Router("/logout", &controllers.UserController{}, `get:Logout`)

	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
	beego.Router("/api/user/profile", &controllers.UserController{}, `get:API_Profile`)
}
