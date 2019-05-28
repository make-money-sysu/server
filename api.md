## 1./
本地址用来测试服务器是否在正常运行，显示beego默认界面
## 2./login
本地址用来实现用户登录（暂未实现cookie）
只接受POST类型的访问（为了用户密码不会直接暴露在地址中）
POST的body为如下格式的json:
```
{
	"id" : sampleId,
	"password" : "password"
}
```
如果登录成功，返回：
```
{
	"status" : "success"
}
```
如果登录失败，返回：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```
## 3./user/?:id
注： /?:id即为id不一定存在，并且id为变量，如上述地址同时匹配/user和/user/123
此地址用来对用户表进行修改
符合RESTful设计支持GET,POST,PUT和DELETE
GET为获取用户信息，必须提供id(方式如上)：

输入body：raw
```
{
	"id" : 16340132,
	"password" : "dick20",
	"real_name" : "dick",
	"nick_name" : "dick",
	"age" : 18,
	"gender" : "1",
	"head_picture" : " blob not null1",
	"balance" : 100000,
	"profession" : "1",
	"grade" : "3",
	"phone" : "13680473185",
	"email" : "dic0k@qq.com"
}
```

返回为json格式，除了密码外此用户的所有信息都会返回，字段名同数据库中列名(见init.sql)
POST即为注册功能，http包body体中提供json，字段名同数据库即可，如果注册成功，则返回json:
```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```
PUT为修改用户信息功能，使用方法和POST类似，返回同上

## 4./survey/?:id
此地址为对survey进行操作，同样符合RESTful规范用GET,POST,PUT,DELETE来实现增删改查
GET为获取问卷列表，注意这里支持对问卷id,发布者id,问卷名称进行筛选，而且支持设置offset和limit,注意不要用survey/123这种方式来进行id的筛选，所有的筛选字段统一为GET请求地址的参数，比如/survey?id=1&publisher_id=1&name=aaa&name=1&offset=1，如果不设置limit，默认返回所有，不设置offset默认为0
如果查询成功，则返回json:
```
{
	"status" : "success",
	"data" ： survey数组，字段名同数据库
}
```
如果查询失败，则返回json:
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

POST为创建问卷，请求体为json格式，字段名同数据库(id随便填，尽量不要为空，不确定为空可不可以解析)，如果创建成功，则返回json:
```
{
	"status" : "success",
	"id" : survey的id
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

PUT为修改问卷，请求Body类似POST，但是需要提供id（地址中），如果修改成功，则返回json:
```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

DELETE为删除问卷，需要提供id（地址中），如果删除成功，则返回json:
```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

## 5./friends
此地址处理好友申请，同意，删除，获取
支持GET,POST,DELETE，不支持PUT

GET请求获取好友列表或者请求列表，需要在请求参数中添加id，如/friends?id=1,支持offset和limit
同时必须添加参数method,method可以是friends或者request,如果method为friends，则为获取好友列表，如果method为request，则为获取好友请求列表
返回json：
```
{
	"status" : "success"
	"data" : 用户数组，除了password字段不存在其他字段均与数据库相同
}
```

POST为接受好友请求或者添加好友
需要在body体中提供user1_id和user2_id(不是json),body中为类似user1_id=1&user2_id=2即可
user1_id为本用户id，user2为对方id，如果接收请求或者发送添加好友请求成功，则返回json:
```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

DELETE为删除好友，同样需要提供user1_id和user2_id，删除成功，则返回json:
```
{
	"status" : "success"
}
```
失败则返回json：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```
## /package
package为快递相关api，支持GET，PUT和POST请求
GET为获取快递列表，在参数列表里可以提供id(快递的)，owner_id，receiver_id，state，offset，limit，返回格式如下:
```
{
	“status" : "success",
	"data" : 查询出的快递列表，字段名同数据库
}
```

POST为上传快递单，在boby中提供快递的owner_id(所有者)，reward，note即可，其他字段需要存在，但是值无所谓，不然可能会出现不可预知的错误，如果目前登录用户和owner_id不符会LOgin expired，成功则返回：
```
{
	"status" : "success"
}
```
失败则返回：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```

PUT为接单和确认收货功能，在参数列表中必须提供method，method为receive则为接单，confirm则为确认收货，如果method为receive，就必须同时提供receiver_id，并且receiver_id必须是当前登录用户，成功则返回：
```
{
	"status" : "success"
}
```
失败则返回：
```
{
	"status" : "failed",
	"msg" : "报错信息"
}
```
