package controllers

import (
	_ "DDN_XS/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//fmt.Println("request:", c.Ctx.Request.URL.Path)
	switch c.Ctx.Request.URL.Path {
	case "/":
		c.TplName = "index.html"
		break
	default:
		c.Ctx.WriteString("error..")
		break
	}
}

func (c *MainController) Post() {
	fmt.Println("post request:", c.Ctx.Request.URL.Path)
	switch c.Ctx.Request.URL.Path {
	case "/login":
		uniq := c.GetString("uniq")
		fmt.Println(uniq)
		//操作数据库
		o := orm.NewOrm()
		var level string
		sqlerr := o.Raw("select level from roles where uniq=?", uniq).QueryRow(&level)
		if sqlerr != nil {
			fmt.Println(sqlerr)

			return
		}
		fmt.Println("等级是" + level)
		break
	default:
		c.Ctx.WriteString("error")
		break
	}
}
