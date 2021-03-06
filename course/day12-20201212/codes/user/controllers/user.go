package controllers

import (
	"user/forms"
	"user/services"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Add() {
	if c.Ctx.Input.IsPost() {
		var form forms.AddUserForm
		c.ParseForm(&form)
		services.AddUser(form.Name, form.Addr, form.Sex)
		c.Redirect("/user/query", 302)
	} else {
		c.Data["form"] = &forms.AddUserForm{Name: "kk123"}
		c.TplName = "user/add.html"
	}
}

func (c *UserController) Query() {
	c.Data["users"] = services.GetUsers()
	c.TplName = "user/query.html"
}

func (c *UserController) Delete() {
	if id, err := c.GetInt64("id"); err == nil {
		services.DeleteUser(id)
	}
	c.Redirect("/user/query", 302)
}
