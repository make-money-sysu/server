package controllers

import (
	"server/models"
	"time"

	"github.com/astaxie/beego"
	"github.com/bitly/go-simplejson"
)

// PackageController operations for Package
type PackageController struct {
	beego.Controller
}

func (this *PackageController) Post() {
	var thisPackage models.Package
	packageJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if err != nil {
		this.Abort("invalid json format")
	}
	thisPackage.OwnerId, err = models.GetUserById(packageJSON.Get("owner_id").MustInt())
	if thisPackage.OwnerId.Id != this.GetSession("id").(int) {
		this.Abort("Login expired")
	}
	thisPackage.CreateTime = time.Now()
	thisPackage.Reward = float32(packageJSON.Get("reward").MustFloat64())
	thisPackage.State = 0
	thisPackage.Note = packageJSON.Get("note").MustString()
	if thisPackage.OwnerId.Balance < thisPackage.Reward {
		this.Abort("user balance doesn't enough")
	}
	bodyJSON := simplejson.New()
	if err == nil {
		_, err := models.AddPackage(&thisPackage)
		if err == nil {
			bodyJSON.Set("status", "success")
		} else {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", err.Error())
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "this user doesn't not exist")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *PackageController) Put() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Abort("invalid id")
	}
	method := this.GetString("method")
	bodyJSON := simplejson.New()
	//接单
	if method == "receive" {
		receiver_id, _ := this.GetInt("receiver_id")
		if this.GetSession("id").(int) != receiver_id {
			this.Abort("Login expired")
		}
		err = models.ReceivePackage(id, receiver_id)
		if err != nil {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "the package or the user doesn't exist")
		} else {
			bodyJSON.Set("status", "success")
		}
	} else if method == "confirm" {
		thisPackage, _ := models.GetPackageById(id)
		if this.GetSession("id").(int) != thisPackage.OwnerId.Id {
			this.Abort("Login expired")
		}
		err = models.ConfirmPackage(id)
		if err != nil {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", err.Error())
		} else {
			bodyJSON.Set("status", "success")
		}
	} else {
		this.Abort("invalid method")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *PackageController) Get() {
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
		tmpMap["receiver_id"] = p.ReceiverId.Id
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
