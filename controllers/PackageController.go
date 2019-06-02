package controllers

import (
	"server/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
	// "fmt"
)

// PackageController operations for Package
type PackageController struct {
	beego.Controller
}

func (this *PackageController) Post() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	bodyJSON := simplejson.New()
	var thisPackage models.Package
	packageJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if err != nil {
		//校验格式
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid json format")
	}else if this.GetSession("id") == nil {
		// 检查session
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "Login expired")
	}else{
		thisPackage.OwnerId, err = models.GetUserById(this.GetSession("id").(int))
		
		thisPackage.CreateTime = time.Now()
		thisPackage.Reward = float32(packageJSON.Get("reward").MustFloat64())
		thisPackage.State = 0
		thisPackage.Note = packageJSON.Get("note").MustString()
		if thisPackage.OwnerId.Balance < thisPackage.Reward {
			this.Abort("user balance doesn't enough")
		}
		if err == nil {
			_, err := models.AddPackage(&thisPackage)
			if err == nil {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", "post success")
			} else {
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "create the pakage error, please contact with us")
			}
		} else {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "this user doesn't not exist")
		}
	}
	
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *PackageController) Put() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	id, err := this.GetInt("id")
	if err != nil {
		this.Abort("invalid id")
	}
	method := this.GetString("method")
	bodyJSON := simplejson.New()
	//接单
	if method == "receive" {
		// this.GetSession("id").(int)
		if this.GetSession("id") == nil {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			err = models.ReceivePackage(id, this.GetSession("id").(int))
			if err != nil {
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "the package or the user doesn't exist")
			} else {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", "you have recived it")
			}
		}
	} else if method == "confirm" {
		if this.GetSession("id") == nil {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			thisPackage, _ := models.GetPackageById(id)
			if this.GetSession("id") == nil || thisPackage.OwnerId.Id != this.GetSession("id") {
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "Login expired")
			}else{
				err = models.ConfirmPackage(id)
				if err != nil {
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", "create the pakage error, please contact with us")
				} else {
					bodyJSON.Set("status", "success")
					bodyJSON.Set("msg", "confirmed")
				}
			}
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "found no method")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *PackageController) Get() {
	// this.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	
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
		// fmt.Println(p.ReceiverId)
		tmpMap["owner_id"] = p.OwnerId.Id
		if p.ReceiverId == nil {
			tmpMap["receiver_id"] = "none"
		}else{
			tmpMap["receiver_id"] = p.ReceiverId.Id
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
