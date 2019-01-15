# Go Admin

一个go api 简单例子 ,包含jwt,rbac等 直接拿过来就可以用！

## Installation
```
$ go get github.com/hequan2017/go-admin
```

## How to run

### Required

- Mysql

### Ready

Create a **blog database** and import [SQL](https://github.com/hequan2017/go-admin/blob/master/docs/sql/go.sql)

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

...
```

### Run
```
$ cd $GOPATH/src/go-admin

$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)


Listening port is 8000
Actual pid is 4393


## Features

- RESTful API
- Gorm
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable