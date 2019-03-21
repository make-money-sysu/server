package models

type DoSurvey struct {
	SurveyId    *Survey `orm:"column(survey_id);rel(fk)"`
	RecipientId *User   `orm:"column(recipient_id);rel(fk)"`
	Content     string  `orm:"column(content);size(1000)"`
}
