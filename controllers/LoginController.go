package controllers

import (
	"fmt"
	"server/models"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

type LoginController struct {
	beego.Controller
}

//登录相关逻辑
func (this *LoginController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	fmt.Println(this.Ctx.Input.Header("cookie"))
	if inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		id, _ := inputJSON.Get("id").Int()
		passord, _ := inputJSON.Get("password").String()
		success := models.UserLogin(id, passord)
		bodyJSON := simplejson.New()
		if success {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("msg", "success")
			fmt.Println("Set Session ID: ")
			fmt.Println(id)
			this.SetSession("id", id)
		} else {
			this.Ctx.Output.SetStatus(403)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "id and password doesn't match")
		}
		body, _ := bodyJSON.MarshalJSON()
		this.Ctx.Output.Body(body)
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON := simplejson.New()
		bodyJSON.Set("status", "fail")
		bodyJSON.Set("msg", "invalid login format")
		body, _ := bodyJSON.MarshalJSON()
		this.Ctx.Output.Body(body)
	}
}
