### 关于运行

当然先安装go

然后增加包beego（warning：切莫反着读）。

> go get github.com/astaxie/beego
>
> 反正缺的都go get 就好了。
>
> 然后会说找不到routers，虽然路径上显然是可以找到的。但是他就是不行。因为这就是个[bug](<https://github.com/astaxie/beego/issues/810>)把他放在go的安装目录里面的src之中就好了。
>
> 然后就是新的报错
>
> > non-standard import "github.com/astaxie/beego" in standard package
> >
> > 贼刺激。
> >
> > 忽然发现他500+个问题未解决，感觉头皮发麻。
>
> 原来是位置不对。应该是要放在项目目录之下，GOPATH下面而不是GOROOT下。再挪一下就好了。继续go get然后就能go build main.go出来一个exe，运行他就能够打开了。
>
> （当然现在的项目还有bug而没成功，不影响我继续~



## 关于friends功能的修改

目前的功能需求是实现 好友关系的添加和删除功能。本来的设计是使用两个user的id联合作为联合主键来建立表格然后用一个bool值来标记他们之间的关系（是否已经接受，没有记录就表示还没有发出申请）。据刘恒伟描述是，orm不支持联合主键，那样的话直接不用联合主键就好了。（另外的话，现在是一个带方向的映射关系，接受之后应该要注意一下把另一个方向的关掉，之类的操作。。直觉上，应该是会有一个一对一的意思的表格而不用user1，user2这样来搞，但是没找到，先这样用着，增删的时候留一下就好）。增加一个自增字段fid，表示两者之间的关系映射id。



TODO::

1. 修改model.friends里面的查改内容
2. 留意增删的时候反方向的处理
3. 跑起来~