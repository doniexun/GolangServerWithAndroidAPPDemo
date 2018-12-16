package routers

import (
	"github.com/doniexun/GolangServerWithAndroidAPPDemo/ServerDemo/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.BaseController{})
    beego.Router("/register", &controllers.UserController{}, "post:Register")
}
