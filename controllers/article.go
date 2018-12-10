package controllers

import (
	"DDN_XS/models/class"
	"fmt"
	"strconv"
)

type ArticleController struct {
	BaseController
	ret RET
}

func (c *ArticleController) PageNew() {
	c.CheckLogin()
	c.TplName = "article/new.html"
}

func (c *ArticleController) Get() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()
	c.Data["article"] = a
	c.TplName = "article/article.html"
}

func (c *ArticleController) PageEdit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()
	a.Author.ReadDB()
	c.Data["article"] = a
	c.TplName = "article/edit.html"
}

func (c *ArticleController) Edit() {
	c.CheckLogin()
	u := c.GetSession("user").(class.User)

	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()

	if u.Id != a.Author.Id {
		c.DoLogout()
	}

	a.Title = c.GetString("title")
	a.Content = c.GetString("content")

	a.Update()

	c.ret.Ok = true
	c.Data["json"] = c.ret
	c.ServeJSON()
}

//这里没有做数据验证 todo
func (c *ArticleController) New() {
	fmt.Println("创建文章")
	c.CheckLogin()
	u := c.GetSession("user").(class.User)
	a := &class.Article{
		Title:   c.GetString("title"),
		Content: c.GetString("content"),
		Author:  &u,
	}

	n, err := a.Create()

	if err == nil {
		c.ret.Ok = true
		c.ret.Content = n //创建成功则返回id
		c.Data["json"] = c.ret
		c.ServeJSON()
		return
	}

	c.ret.Content = fmt.Sprint(err)

	c.Data["json"] = c.ret
	c.ServeJSON()
}

func (c *ArticleController) Del() {
	c.CheckLogin()
	u := c.GetSession("user").(class.User)

	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	a := &class.Article{Id: id}
	a.ReadDB()

	if u.Id != a.Author.Id {
		c.DoLogout()
	}

	a.Defunct = true
	a.Update()

	c.Redirect("/user/"+a.Author.Id, 302)
}
