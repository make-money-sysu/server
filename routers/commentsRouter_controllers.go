package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["server/controllers:DoSurveyController"] = append(beego.GlobalControllerRouter["server/controllers:DoSurveyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:DoSurveyController"] = append(beego.GlobalControllerRouter["server/controllers:DoSurveyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:LoginController"] = append(beego.GlobalControllerRouter["server/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:LoginController"] = append(beego.GlobalControllerRouter["server/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:PackageController"] = append(beego.GlobalControllerRouter["server/controllers:PackageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:PackageController"] = append(beego.GlobalControllerRouter["server/controllers:PackageController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:PackageController"] = append(beego.GlobalControllerRouter["server/controllers:PackageController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:SurveyController"] = append(beego.GlobalControllerRouter["server/controllers:SurveyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UserController"] = append(beego.GlobalControllerRouter["server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UserController"] = append(beego.GlobalControllerRouter["server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UserController"] = append(beego.GlobalControllerRouter["server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["server/controllers:UserController"] = append(beego.GlobalControllerRouter["server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
