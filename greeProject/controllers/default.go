package controllers

import (
	"fmt"
	"greeProject/models"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) GetUsers() {
	fmt.Println("hello beego!")
	//models.PrintUsers()
	models.PrintUserByORM()

	c.TplName = "index.tpl"
}
