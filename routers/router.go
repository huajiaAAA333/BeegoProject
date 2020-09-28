package routers

import (
	"BeeGoDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
	
    beego.Router("/login", &controllers.MainController{})
}
