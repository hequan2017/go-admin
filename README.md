# Go Admin

一个go api 后端例子,包含JWT,RBAC等!(逐步完善中...)

## 主要说明

### 表
* user      username  password   
* role      name 
* menu      path   method


### 权限验证说明
>  利用的casbin库, 将  user  role  menu 进行自动关联

```
项目启动时,会自动加载权限. 如有更改,会删除对应的权限,重新加载.

用户 关联 角色   角色关联菜单 , 只需要添加对应的数据,就会自动判断

权限关系为:
角色(role.name,menu.path,menu.method)  
用户(user.username,role.name)

例如:
test  /api/v1/users  GET
hequan  test

当hequan  GET  /api/v1/users 地址的时候，会去检查权限，因为他属于test组，同时组有对应权限，所以本次请求会通过。

因为现在是测试 在模拟请求的时候 需要加上   /api/v1/users?token=xxxxxxxxxxx  

用户 admin 有所有的权限
```
## demo
```html
先获取token
http://47.104.140.38:8002/auth?username=hequan&password=123456
然后把token 放到下面
http://47.104.140.38:8002/api/v1/users?token=xxxxxxxxxxxxxxx
```

## How to run

### 特别注意

>  github.com\hequan2017\go-admin\middleware\inject\inject.go   22 行, 现在是windows开发,如果想在linux运行，需要修改一下为linux格式    path := dir + "\\\conf\\\rbac_model.conf" 

### Required

- Mysql

### Ready

Create a **go database** and import [SQL](https://github.com/hequan2017/go-admin/docs/sql/go.sql)

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
git clone  http://github.com/hequan2017/go-admin
go build main.go
go run  main.go
```

### Run
```
Project information and existing API

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

Listening port is 8000
```

## Features
```
- RESTful API
- Gorm
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
```

## 开发者
* 何全



## 特别感谢
```
本项目主要参考了:
https://github.com/EDDYCJY/go-gin-example  包含更多的例子，上传文件图片等。本项目进行了增改。
https://github.com/LyricTian/gin-admin     主要为 gin+ casbin例子。
```