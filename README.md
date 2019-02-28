# Go Admin

一个go api 后端例子,包含JWT,RBAC(Casbin),增删改查, 一键生成 Restful API接口(不依赖orm)。

## 主要说明
* v1.1.1

### 表
* user     
    * username  password   
* role      
    * name 
* menu     
    * name path   method

### 目录结构
* conf：用于存储配置文件
* docs： 文档(SQL和API注释)
* logs： 日志
* middleware：应用中间件
* models：应用数据库模型
* pkg：第三方包
* routers： 路由逻辑处理
* service： 逻辑处理
* test: 单元测试


### 权限验证说明
>  利用的casbin库, 将  user  role  menu 进行自动关联

```
项目启动时,会自动加载权限. 如有更改,会删除对应的权限,重新加载.

用户关联角色  
角色关联菜单  

权限关系为:
角色(role.name,menu.path,menu.method)  
用户(user.username,role.name)

例如:
test      /api/v1/users       GET
hequan     test

当hequan  GET  /api/v1/users 地址的时候，会去检查权限，因为他属于test组，同时组有对应权限，所以本次请求会通过。

用户 admin 有所有的权限,不进行权限匹配

登录接口 /auth  不进行验证
```

### 请求

> 请求和接收 都是 传递 json 格式 数据
```
例如:
访问 /auth    获取token
{
	"username": "admin",
	"password": "123456"
}

访问  /api/v1/users   
请求头设置  Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx

```

## How to run

### Required

- Mysql

### Ready

Create a **go database** and import [SQL](https://github.com/hequan2017/go-admin/blob/master/docs/sql/go.sql)

创建一个库 go,然后导入sql,创建表！

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = go
TablePrefix = go_
```

## Installation
```

yum install go -y 


export GOPROXY=https://goproxy.io
go get github.com/hequan2017/go-admin
cd $GOPATH/src/github.com/hequan2017/go-admin
go build main.go
go run  main.go 

## 热编译,开发时使用

go get github.com/silenceper/gowatch

gowatch   

```

### Run
```
Project information and existing API

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

Listening port is 8000

默认 账户 密码 都为  123456

```

* 不用orm依赖,直接输入表名字就可以 增删改查 

```
http://127.0.0.1:8000/api/restful/go_menu
```

###  API  注释

> http://127.0.0.1:8000/swagger/index.html


## Features
```
- RESTful API
- Gorm
- logging
- Jwt-go
- Swagger
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- 一键生成 Restful API接口
```


## 开发者
* 何全

## 特别感谢
```
本项目主要参考了:
https://github.com/EDDYCJY/go-gin-example  包含更多的例子，上传文件图片等。本项目进行了增改。
https://github.com/LyricTian/gin-admin     主要为 gin+ casbin例子。
https://gitee.com/konyshe/gogo             一行代码搞定RESTFul的轻量web框架。
```


## 其他
```shell
##更新注释
swag init


```