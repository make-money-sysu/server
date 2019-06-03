package controllers

import (
	// "fmt"
	"server/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

// PackageController operations for Package
type msgController struct {
	beego.Controller
}

func (this *msgController) Post() {
	var msg models.Msg
	bodyJSON := simplejson.New()
	if inputJson, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		msg.Fromid = inputJson.Get("from").MustInt()
		msg.Toid = inputJson.Get("to").MustInt()
		msg.Content = inputJson.Get("msg").MustString()
		msg.Createtime = time.Now()
		msg.State = 10
		// 0为系统消息 ，10为未查看，11为已查看，但未知悉，12为已查看，已知悉，13为已撤回
	
		result, err := models.SendMessage(&msg)
		if err == nil {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("msg", result)
		} else {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", result)
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid msg format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *msgController) Put() {

}

func (this *msgController) Get() {

}
