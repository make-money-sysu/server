package controllers

import (
	// "fmt"
	// "server/models"
	// "time"

	"github.com/astaxie/beego"
	// "github.com/bitly/go-simplejson"
)

// PackageController operations for Package
type msgController struct {
	beego.Controller
}

func (this *msgController) Post() {
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

func (this *msgController) Put() {

}

func (this *msgController) Get() {

}
