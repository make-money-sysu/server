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
	// Cross
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	if inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		id, _ := inputJSON.Get("id").Int()
		passord, _ := inputJSON.Get("password").String()
		fmt.Println(id)
		fmt.Println(passord)
		success := models.UserLogin(id, passord)
		bodyJSON := simplejson.New()
		if success {
			bodyJSON.Set("status", "success")
			this.SetSession("id", id)
		} else {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "id and password doesn't match")
		}
		body, _ := bodyJSON.MarshalJSON()
		this.Ctx.Output.Body(body)
	} else {
		fmt.Println(err.Error())
		bodyJSON := simplejson.New()
		bodyJSON.Set("status", "fail")
		bodyJSON.Set("msg", "invalid login format")
		body, _ := bodyJSON.MarshalJSON()
		this.Ctx.Output.Body(body)
	}
}
