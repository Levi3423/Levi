package routers

import (
	"greeProject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/getusers", &controllers.MainController{},"get:GetUsers")
}
