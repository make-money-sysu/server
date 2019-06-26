// @APIVersion 1.0.0
// @Title make-money api
// @Description a api for make-money web application
// @Contact 935841375@qq.com
package routers

import (
	"server/controllers"

	"github.com/astaxie/beego"
)

func init() {
	/*
		beego.Router("/", &controllers.MainController{})
		beego.Router("/login", &controllers.LoginController{})
		beego.Router("/user", &controllers.UserController{})
		beego.Router("/survey/?:id", &controllers.SurveyController{})
		beego.Router("/friends", &controllers.FriendsController{})
		beego.Router("/package?:id", &controllers.PackageController{})
		beego.Router("/msg", &controllers.MsgController{})
		beego.Router("/do_survey", &controllers.DoSurveyController{})
	*/
	ns := beego.NewNamespace("/api",

		beego.NSNamespace("/do_survey",
			beego.NSInclude(
				&controllers.DoSurveyController{},
			),
		),

		beego.NSNamespace("/friends",
			beego.NSInclude(
				&controllers.FriendsController{},
			),
		),

		beego.NSNamespace("/msg",
			beego.NSInclude(
				&controllers.MsgController{},
			),
		),

		beego.NSNamespace("/package",
			beego.NSInclude(
				&controllers.PackageController{},
			),
		),

		beego.NSNamespace("/survey",
			beego.NSInclude(
				&controllers.SurveyController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
