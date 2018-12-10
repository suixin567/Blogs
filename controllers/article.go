package controllers

import (
	"DDN_XS/models/class"
	"fmt"
	"strconv"
	"strings"
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
	c.CheckLogin()
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

	//读取标签
	strs := strings.Split(c.GetString("tag"), ",")
	tags := []*class.Tag{}
	for _, v := range strs {
		tags = append(tags, class.Tag{Name: strings.TrimSpace(v)}.GetOrNew())
	}
	a.Tags = tags
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

//根据一个标签 及一个用户id 获取文章列表
func (c *ArticleController) Archive() {

	errmsg := ""

	a := class.Article{}
	if len(c.GetString("tag")) > 0 {
		tag := class.Tag{Name: c.GetString("tag")}.Get()
		if tag == nil {
			errmsg += fmt.Sprintf("Tag[%s] is not exist.\n", c.GetString("tag"))
		} else {
			a.Tags = []*class.Tag{tag}
		}
	}

	if len(c.GetString("author")) > 0 {
		author := class.User{Id: c.GetString("author")}.Get()
		if author == nil {
			errmsg += fmt.Sprintf("User[%s] is not exist.\n", c.GetString("author"))
		} else {
			a.Author = author
		}
	}

	if len(errmsg) == 0 {
		rets := a.Gets()
		c.Data["articles"] = rets
	}

	c.Data["err"] = errmsg

	c.TplName = "article/archive.html"

}
