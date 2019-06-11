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
	
	bodyJSON := simplejson.New()
	fmt.Println(this.Ctx.Input.Header("cookie"))
	if inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		id, err1 := inputJSON.Get("id").Int()
		passord, err2 := inputJSON.Get("password").String()
		if err1 != nil || err2 != nil {
			this.Ctx.Output.SetStatus(400)
			bodyJSON.Set("status", "fail")
			bodyJSON.Set("msg", "invalid login format")
		}else{		
			success := models.UserLogin(id, passord)
			if success {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", "success")
				this.SetSession("id", id)
			} else {
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "id and password doesn't match")
			}
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "fail")
		bodyJSON.Set("msg", "invalid login format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}
