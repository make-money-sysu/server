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
	// o := orm.NewOrm()
	
	return "",err
}
