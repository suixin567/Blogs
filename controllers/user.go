package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {
	c.Data["userid"] = "geek"
	c.Data["tag"] = "i am a fish"
	c.Data["hobby"] = []string{"wan", "you", "xi"}
	c.TplName = "user/profile.html"
}

func (c *UserController) PageJoin() {
	c.TplName = "user/join.html"
}
