package routers

import (
	"server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/survey/?:id", &controllers.SurveyController{})
	beego.Router("/friends", &controllers.FriendsController{})
	beego.Router("/package?:id", &controllers.PackageController{})
	beego.Router("/msg", &controllers.MsgController{})
	beego.Router("/do_survey", &controllers.DoSurveyController{})
}
