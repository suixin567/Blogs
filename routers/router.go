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

	beego.Router("/setting", &controllers.UserController{}, `get:PageSetting;post:Setting`)

	beego.Router("/article/new", &controllers.ArticleController{}, `get:PageNew;post:New`)
	beego.Router("/article/:id([0-9]+)", &controllers.ArticleController{}, `get:Get`)
	beego.Router("/article/edit/:id([0-9]+)", &controllers.ArticleController{}, `get:PageEdit;post:Edit`)
	//查看个人主页
	beego.Router("/user/:id", &controllers.UserController{}, `get:Profile`)
	//删除文章
	beego.Router("/article/del/:id([0-9]+)", &controllers.ArticleController{}, `get:Del`)
	//根据标签检索文章
	beego.Router("/archive", &controllers.ArticleController{}, `get:Archive`)
	//发表评论
	beego.Router("/reply/new", &controllers.ReplyController{}, `post:New`)
	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
	beego.Router("/api/user/profile", &controllers.UserController{}, `get:API_Profile`)
}
