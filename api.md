## 1./
本地址用来测试服务器是否在正常运行，显示beego默认界面
## 2./login
**目前注册时没给session的，请登录获取**

| 方法 | 参数     | 返回值               |
| ---- | -------- | -------------------- |
| post | id       | status  string 状态  |
|      | password | msg  string 额外信息 |



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
|      |      |      |
| ---- | ---- | ---- |
|      |      |      |
|      |      |      |
|      |      |      |



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

下面的人信息数组一般为以下字段：

id，real_name，nick_name，age，gender，head_picture，balance，

profession，grade，phone，email



|      |                  | 参数                               | 返回值          |
| ---- | ---------------- | ---------------------------------- | --------------- |
| get  | 获取好友请求列表 | method=request   支持offset  limit | "status":string |
|      |                  |                                    | "msg":  string  |
|      |                  |                                    | "data": 人数组  |



此地址处理好友申请，同意，删除，获取
支持GET,POST,DELETE，不支持PUT

GET请求获取好友列表或者请求列表，需要在请求参数中添加id，如/friends?id=1,支持offset和limit
同时必须添加参数method,method可以是friends或者request,如果method为friends，则为获取好友列表，如果method为request，则为
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

快递业务相关

**目前设定全部全体可见**



[可选条件]

包裹信息包含以下字段：

| 字段               | 类型   | 备注                                |
| ------------------ | ------ | ----------------------------------- |
| id                 | int    |                                     |
| owner_id           | int    |                                     |
| owner_real_name    | string |                                     |
| owner_nick_name    | string |                                     |
| owner_Phone        | string |                                     |
| receiver_id        | int    |                                     |
| receiver_real_name | string |                                     |
| receiver_nick_name | string |                                     |
| receiver_Phone     | string |                                     |
| create_time        | string |                                     |
| reward             | int    |                                     |
| state              | int    | 0为刚发布，1为已被接单，2为确认送达 |
| note               | string |                                     |



| 方法 | 作用         | 参数                                                         | 返回值                    |
| ---- | ------------ | ------------------------------------------------------------ | ------------------------- |
| get  | 获得快递列表 | [id (快递id)] [owner_id] [receiver_id] [state] [limit] [offset] | data数组，status 字符串   |
| post | 上传快递单   | reward 整数; note 字符串                                     | status 字符串，msg 字符串 |
| PUT  | 接单         | method:receive ; id (快递id)                                 | status 字符串，msg 字符串 |
| PUT  | 确认收货     | method: confirm ; id (快递id)                                | status 字符串，msg 字符串 |
|      |              |                                                              |                           |



