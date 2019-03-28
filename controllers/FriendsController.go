package controllers

import (
	"server/models"

	"github.com/bitly/go-simplejson"

	"github.com/astaxie/beego"

	"fmt"
)

type FriendsController struct {
	beego.Controller
}

//查询一个用户的好友列表
func (this *FriendsController) Get() {
	id, err := this.GetInt("id")
	if err != nil {
		this.Abort("invalid id")
	}
	limit, err := this.GetInt64("limit")
	if err != nil {
		limit = -1
	}
	offset, err := this.GetInt64("offset")
	if err != nil {
		offset = 0
	}
	method := this.GetString("method")
	// fmt.Println(method)
	var friends []models.Friends
	if method == "" {
		this.Abort("must have a method")
	} else if method == "friends" {
		friends = models.GetFriends(id, limit, offset)
	} else if method == "request" {
		friends = models.GetFriendsRequest(id, limit, offset)
	} else{
		fmt.Println("WTF!")
	}
	bodyJSON := simplejson.New()
	bodyJSON.Set("status", "success")
	tmpMapArr := make([]interface{}, len(friends))
	for i, f := range friends {
		tmpMap := make(map[string]interface{})
		if f.User1Id.Id == id {
			tmpMap["id"] = f.User2Id.Id
			tmpMap["real_name"] = f.User2Id.RealName
			tmpMap["nick_name"] = f.User2Id.NickName
			tmpMap["age"] = f.User2Id.Age
			tmpMap["gender"] = f.User2Id.Gender
			tmpMap["head_picture"] = f.User2Id.HeadPicture
			tmpMap["balance"] = f.User2Id.Balance
			tmpMap["profession"] = f.User2Id.Profession
			tmpMap["grade"] = f.User2Id.Grade
			tmpMap["phone"] = f.User2Id.Phone
			tmpMap["email"] = f.User2Id.Email
		} else {
			tmpMap["id"] = f.User1Id.Id
			tmpMap["real_name"] = f.User1Id.RealName
			tmpMap["nick_name"] = f.User1Id.NickName
			tmpMap["age"] = f.User1Id.Age
			tmpMap["gender"] = f.User1Id.Gender
			tmpMap["head_picture"] = f.User1Id.HeadPicture
			tmpMap["balance"] = f.User1Id.Balance
			tmpMap["profession"] = f.User1Id.Profession
			tmpMap["grade"] = f.User1Id.Grade
			tmpMap["phone"] = f.User1Id.Phone
			tmpMap["email"] = f.User1Id.Email
		}
		tmpMapArr[i] = tmpMap
	}
	bodyJSON.Set("data", tmpMapArr)
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *FriendsController) Post() {
	user1_id, err := this.GetInt("user1_id")
	if err != nil {
		this.Abort("invalid user1 id")
	}
	user2_id, err := this.GetInt("user2_id")
	if err != nil {
		this.Abort("invalid user2 id")
	}
	status := models.AddFriends(user1_id, user2_id)
	bodyJSON := simplejson.New()
	if status == 0 {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "user not exist or database error")
	} else if status == 1 {
		bodyJSON.Set("status", "success")
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "two user have been friends")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}

func (this *FriendsController) Delete() {
	user1_id, err := this.GetInt("user1_id")
	if err != nil {
		this.Abort("invalid user1 id")
	}
	user2_id, err := this.GetInt("user2_id")
	if err != nil {
		this.Abort("invalid user2 id")
	}
	err = models.DeleteFriends(user1_id, user2_id)
	bodyJSON := simplejson.New()
	if err != nil {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", err.Error())
	} else {
		bodyJSON.Set("status", "success")
	}
	body, _ := bodyJSON.Encode()
	this.Ctx.Output.Body(body)
}
