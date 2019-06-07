package controllers

import (
	"server/models"
	"strconv"
	// "fmt"
	"time"
	"github.com/bitly/go-simplejson"

	"github.com/astaxie/beego"
)

type SurveyController struct {
	beego.Controller
}

func (this *SurveyController) Get() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	id := this.GetString("id")
	publisherId := this.GetString("publisher_id")
	title := this.GetString("title")
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
	if title != "" {
		queryMap["Title"] = title
	}
	var result []interface{}
	result, err = models.GetAllSurvey(queryMap, []string{}, []string{}, []string{}, offset, limit)
	bodyJSON := simplejson.New()
	if err == nil {
		bodyJSON.Set("status", "success")
		bodyJSON.Set("msg", "just give you")
		tmpMapArr := make([]interface{}, len(result))
		for i, v := range result {
			survey := v.(models.Survey)
			tmpMap := make(map[string]interface{})
			tmpMap["id"] = survey.Id
			tmpMap["publisher_id"] = survey.PublisherId.Id
			tmpMap["title"] = survey.Title
			tmpMap["state"] = survey.State
			tmpMap["checked"] = survey.Checked
			tmpMap["content"] = survey.Content
			tmpMap["create_time"] = survey.CreateTime.String()
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
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")
	// fmt.Println(this.Ctx.Input.Header("cookie"))
	bodyJSON := simplejson.New()
	var survey models.Survey
	inputJSON, err := simplejson.NewJson(this.Ctx.Input.RequestBody)
	if err == nil {
		if nil == this.GetSession("id")  {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			publisher_id := this.GetSession("id").(int)
			survey.PublisherId, _ = models.GetUserById(publisher_id)
			survey.Title = inputJSON.Get("title").MustString()
			survey.Content = inputJSON.Get("content").MustString()
			survey.CreateTime = time.Now()
			if id, err := models.AddSurvey(&survey); err == nil {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("msg", "created")
				bodyJSON.Set("id", id)
			} else {
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "create survey failed")
			}
		}

	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("msg", "invalid json format")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *SurveyController) Put() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	bodyJSON := simplejson.New()
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	// fmt.Println("ID: ")
	// fmt.Println(id)
	if err == nil {
		var survey models.Survey
		inputJSON, _ := simplejson.NewJson(this.Ctx.Input.RequestBody)
		
		if nil == this.GetSession("id")  {
			this.Ctx.Output.SetStatus(401)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("msg", "Login expired")
		}else{
			publisher_id := this.GetSession("id").(int)
			survey.PublisherId, _ = models.GetUserById(publisher_id)
			survey.Title = inputJSON.Get("Title").MustString()
			survey.Content = inputJSON.Get("content").MustString()
			survey.State = inputJSON.Get("state").MustInt()
			survey.Checked = inputJSON.Get("checked").MustInt()
			survey.CreateTime = time.Now()
			survey.Id = id
			err = models.UpdateSurveyById(&survey)
			if err == nil {
				bodyJSON.Set("status", "success")
				bodyJSON.Set("status", "updated")
			} else {
				this.Ctx.Output.SetStatus(403)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("status", "update survey failed")
			}
		}
	}else{
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("status", "formate error")
	}

	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}

func (this *SurveyController) Delete() {
	this.Ctx.Output.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	this.Ctx.Output.Header("Access-Control-Allow-Credentials", "true")

	bodyJSON := simplejson.New()
	id, err := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if err == nil {
		to_delete, err := models.GetSurveyById(id)
		if err == nil {
			if nil == this.GetSession("id") || to_delete.PublisherId.Id != this.GetSession("id").(int)  {
				this.Ctx.Output.SetStatus(401)
				bodyJSON.Set("status", "failed")
				bodyJSON.Set("msg", "Login expired")
			}else{
				err = models.DeleteSurvey(id)
				if err == nil {
					bodyJSON.Set("status", "success")
					bodyJSON.Set("msg", "deleted")
				} else {
					this.Ctx.Output.SetStatus(403)
					bodyJSON.Set("status", "failed")
					bodyJSON.Set("msg", "the id doesn't exist")
				}
			}
		}else{	
			this.Ctx.Output.SetStatus(404)
			bodyJSON.Set("status", "failed")
			bodyJSON.Set("status", "not found")
		}
	} else {
		this.Ctx.Output.SetStatus(400)
		bodyJSON.Set("status", "failed")
		bodyJSON.Set("status", "formate error")
	}
	body, _ := bodyJSON.MarshalJSON()
	this.Ctx.Output.Body(body)
}
