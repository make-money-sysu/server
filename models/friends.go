package models

import (
	"errors"

	"github.com/astaxie/beego/orm"
)

type Friends struct {
	Fid			int		`orm:"column(fid);pk"`
	User1Id  	*User 	`orm:"column(user1_id);rel(fk)"`
	User2Id  	*User 	`orm:"column(user2_id);rel(fk)"`
	Accepted 	bool  	`orm:"column(accepted)"`
}

func (f *Friends) TableName() string {
	return "friends"
}

func init() {
	orm.RegisterModel(new(Friends))
}

// 返回0为失败，返回1为发送好友请求成功或者接受好友请求成功，返回2为已经为好友关系
func AddFriends(user1_id int, user2_id int) int {
	// 找这两个人是否都存在
	user1, err := GetUserById(user1_id)
	if err != nil {
		return 0
	}
	user2, err := GetUserById(user2_id)
	if err != nil {
		return 0
	}
	//查找是否存在反向关系（请求或者已经接受）
	o := orm.NewOrm()
	qs := o.QueryTable(new(Friends))
	cond := orm.NewCondition()
	cond = cond.And("User1Id", user2).And("User2Id", user1)
	var values []orm.Params
	count, err := qs.SetCond(cond).Values(&values)
	if err != nil {
		return 0
	}
	if count == 0 {
		// 没有反向，即不是朋友，也对方没有申请加申请者为好友
		friends := Friends{User1Id: user1, User2Id: user2, Accepted: false}
		_, err := o.Insert(&friends)
		if err == nil {
			return 1
		} else {
			return 0
		}
	} else {
		for _, v := range values {
			if v["Accepted"] == false {
				// 对面有申请加好友，同意~
				friends := Friends{User1Id: user2, User2Id: user1, Accepted: true}
				o.Update(&friends) // TODO:: 这里是否应该加入容错？

				// 加入本人加对面的好友
				friends1 := Friends{User1Id: user1, User2Id: user2, Accepted: true}
				_, err := o.Insert(&friends1)
				if err == nil {
					return 1
				} else {
					return 0
				}

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
	cond = cond.Or("User1Id", user)  //.Or("User2Id", user)
	qs.SetCond(cond).Filter("Accepted", true).All(friends)
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
	qs.Filter("User2Id", user).Filter("Accepted", false).All(friends)
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
	friends := Friends{User1Id: user1, User2Id: user2, Accepted: true} //TODO::待测试 直觉上不用写这么详细。。
	num1, _ := o.Delete(&friends)
	friends = Friends{User1Id: user2, User2Id: user1, Accepted: true}
	num2, _ := o.Delete(&friends)
	if num1 == 0 && num2 == 0 {
		return errors.New("these two users are not friends")
	}
	return nil
}
