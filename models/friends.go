package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type Friends struct {
	Fid			int		`orm:"column(fid);pk"`
	User1Id  	*User 	`orm:"column(user1_id);rel(fk)"`
	User2Id  	*User 	`orm:"column(user2_id);rel(fk)"`
	Accepted 	int8  	`orm:"column(accepted)"`
}

func (f *Friends) TableName() string {
	return "friends"
}

func init() {
	orm.RegisterModel(new(Friends))
}

// 返回0为失败，返回1为发送好友请求成功或者接受好友请求成功，返回2为已经为好友关系
func AddFriends(user1_id int, user2_id int) int {
	user1, err := GetUserById(user1_id)
	if err != nil {
		return 0
	}
	user2, err := GetUserById(user2_id)
	if err != nil {
		return 0
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Friends))
	cond := orm.NewCondition()
	cond = cond.Or("User1Id", user2).Or("User2Id", user1)
	var values []orm.Params
	count, err := qs.SetCond(cond).Values(&values)
	if err != nil {
		return 0
	}
	if count == 0 {
		friends := Friends{User1Id: user1, User2Id: user2, Accepted: 0}
		_, err := o.Insert(&friends)
		if err == nil {
			return 1
		} else {
			return 0
		}
	} else {
		for _, v := range values {
			if v["Accepted"].(int8) == 0 {
				friends := Friends{User1Id: user2, User2Id: user1, Accepted: 1}
				o.Update(&friends)
				return 1
			} else {
				return 2
			}
		}
	}
	return 0
}

//获得好友列表
func GetFriends(id int, limit int64, offset int64) []Friends {
	var friends []Friends
	user, err := GetUserById(id)
	if err != nil {
		return friends
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Friends))
	cond := orm.NewCondition()
	cond = cond.Or("User1Id", user).Or("User2Id", user)
	qs.SetCond(cond).Filter("Accepted", 1).All(friends)
	return friends
}

//获得请求列表
func GetFriendsRequest(id int, limit int64, offset int64) []Friends {
	var friends []Friends
	user, err := GetUserById(id)
	if err != nil {
		return friends
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(Friends))
	qs.Filter("User2Id", user).Filter("Accepted", 0).All(friends)
	return friends
}

func DeleteFriends(user1_id int, user2_id int) error {
	user1, err := GetUserById(user1_id)
	if err != nil {
		return errors.New("invalid user id")
	}
	user2, err := GetUserById(user2_id)
	if err != nil {
		return errors.New("invalid user id")
	}
	o := orm.NewOrm()
	friends := Friends{User1Id: user1, User2Id: user2, Accepted: 1}
	num1, _ := o.Delete(&friends)
	friends = Friends{User1Id: user2, User2Id: user1, Accepted: 1}
	num2, _ := o.Delete(&friends)
	if num1 == 0 && num2 == 0 {
		return errors.New("these two users are not friends")
	}
	return nil
}
