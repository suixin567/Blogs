package controllers

import (
	"DDN_XS/models/class"
	"fmt"
	"time"

	"github.com/astaxie/beego/validation"
)

func (c *UserController) API_Profile() {
	type user struct {
		Userid string
		Hobby  []string
	}
	u := user{
		"jike",
		[]string{"aaa", "bbb"},
	}
	c.Data["json"] = u
	c.ServeJSON()
}

type RET struct {
	Ok      bool        `json:"success"`
	Content interface{} `json:"content"`
}

func (c *UserController) Register() {

	ret := RET{
		Ok:      true,
		Content: "success",
	}

	defer func() {
		c.Data["json"] = ret
		c.ServeJSON()
	}()

	id := c.GetString("userid")
	nick := c.GetString("nick")
	pwd1 := c.GetString("password")
	pwd2 := c.GetString("password2")
	email := c.GetString("email")

	if len(nick) < 1 {
		nick = id
	}
	//合法性验证
	valid := validation.Validation{}
	valid.Email(email, "Email")
	valid.Required(id, "Userid")
	valid.Required(pwd1, "Password")
	valid.Required(pwd2, "Password2")

	valid.MaxSize(id, 20, "UserID")
	valid.MaxSize(nick, 30, "Nick")

	if pwd1 != pwd2 {
		valid.Error("两次密码不一致")
	}

	if valid.HasErrors() {
		fmt.Println("初步验证有错误")
		ret.Ok = false
		ret.Content = valid.Errors[0].Key + valid.Errors[0].Message
		return
	}
	u := &class.User{
		Id:       0,
		Email:    email,
		Nick:     nick,
		Password: pwd1,
		Regtime:  time.Now(),
	}

	if u.ExistId() {
		valid.Error("用户名被占用")
	}
	if u.ExistEmail() {
		valid.Error("邮箱被占用")
	}
	if valid.HasErrors() {
		fmt.Println("进一步验证有错误")
		ret.Ok = false
		ret.Content = valid.Errors[0].Key + valid.Errors[0].Message
		return
	}

	err := u.Create()
	if err != nil {
		valid.Error(fmt.Sprintf("%v", err))
	}

	if valid.HasErrors() {
		fmt.Println("最终创建有错误")
		ret.Ok = false
		ret.Content = valid.Errors[0].Key + valid.Errors[0].Message
	}

}

func (c *UserController) Login() {
	c.ServeJSON()
}
