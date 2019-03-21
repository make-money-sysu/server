package controllers

import (
	"encoding/json"
	"fmt"
	"server/models"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Post() {
	var user models.User
	bodyJSON := simplejson.New()
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		fmt.Printf("add user: %+v\n", user)
		_, err = models.AddUser(&user)
		if err == nil {
			bodyJSON.Set("status", "success")
		} else {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "this user already registered")
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid user infomation format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *UserController) Put() {
	var user models.User
	bodyJSON := simplejson.New()
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &user); err == nil {
		err = models.UpdateUserById(&user)
		if err == nil {
			bodyJSON.Set("status", "success")
		} else {
			bodyJSON.Set("status", "fail")
			bodyJSON.Set("msg", "this user doesn't exist")
		}
	} else {
		bodyJSON.Set("status", "success")
		bodyJSON.Set("msg", "invalid user infomation format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *UserController) Delete() {
	bodyJSON := simplejson.New()
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err == nil {
		err = models.DeleteUser(id)
	}
	if err == nil {
		bodyJSON.Set("status", "success")
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid user id")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *UserController) Get() {
	bodyJSON := simplejson.New()
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err == nil {
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
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "user doesn't exist")
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid user id format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}
