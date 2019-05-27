package models

import (
	// "errors"
	// "fmt"
	// "reflect"
	// "strings"

	"github.com/astaxie/beego/orm"
)

type Msg struct {
	mid         uint    `orm:"column(mid);pk;auto"`
	fromid    	uint  	`orm:"column(fromid);rel(fk)"`
	toid    	uint  	`orm:"column(toid);rel(fk)"`
	create_time string  `orm:"column(create_time);type(datetime);"`
	state       uint  	`orm:"column(state)"`
	content      string  `orm:"column(content);size(140)"`
}

func (t *Msg) TableName() string {
	return "msg"
}

func init() {
	orm.RegisterModel(new(Msg))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func Sendmsg(m *Msg) (fid int64, tid int64, msg string, err error) {
	o := orm.NewOrm()
	
	o.
	return
}
