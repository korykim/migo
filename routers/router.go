package routers

import (
	"migo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/hello",&controllers.HelloController{})
}
