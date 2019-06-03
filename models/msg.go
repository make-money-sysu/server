package models

import (
	// "errors"
	// "fmt"
	// "reflect"
	// "strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Msg struct {
	Mid         	int    		`orm:"column(mid);pk;auto"`
	Fromid    		int  		`orm:"column(fromid);rel(fk)"`
	Toid    		int  		`orm:"column(toid);rel(fk)"`
	Createtime 		time.Time  	`orm:"column(create_time);type(datetime);"`
	State       	uint  		`orm:"column(state)"`
	Content     	string  	`orm:"column(content);size(140)"`
}

func (t *Msg) TableName() string {
	return "msg"
}

func init() {
	orm.RegisterModel(new(Msg))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func SendMessage(m *Msg) (msg string, err error) {
	o := orm.NewOrm()
	_, err = o.Insert(m)
	return "sent", err
}

func WithdrawalMessage(from int,mid int) (msg string, err error) {
	o := orm.NewOrm()
	//  0为系统消息 ，10为未查看，11为已查看，但未知悉，12为已查看，已知悉，13为已撤回
	if num, err := o.Delete(&Msg{Mid: mid, State: 10}); err == nil && num != 0 {
		return "delete successed", err
	}else{
		return "some error happend",err
	}
}

// 返回 这个人未查看，或者被查看但未知悉的信息
func GetMessage(fromid int) (readData []Msg, unreadData []Msg,err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Msg))
	cond := orm.NewCondition()

	user, err := GetUserById(fromid)
	if err != nil {
		return readData, unreadData, err
	}
	cond1 := cond.Or("Fromid", user)
	cond1 = cond1.And("State", 11)

	
	cond2 := cond.Or("Toid", user)
	cond2 = cond2.And("State", 10)

	// cond = cond1.OrCond(cond2)
	qs.SetCond(cond1).All(&readData)

	qs.SetCond(cond2).All(&unreadData)
	return readData, unreadData, err
}

func GetHistory(fromid int, toid int, limit int, offset int) (readData []Msg, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Msg))
	cond := orm.NewCondition()

	user, err := GetUserById(fromid)
	if err != nil {
		return readData, err
	}
	cond = cond.Or("Fromid", user)
	cond = cond.Or("Toid", user)

	qs.SetCond(cond).Limit(limit).Offset(offset).All(&readData)
	return readData, err
}
