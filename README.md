# Go Web Admin 

![构建](https://travis-ci.org/hequan2017/go-admin.svg?branch=master)
![版本](https://img.shields.io/badge/release-1.5.1-blue.svg)
![语言](https://img.shields.io/badge/language-goland1.2-blue.svg)
![base](https://img.shields.io/badge/base-gin-blue.svg)
![base](https://img.shields.io/badge/base-casbin-blue.svg)

> 一个Go Web Api 后端 简单例子,包含 用户、权限、菜单、JWT 、 RBAC(Casbin)等！


## 表
* user     
    * username  password   
* role      
    * name 
* menu     
    * name   path    method


## API  注释

> http://127.0.0.1:8000/swagger/index.html

![demo](docs/demo.jpg)



## 目录结构
* conf：用于存储配置文件
* docs： 文档
    * sql执行命令
    * API注释
* logs： 日志
* middleware：应用中间件
    * inject 初始化对象
    * jwt
    * permission  权限验证
* models：应用数据库模型
* pkg：第三方包
* routers： 路由逻辑处理
* service： 逻辑处理
* test: 单元测试


## 权限验证说明

```
项目启动时,会自动user  role  menu 进行自动关联!  如有更改,会删除对应的权限,重新加载!

用户  关联  角色  
角色  关联  菜单  
```


```
权限关系为:
角色(role.name,  menu.path,  menu.method)  
用户(user.username,   role.name)

例如:
运维部      /api/v1/users       GET
hequan     运维部

当hequan  GET  /api/v1/users 地址的时候，会去检查权限，因为他属于运维部 ，同时 运维部 有对应权限，所以本次请求会通过。

用户 admin 有所有的权限,不进行权限匹配

登录接口 /auth    /api/v1/userInfo 不进行验证
```

## 请求

> 请求和接收 都是 传递 json 格式 数据
```
例如:
访问 /auth    获取token
{
	"username": "admin",
	"password": "123456"
}

访问      /api/v1/menus?page=2    页面是page
请求头设置  Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx


访问   /api/v1/userInfo  获取用户信息
前端所需的权限 放在password字段里面，已经去重了。

"data": {
        "lists": {
            "id": 2,
            "created_on": 1550642309,
            "modified_on": 1550642309,
            "deleted_on": 0,
            "username": "hequan",
            "password": ",system,menu,create_menu,update_menu,delete_menu,user,create_user,update_user,delete_user,role,create_role,update_role,delete_role",
            "role": [
                {
                    "id": 2,
                    "created_on": 0,
                    "modified_on": 0,
                    "deleted_on": 0,
                    "name": "运维部",
                    "menu": null
                }
            ]
        }
    },
```

## 部署

### 支持

- 部署 Mysql

### 库

Create a **go database** and import [SQL](https://go-admin/blob/master/docs/sql/go.sql)

创建一个库 go,然后导入sql,创建表！

### 配置文件

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

### 安装部署
```

yum install go -y 


export GOPROXY=https://goproxy.io
go get go-admin
cd $GOPATH/src/go-admin
go build main.go
go run  main.go 
```


### 热编译(开发时使用)
```bash

go get github.com/silenceper/gowatch

gowatch   
```

## 运行


```
更新角色权限关系 [[hequan 运维部]]
角色权限关系 [[hequan 运维部]]
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /auth                     --> go-admin/routers/api.Auth (4 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/swaggo/gin-swagger.WrapHandler.func1 (4 handlers)
[GIN-debug] GET    /api/v1/menus             --> go-admin/routers/api/v1.GetMenus (6 handlers)
[GIN-debug] GET    /api/v1/menus/:id         --> go-admin/routers/api/v1.GetMenu (6 handlers)
[GIN-debug] POST   /api/v1/menus             --> go-admin/routers/api/v1.AddMenu (6 handlers)
[GIN-debug] PUT    /api/v1/menus/:id         --> go-admin/routers/api/v1.EditMenu (6 handlers)
[GIN-debug] DELETE /api/v1/menus/:id         --> go-admin/routers/api/v1.DeleteMenu (6 handlers)
[GIN-debug] GET    /api/v1/roles             --> go-admin/routers/api/v1.GetRoles (6 handlers)
[GIN-debug] GET    /api/v1/roles/:id         --> go-admin/routers/api/v1.GetRole (6 handlers)
[GIN-debug] POST   /api/v1/roles             --> go-admin/routers/api/v1.AddRole (6 handlers)
[GIN-debug] PUT    /api/v1/roles/:id         --> go-admin/routers/api/v1.EditRole (6 handlers)
[GIN-debug] DELETE /api/v1/roles/:id         --> go-admin/routers/api/v1.DeleteRole (6 handlers)
[GIN-debug] GET    /api/v1/users             --> go-admin/routers/api.GetUsers (6 handlers)
[GIN-debug] GET    /api/v1/users/:id         --> go-admin/routers/api.GetUser (6 handlers)
[GIN-debug] POST   /api/v1/users             --> go-admin/routers/api.AddUser (6 handlers)
[GIN-debug] PUT    /api/v1/users/:id         --> go-admin/routers/api.EditUser (6 handlers)
[GIN-debug] DELETE /api/v1/users/:id         --> go-admin/routers/api.DeleteUser (6 handlers)
2019/05/31 14:48:43 [info] start http server listening :8000

```




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
```

## 前端地址

> https://go-admin-vue

## 开发者
* 何全


## 特别感谢

```
本项目主要参考了:
https://github.com/EDDYCJY/go-gin-example  包含更多的例子，上传文件图片等。本项目进行了增改。
https://github.com/LyricTian/gin-admin     主要为 gin+ casbin例子。
```

## 其他
```shell
## 更新API文档
swag init

## 
cd /opt/go-admin
nohup  go run  main.go   >>  /tmp/go-http.log   2>&1  & 
```