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

func (this *LoginController) Delete() {
	bodyJSON := simplejson.New()
	//id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	fmt.Println(this.GetSession("id"))
	if this.GetSession("id") != nil {
		// id := this.GetSession("id").(int)
		this.DelSession("id")
		bodyJSON.Set("status", "success")
		bodyJSON.Set("msg", "you have log out")
	} else {
		this.Ctx.Output.SetStatus(200)
		bodyJSON.Set("status", "success")
		bodyJSON.Set("msg", "you have not logined yet")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}