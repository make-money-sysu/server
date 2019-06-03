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

// 发送信息
func (this *msgController) Post() {
	var msg models.Msg
	bodyJSON := simplejson.New()
	if inputJson, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		if nil != this.GetSession("id") {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			msg.Fromid = this.GetSession("id").(int)
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
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid msg format")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

// 撤回信息，（只能是未读的） WithdrawalMessage
func (this *msgController) delete() {

	bodyJSON := simplejson.New()
	if inputJson, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		if nil != this.GetSession("id") {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			fromid := this.GetSession("id").(int)
			mid := inputJson.Get("mid").MustInt()
			// 0为系统消息 ，10为未查看，11为已查看，但未知悉，12为已查看，已知悉，13为已撤回
		
			result, err := models.WithdrawalMessage(fromid,mid)
			if err == nil {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", result)
			} else {
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", result)
			}
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid msg format")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

//获取消息, 被获取了，数据库就算已读
func (this *msgController) Get() {
	
	bodyJSON := simplejson.New()
	if inputJson, err := simplejson.NewJson(this.Ctx.Input.RequestBody); err == nil {
		if nil != this.GetSession("id") {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			var readData []models.Msg
			var unreadData []models.Msg
			fromid := this.GetSession("id").(int)
			// mode 0 is history, 1 (not zero) is for change
			if inputJson.Get("history").MustInt() == 0 {
				limit, err := this.GetInt("limit")
				if err != nil {
					limit = -1
				}
				offset, err := this.GetInt("offset")
				if err != nil {
					offset = 0
				}
				toid := inputJson.Get("with").MustInt()
				readData, err = models.GetHistory(fromid,toid, limit, offset)
			}else{
				readData,unreadData, err = models.GetMessage(fromid)
			}

			if err == nil {
				bodyJSON.Set("status", "success")

				tmpMapArr := make([]interface{}, len(readData))
				for i, p := range readData {
					tmpMap := make(map[string]interface{})
					tmpMap["mid"] = p.Mid
					tmpMap["fromid"] = p.Fromid // TODO:: add name
					tmpMap["toid"] = p.Toid
					tmpMap["create_time"] = p.Createtime.String()
					tmpMap["Content"] = p.Content
					tmpMap["state"] = p.State
					tmpMapArr[i] = tmpMap
				}
				bodyJSON.Set("readData", tmpMapArr)

				if inputJson.Get("history").MustInt() != 0 {
					tmpMapArr := make([]interface{}, len(unreadData))
					for i, p := range unreadData {
						tmpMap := make(map[string]interface{})
						tmpMap["mid"] = p.Mid
						tmpMap["fromid"] = p.Fromid // TODO:: add name
						tmpMap["toid"] = p.Toid
						tmpMap["create_time"] = p.Createtime.String()
						tmpMap["Content"] = p.Content
						tmpMap["state"] = p.State
						tmpMapArr[i] = tmpMap
					}
					bodyJSON.Set("unreadData", tmpMapArr)
				}
			} else {
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "get message error,you can try to find us for help")
			}
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid msg format")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}
