package routers

import (
	"server/controllers"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowOrigins:     []string{"http://localhost:8080"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
        AllowCredentials: true,
    }))

	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/survey/?:id", &controllers.SurveyController{})
	beego.Router("/friends", &controllers.FriendsController{})
	beego.Router("/package?:id", &controllers.PackageController{})
	beego.Router("/msg", &controllers.MsgController{})
	beego.Router("/do_survey", &controllers.DoSurveyController{})
}
