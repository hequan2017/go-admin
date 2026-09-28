[简体中文](README.md) | [English](README.en.md)

# go-admin

![版本](https://img.shields.io/badge/release-2.0-blue.svg)
![语言](https://img.shields.io/badge/language-goland1.16.5-blue.svg)
![base](https://img.shields.io/badge/base-gin-blue.svg)
![base](https://img.shields.io/badge/base-casbin-blue.svg)

> 一个 Go Web API 后端简单例子，包含 用户、权限、菜单、JWT、RBAC(Casbin) 等。

> **⚠️ 本项目已停止维护，请仅供参考！**

> 交流QQ群： 620176501

## 项目简介

go-admin 是一个使用 Gin + Gorm + Casbin 搭建的 Go Web API 后端示例项目，实现了常见的后台管理核心模块：用户、角色、菜单（权限资源）的管理，以及 JWT 登录认证和基于 Casbin 的 RBAC 接口权限校验。

项目结构清晰、代码量不大，适合作为 Go 后端开发的学习模板：想了解 Gin 中间件如何做 token 校验、Casbin 如何落地 RBAC 权限模型、Gorm 软删除与自动时间戳如何实现，都可以直接阅读源码。权限数据在项目启动时会根据用户-角色-菜单的关联关系自动加载到 Casbin，改动后自动重载，开箱即可跑通。

## ✨ 功能特性

- **JWT 认证**：`/auth` 登录换取 token（HS256，24 小时有效期），API 统一通过 `Authorization: Token xxx` 请求头校验
- **RBAC 权限（Casbin）**：用户关联角色、角色关联菜单（权限资源），启动时自动加载全部策略，用户/角色变动后自动重载；`admin` 用户拥有全部权限不做匹配
- **RESTful API**：用户 / 角色 / 菜单三套资源的完整增删改查，统一 JSON 响应格式（`code` / `msg` / `data`）
- **Swagger 在线文档**：通过 swag 注释生成，启动后访问 `http://127.0.0.1:8000/swagger/index.html`
- **Gorm ORM**：MySQL 驱动，支持软删除（`deleted_on`）、创建/更新时间自动填充、表名前缀、连接池（MaxIdle 10 / MaxOpen 100）
- **参数校验**：基于 beego validation 的入参校验
- **分页**：`page` 参数 + 配置文件中的 `PageSize` 自动计算偏移量
- **日志**：分级日志（DEBUG / INFO / WARN / ERROR / FATAL），按天写入 `logs/` 目录
- **跨域（CORS）中间件**：放开常见跨域请求头，方便前端联调
- **配置化**：应用参数、服务器参数、数据库连接全部收口在 `conf/app.ini`
- **依赖注入**：基于 facebookgo/inject 注入 Casbin Enforcer 与业务对象
- **Docker 部署**：提供 Dockerfile，构建 alpine 镜像并使用 UPX 压缩二进制
- **热编译（开发）**：支持 gowatch，改动代码自动重新编译运行

## 🛠 技术栈

**后端**（版本取自 `go.mod`）：

| 组件 | 版本 | 用途 |
| --- | --- | --- |
| Go | 1.16 | 开发语言 |
| gin-gonic/gin | v1.5.0 | Web 框架 |
| jinzhu/gorm | v1.9.11 | ORM（MySQL） |
| casbin/casbin | v1.9.1 | RBAC 权限模型 |
| dgrijalva/jwt-go | v3.2.0 | JWT 认证 |
| swaggo/gin-swagger / swaggo/swag | v1.2.0 / v1.6.2 | Swagger 文档 |
| go-ini/ini | v1.44.0 | INI 配置解析 |
| astaxie/beego | v1.12.0 | 仅使用其 validation 参数校验 |
| facebookgo/inject | — | 依赖注入 |

**数据库**：MySQL 5.7（`docs/sql/go.sql`，含 `go_user` / `go_role` / `go_menu` 及 `go_user_role` / `go_role_menu` 关联表和初始数据）

**部署**：Dockerfile（alpine + UPX 压缩）、Travis CI、gowatch 热编译

## 🚀 快速开始

### 环境要求

- Go 1.16+
- MySQL 5.7+
- Windows 下开发需要安装 gcc（否则报错 `"gcc" executable file not found in %PATH%`，可参考[这篇文档](https://blog.csdn.net/xia_2017/article/details/105545789)安装）

### 1. 拉取代码并安装依赖

```bash
export GOPROXY=https://goproxy.io
go get go-admin
cd $GOPATH/src/go-admin
```

### 2. 初始化数据库

创建一个名为 `go` 的库，然后导入 `docs/sql/go.sql` 创建表并写入初始数据（含角色：开发部 / 运维部 / 测试部）：

```sql
CREATE DATABASE go;
USE go;
SOURCE docs/sql/go.sql;
```

### 3. 修改配置文件

按需修改 `conf/app.ini`：

```ini
[app]
PageSize = 10          # 分页大小
JwtSecret = 1234567890 # JWT 签名密钥
LogSavePath = logs/    # 日志目录

[server]
RunMode = debug        # debug 或 release
HttpPort = 8000        # 监听端口
ReadTimeout = 60
WriteTimeout = 60

[database]
Type = mysql
User = root
Password = 123456
Host = 127.0.0.1:3306
Name = go              # 库名
TablePrefix = go_      # 表名前缀
```

### 4. 编译运行

```bash
go build main.go
go run main.go
# 启动后监听 :8000
```

### 5. 查看接口文档

启动后访问 Swagger：`http://127.0.0.1:8000/swagger/index.html`

### Docker 部署

仓库根目录提供 `Dockerfile`（两段构建，产出 UPX 压缩后的 alpine 镜像，默认暴露 80 端口）：

```bash
docker build -t go-admin .
docker run -d -p 8000:80 go-admin
```

### API 调用示例

请求和响应均为 JSON 格式。

访问 `/auth` 获取 token（初始账号见 `docs/sql/go.sql`）：

```json
{
    "username": "admin",
    "password": "123456"
}
```

之后访问业务接口（如 `/api/v1/menus?page=2`）需在请求头带上 token：

```
Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

访问 `/api/v1/userInfo` 获取用户信息：前端所需的权限列表放在返回结果的 `password` 字段里（已去重）：

```json
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
}
```

### 接口一览

| 方法 | 路径 | 说明 | 中间件 |
| --- | --- | --- | --- |
| POST | `/auth` | 登录，获取 token | — |
| GET | `/swagger/*any` | Swagger 文档 | — |
| GET | `/api/v1/userInfo` | 当前登录用户信息（含权限列表） | JWT |
| GET / POST / PUT / DELETE | `/api/v1/menus`、`/api/v1/menus/:id` | 菜单（权限资源）管理 | JWT + Casbin |
| GET / POST / PUT / DELETE | `/api/v1/roles`、`/api/v1/roles/:id` | 角色管理 | JWT + Casbin |
| GET / POST / PUT / DELETE | `/api/v1/users`、`/api/v1/users/:id` | 用户管理 | JWT + Casbin |

## 🔑 权限验证说明

项目启动时，会根据 `user`、`role`、`menu` 的关联自动加载 Casbin 策略；如有更改，会删除对应的权限并重新加载：

- 用户 关联 角色（`user_role`）
- 角色 关联 菜单（`role_menu`）

```
权限关系为:
角色(role.name,  menu.path,  menu.method)
用户(user.username,   role.name)

例如:
运维部      /api/v1/users       GET
hequan     运维部

当 hequan GET /api/v1/users 地址的时候，会去检查权限，因为他属于运维部，
同时运维部有对应权限，所以本次请求会通过。

用户 admin 有所有的权限，不进行权限匹配
登录接口 /auth 、 /api/v1/userInfo 不进行验证
```

Casbin 模型定义在 `conf/rbac_model.conf`（keyMatch2 匹配路径、regexMatch 匹配请求方式）。

## 📁 目录结构

```
go-admin/
├── conf/                 # 配置文件（app.ini、rbac_model.conf）
├── docs/                 # 文档
│   ├── sql/go.sql        # 建表与初始数据 SQL
│   └── swagger/          # Swagger 生成的 API 文档
├── logs/                 # 日志
├── middleware/           # 应用中间件
│   ├── inject/           # 依赖注入、Casbin 策略加载
│   ├── jwt/              # JWT 校验
│   └── permission/       # Casbin 权限验证
├── models/               # 数据库模型
├── pkg/                  # 通用工具包（app/e/file/logging/setting/util）
├── routers/              # 路由与 API 处理
├── service/              # 业务逻辑
├── test/                 # 单元测试
├── main.go
└── Dockerfile
```

## 📸 截图

![demo](docs/demo.jpg)

## 🧰 开发

```shell
# 更新 API 文档（修改 swagger 注释后执行）
swag init

# 热编译（开发时使用）
go get github.com/silenceper/gowatch
gowatch

# 后台常驻运行
cd /opt/go-admin
nohup go run main.go >> /tmp/go-http.log 2>&1 &
```

> 优雅重启（Graceful Restart）依赖 fvbock/endless，仅支持 Unix 系统；`main.go` 中保留了对应注释代码，需要时可自行开启。

## 🔗 相关项目

本项目主要参考了：

- [EDDYCJY/go-gin-example](https://github.com/EDDYCJY/go-gin-example) —— 包含更多的例子，上传文件图片等。本项目进行了增改。
- [LyricTian/gin-admin](https://github.com/LyricTian/gin-admin) —— 主要为 gin + casbin 例子。

## 📄 License

[MIT](LICENSE)

## 👤 作者

* 何全（[hequan2017](https://github.com/hequan2017)）
* 交流QQ群： 620176501
