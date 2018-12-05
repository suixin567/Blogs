package controllers

import (
	_ "DDN_XS/models"
	"fmt"

	"os"
	"text/template"

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
	case "/home":
		c.TplName = "xs.html"
		break
	case "/test":
		s1, _ := template.ParseFiles("header.tmpl", "content.tmpl", "footer.tmpl")
		s1.ExecuteTemplate(os.Stdout, "header", nil)
		fmt.Println()
		s1.ExecuteTemplate(os.Stdout, "content", nil)
		fmt.Println()
		s1.ExecuteTemplate(os.Stdout, "footer", nil)
		fmt.Println()
		s1.Execute(os.Stdout, nil)
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
			res := JsonResult{Success: false, Content: "无此用户！"}
			c.Data["json"] = res
			c.ServeJSON()
			return
		}
		fmt.Println("等级是" + level)
		//		c.Data["level"] = level
		//		c.TplName = "xs.tpl" //(文件、文件夹必须小写)
		res := JsonResult{Success: true, Content: level}
		c.Data["json"] = res
		c.ServeJSON()
		break
	default:
		c.Ctx.WriteString("error")
		break
	}
}
