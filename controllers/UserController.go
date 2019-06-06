package controllers

import (
	"encoding/json"
	"fmt"
	"server/models"
	// "strconv"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	
	var user models.User
	bodyJSON := simplejson.New()
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		// fmt.Printf("add user: %+v\n", user)
		_, err = models.AddUser(&user)
		if err == nil {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("msg", "just a msg")
		} else {
			this.Ctx.Output.SetStatus(403)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "this user already registered")
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid user infomation format")
		bodyJSON.Set("err", err)
		fmt.Println(err)
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *UserController) Put() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	
	var user models.User
	bodyJSON := simplejson.New()
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		if nil != this.GetSession("id") && user.Id != this.GetSession("id").(int) {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			err = models.UpdateUserById(&user)
			if err == nil {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", "edited")
			} else {
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "fail")
				bodyJSON.Set("msg", "this user doesn't exist")
			}
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid user infomation format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *UserController) Delete() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	bodyJSON := simplejson.New()

	if nil == this.GetSession("id") {
		this.Ctx.Output.SetStatus(401)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "Login expired")
	}else{
		id :=this.GetSession("id").(int)
		err := models.DeleteUser(id)

		if err == nil {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("msg", "bye~")
		} else {
			this.Ctx.Output.SetStatus(403)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "invalid user")
		}
	}

	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *UserController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	bodyJSON := simplejson.New()
	//id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if this.GetSession("id") != nil {
		id := this.GetSession("id").(int)
		var user *models.User
		user, err := models.GetUserById(id)
		if err == nil {
			bodyJSON.Set("status", "success")
			dataMap := make(map[string]interface{})
			dataMap["id"] = id
			dataMap["real_name"] = user.RealName
			dataMap["nick_name"] = user.NickName
			dataMap["age"] = user.Age
			dataMap["gender"] = user.Gender
			dataMap["head_picture"] = user.HeadPicture
			dataMap["balance"] = user.Balance
			dataMap["profession"] = user.Profession
			dataMap["grade"] = user.Grade
			dataMap["phone"] = user.Phone
			dataMap["email"] = user.Email
			bodyJSON.Set("data", dataMap)
		} else {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "user doesn't exist")
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "Login expired")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}