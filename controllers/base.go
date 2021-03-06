package controllers

import (
	"DDN_XS/models/class"
	"fmt"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {

	if c.IsLogin() {
		c.Data["user"] = c.GetSession("user").(class.User)
	}
}

func (c *BaseController) DoLogin(u class.User) {
	c.SetSession("user", u)
}
func (c *BaseController) DoLogout() {
	c.DestroySession()
	c.Redirect("/join", 302)
}

func (c *BaseController) IsLogin() bool {
	return c.GetSession("user") != nil
}

func (c *BaseController) CheckLogin() {

	if !c.IsLogin() {
		fmt.Println("未登录")
		c.Redirect("/join", 302)
		c.Abort("302")
	}
}
