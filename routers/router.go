package routers

import (
	"server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/user/?:id", &controllers.UserController{})
	beego.Router("/survey/?:id", &controllers.SurveyController{})
	beego.Router("/friends", &controllers.FriendsController{})
}
