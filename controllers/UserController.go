package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/make-money-sysu/server/models"
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
			// this.SetSession("id", id)
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
	
	fmt.Println(this.Ctx.Input.Header("cookie"))

	bodyJSON := simplejson.New()
	if inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		if nil == this.GetSession("id") {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			user, err := models.GetUserById(this.GetSession("id").(int))

			fmt.Println(user)
			fmt.Println(inputJSON)

			if _, ok := inputJSON.CheckGet("real_name");ok{
				user.RealName = inputJSON.Get("real_name").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("password");ok{
				user.Password = inputJSON.Get("password").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("nick_name");ok{
				user.NickName = inputJSON.Get("nick_name").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("age");ok{
				user.Age = inputJSON.Get("age").MustInt()
			}
	
			if _, ok := inputJSON.CheckGet("gender");ok{
				user.Gender = inputJSON.Get("gender").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("head_picture");ok{
				user.HeadPicture = inputJSON.Get("head_picture").MustString()
			}
	
			// if _, ok := inputJSON.CheckGet("balance");ok{
			// 	user.Balance = inputJSON.Get("balance").MustFloat64()
			// }
	
	
			if _, ok := inputJSON.CheckGet("profession");ok{
				user.Profession = inputJSON.Get("profession").MustString()
			}
	
			
			if _, ok := inputJSON.CheckGet("grade");ok{
				user.Grade = inputJSON.Get("grade").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("phone");ok{
				user.Phone = inputJSON.Get("phone").MustString()
			}
	
			if _, ok := inputJSON.CheckGet("email");ok{
				user.Email = inputJSON.Get("email").MustString()
			}
	
		
			
			fmt.Println(user)

			err = models.UpdateUserById(user)
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
	fmt.Println(this.GetSession("id"))
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
