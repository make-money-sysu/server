package controllers

import (
	"server/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	"fmt"
)

// PackageController operations for Package
type PackageController struct {
	beego.Controller
}

func (this *PackageController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	bodyJSON := simplejson.New()
	var thisPackage models.Package
	packageJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	_, ok1 := packageJSON.CheckGet("reward")
	_, ok2 := packageJSON.CheckGet("note")
	// if _, ok := bodyJSON.CheckGet("status");!ok{
	if err != nil || !ok1 || !ok2{
		//校验格式
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid json format")
	}else if this.GetSession("id") == nil {
		// 检查session
		this.Ctx.Output.SetStatus(401)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "Login expired")
	}else{
		thisPackage.OwnerId, err = models.GetUserById(this.GetSession("id").(int))
		
		thisPackage.CreateTime = time.Now()
		thisPackage.Reward = float32(packageJSON.Get("reward").MustFloat64())
		thisPackage.State = 0
		thisPackage.Note = packageJSON.Get("note").MustString()
		if thisPackage.OwnerId.Balance < thisPackage.Reward {
			this.Ctx.Output.SetStatus(403)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "user balance doesn't enough")
		}else{
			if err == nil {
				_, err := models.AddPackage(&thisPackage)
				if err == nil {
					bodyJSON.Set("status", "success")
					bodyJSON.Set("msg", "post success")
				} else {
					this.Ctx.Output.SetStatus(403)
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", "create the pakage error, please contact with us")
				}
			} else {
				this.Ctx.Output.SetStatus(404)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "this user doesn't not exist")
			}
		}
	}
	
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *PackageController) Put() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	bodyJSON := simplejson.New()
	id, err := this.GetInt("id")
	if err != nil {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("status", "formate error,id is invalid")
	}else{
		if this.GetSession("id") == nil {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			method := this.GetString("method")
			//接单
			if method == "receive" {
				err = models.ReceivePackage(id, this.GetSession("id").(int))
				if err != nil {
					this.Ctx.Output.SetStatus(403)
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", "the package or the user doesn't exist")
				} else {
					bodyJSON.Set("status", "success")
					bodyJSON.Set("msg", "you have recived it")
				}
			} else if method == "confirm" {
				thisPackage, _ := models.GetPackageById(id)
				if thisPackage.OwnerId.Id != this.GetSession("id") {
					this.Ctx.Output.SetStatus(401)
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", "Login expired")
				}else{
					err = models.ConfirmPackage(id)
					if err != nil {
						this.Ctx.Output.SetStatus(403)
						bodyJSON.Set("status", "failed")
						bodyJSON.Set("msg", "create the pakage error, please contact with us")
					} else {
						bodyJSON.Set("status", "success")
						bodyJSON.Set("msg", "confirmed")
					}
				}
			} else {
				this.Ctx.Output.SetStatus(400)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "found no method")
			}
		}
	}
	
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *PackageController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	
	
	id, err := this.GetInt("id")
	if err != nil {
		id = 0
	}
	owner_id, err := this.GetInt("owner_id")
	if err != nil {
		owner_id = 0
	}
	receiver_id, err := this.GetInt("receiver_id")
	if err != nil {
		receiver_id = 0
	}
	state, err := this.GetInt("state")
	if err != nil {
		state = -1
	}
	limit, err := this.GetInt("limit")
	if err != nil {
		limit = -1
	}
	offset, err := this.GetInt("offset")
	if err != nil {
		offset = 0
	}
	packages := models.GetPackages(id, owner_id, receiver_id, state, limit, offset)
	bodyJSON := simplejson.New()
	bodyJSON.Set("status", "success")
	tmpMapArr := make([]interface{}, len(packages))
	for i, p := range packages {
		tmpMap := make(map[string]interface{})
		tmpMap["id"] = p.Id
		tmpMap["owner_id"] = p.OwnerId.Id
		owner, err :=models.GetUserById(p.OwnerId.Id)
		if err == nil {
			fmt.Println(owner.RealName)
			tmpMap["owner_real_name"]=owner.RealName
			tmpMap["owner_nick_name"]=owner.NickName
			tmpMap["owner_Phone"]=owner.Phone
		}else{
			tmpMap["owner_real_name"]="none"
			tmpMap["owner_nick_name"]="none"
			tmpMap["owner_Phone"]="none"
		}



		if p.ReceiverId == nil {
			tmpMap["receiver_id"] = "none"
			tmpMap["receiver_real_name"]="none"
			tmpMap["receiver_nick_name"]="none"
			tmpMap["receiver_Phone"]="none"
		}else{
			tmpMap["receiver_id"] = p.ReceiverId.Id
			receiver, err :=models.GetUserById(p.ReceiverId.Id)
			if err == nil {
				tmpMap["receiver_real_name"]=receiver.RealName
				tmpMap["receiver_nick_name"]=receiver.NickName
				tmpMap["receiver_Phone"]=receiver.Phone
			}else{
				tmpMap["receiver_real_name"]="none"
				tmpMap["receiver_nick_name"]="none"
				tmpMap["receiver_Phone"]="none"
			}
		}
		
		tmpMap["create_time"] = p.CreateTime.String()
		tmpMap["reward"] = p.Reward
		tmpMap["state"] = p.State
		tmpMap["note"] = p.Note
		tmpMapArr[i] = tmpMap
	}
	bodyJSON.Set("data", tmpMapArr)
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}
