package controllers

import (
	//	"DDNWeb/models"
	"fmt"

	"github.com/astaxie/beego"
)

type XsController struct {
	beego.Controller
}

func (c *XsController) Get() {
	//fmt.Println("request:", c.Ctx.Request.URL.Path)
	switch c.Ctx.Request.URL.Path {
	case "/xs":
		c.TplName = "sample/xs_login.html"
		break
		//	case "/xs_login":
		//		c.TplName = "sample/xs.html"
		//		break
	default:
		c.Ctx.WriteString("error")
		break
	}
}

func (c *XsController) Post() {
	fmt.Println("post request:", c.Ctx.Request.URL.Path)
	switch c.Ctx.Request.URL.Path {
	case "/login": //登录
		uniq := c.GetString("uniq")
		fmt.Println("识别码" + uniq)
		break

		//	case "/xs": //销售专用网页
		//		loudongshu := c.GetString("loudongshu")
		//		fmt.Println("楼栋数" + loudongshu)
		//		danyuanshu := c.GetString("danyuanshu")
		//		fmt.Println("单元数" + danyuanshu)
		//		break
	default:
		c.Ctx.WriteString("error")
		break
	}
}
