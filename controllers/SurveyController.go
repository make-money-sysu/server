package controllers

import (
	"server/models"
	"strconv"

	"github.com/bitly/go-simplejson"

	"github.com/astaxie/beego"
)

type SurveyController struct {
	beego.Controller
}

func (this *SurveyController) Get() {
	id := this.GetString("id")
	publisherId := this.GetString("publisher_id")
	name := this.GetString("name")
	limit, err := this.GetInt64("limit")
	if err != nil {
		limit = -1
	}
	offset, err := this.GetInt64("offset")
	if err != nil {
		offset = 0
	}
	queryMap := make(map[string]string)
	if id != "" {
		queryMap["Id"] = id
	}
	if publisherId != "" {
		queryMap["PublisherId"] = publisherId
	}
	if name != "" {
		queryMap["Name"] = name
	}
	var result []interface{}
	result, err = models.GetAllSurvey(queryMap, []string{}, []string{}, []string{}, offset, limit)
	bodyJSON := simplejson.New()
	if err == nil {
		bodyJSON.Set("status", "success")
		tmpMapArr := make([]interface{}, len(result))
		for i, v := range result {
			survey := v.(models.Survey)
			tmpMap := make(map[string]interface{})
			tmpMap["id"] = survey.Id
			tmpMap["publisher_id"] = survey.PublisherId.Id
			tmpMap["name"] = survey.Name
			tmpMap["content"] = survey.Content
			tmpMapArr[i] = tmpMap
		}
		bodyJSON.Set("data", tmpMapArr)
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid query")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *SurveyController) Post() {
	bodyJSON := simplejson.New()
	var survey models.Survey
	inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if err == nil {
		publisher_id := inputJSON.Get("publisher_id").MustInt()
		if publisher_id != this.GetSession("id").(int) {
			this.Abort("Login expired")
		}
		survey.PublisherId, _ = models.GetUserById(publisher_id)
		survey.Name = inputJSON.Get("name").MustString()
		survey.Content = inputJSON.Get("content").MustString()
		if id, err := models.AddSurvey(&survey); err == nil {
			bodyJSON.Set("status", "success")
			bodyJSON.Set("id", id)
		} else {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "create survey failed")
		}
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid json format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *SurveyController) Put() {
	bodyJSON := simplejson.New()
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err == nil {
		var survey models.Survey
		inputJSON, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
		publisher_id := inputJSON.Get("publisher_id").MustInt()
		if publisher_id != this.GetSession("id").(int) {
			this.Abort("Login expired")
		}
		survey.PublisherId, _ = models.GetUserById(inputJSON.Get("publisher_id").MustInt())
		survey.Name = inputJSON.Get("name").MustString()
		survey.Content = inputJSON.Get("content").MustString()
		survey.Id = id
		err = models.UpdateSurveyById(&survey)
	}
	if err == nil {
		bodyJSON.Set("status", "success")
	} else {
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("status", "update survey failed")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *SurveyController) Delete() {
	bodyJSON := simplejson.New()
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err == nil {
		to_delete, _ := models.GetSurveyById(id)
		if to_delete.PublisherId.Id != this.GetSession("id").(int) {
			this.Abort("Login expired")
		}
		err = models.DeleteSurvey(id)
		if err == nil {
			bodyJSON.Set("status", "success")
		} else {
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "the id doesn't exist")
		}
	} else {
		this.Abort("invalid id")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}
