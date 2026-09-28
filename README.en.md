[简体中文](README.md) | [English](README.en.md)

# go-admin

![release](https://img.shields.io/badge/release-2.0-blue.svg)
![language](https://img.shields.io/badge/language-goland1.16.5-blue.svg)
![base](https://img.shields.io/badge/base-gin-blue.svg)
![base](https://img.shields.io/badge/base-casbin-blue.svg)

> A simple Go Web API backend example, including users, permissions, menus, JWT, RBAC (Casbin), etc.

> **⚠️ This project is no longer maintained. For reference only!**

> QQ Group: 620176501

## Introduction

go-admin is a Go Web API backend example built with Gin + Gorm + Casbin. It implements the core modules of a typical admin backend: management of users, roles and menus (permission resources), JWT login authentication, and Casbin-based RBAC API authorization.

With a clear structure and a small code base, it works well as a learning template for Go backend development: if you want to see how Gin middleware validates tokens, how Casbin implements an RBAC permission model, or how Gorm handles soft deletes and auto timestamps, just read the source. When the project starts, permission policies are automatically generated from the user-role-menu associations and loaded into Casbin; they are reloaded automatically after changes, so everything works out of the box.

## ✨ Features

- **JWT authentication**: exchange credentials for a token at `/auth` (HS256, valid for 24 hours); APIs are protected via the `Authorization: Token xxx` header
- **RBAC authorization (Casbin)**: users are bound to roles, roles are bound to menus (permission resources); all policies are loaded automatically at startup and reloaded after user/role changes; the `admin` user has full permission without policy matching
- **RESTful API**: full CRUD for users / roles / menus with a unified JSON response format (`code` / `msg` / `data`)
- **Swagger documentation**: generated from swag annotations, available at `http://127.0.0.1:8000/swagger/index.html`
- **Gorm ORM**: MySQL driver with soft delete (`deleted_on`), automatic created/updated timestamps, table name prefix, and connection pool (MaxIdle 10 / MaxOpen 100)
- **Parameter validation**: based on beego validation
- **Pagination**: `page` parameter combined with `PageSize` from the config file
- **Logging**: leveled logs (DEBUG / INFO / WARN / ERROR / FATAL), written daily into `logs/`
- **CORS middleware**: common cross-origin headers are allowed for frontend development
- **Configuration**: app, server and database settings all live in `conf/app.ini`
- **Dependency injection**: Casbin Enforcer and business objects injected via facebookgo/inject
- **Docker deployment**: a Dockerfile that builds an alpine image with a UPX-compressed binary
- **Hot compile (development)**: supports gowatch for automatic rebuild on file changes

## 🛠 Tech Stack

**Backend** (versions from `go.mod`):

| Component | Version | Purpose |
| --- | --- | --- |
| Go | 1.16 | Language |
| gin-gonic/gin | v1.5.0 | Web framework |
| jinzhu/gorm | v1.9.11 | ORM (MySQL) |
| casbin/casbin | v1.9.1 | RBAC authorization |
| dgrijalva/jwt-go | v3.2.0 | JWT authentication |
| swaggo/gin-swagger / swaggo/swag | v1.2.0 / v1.6.2 | Swagger docs |
| go-ini/ini | v1.44.0 | INI config parsing |
| astaxie/beego | v1.12.0 | Only its validation package is used |
| facebookgo/inject | — | Dependency injection |

**Database**: MySQL 5.7 (`docs/sql/go.sql`, containing `go_user` / `go_role` / `go_menu`, the `go_user_role` / `go_role_menu` join tables, and seed data)

**Deployment**: Dockerfile (alpine + UPX compression), Travis CI, gowatch hot compile

## 🚀 Quick Start

### Requirements

- Go 1.16+
- MySQL 5.7+
- On Windows you need gcc for development (otherwise you get `"gcc" executable file not found in %PATH%`; see [this guide](https://blog.csdn.net/xia_2017/article/details/105545789) to install it)

### 1. Clone and install dependencies

```bash
export GOPROXY=https://goproxy.io
go get go-admin
cd $GOPATH/src/go-admin
```

### 2. Initialize the database

Create a database named `go`, then import `docs/sql/go.sql` to create the tables and seed data (roles included: Development / Operations / Testing):

```sql
CREATE DATABASE go;
USE go;
SOURCE docs/sql/go.sql;
```

### 3. Edit the configuration

Adjust `conf/app.ini` as needed:

```ini
[app]
PageSize = 10          # page size
JwtSecret = 1234567890 # JWT signing secret
LogSavePath = logs/    # log directory

[server]
RunMode = debug        # debug or release
HttpPort = 8000        # listen port
ReadTimeout = 60
WriteTimeout = 60

[database]
Type = mysql
User = root
Password = 123456
Host = 127.0.0.1:3306
Name = go              # database name
TablePrefix = go_      # table name prefix
```

### 4. Build and run

```bash
go build main.go
go run main.go
# The server listens on :8000
```

### 5. Browse the API docs

Open Swagger at: `http://127.0.0.1:8000/swagger/index.html`

### Docker deployment

The repository root provides a `Dockerfile` (multi-stage build producing an alpine image with a UPX-compressed binary, exposing port 80 by default):

```bash
docker build -t go-admin .
docker run -d -p 8000:80 go-admin
```

### API usage example

Both requests and responses are JSON.

POST `/auth` to get a token (default account from `docs/sql/go.sql`):

```json
{
    "username": "admin",
    "password": "123456"
}
```

Then call business APIs (e.g. `/api/v1/menus?page=2`) with the token in the request header:

```
Authorization: Token xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

GET `/api/v1/userInfo` returns the current user; the permission list required by the frontend is placed in the `password` field (already de-duplicated):

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

### API overview

| Method | Path | Description | Middleware |
| --- | --- | --- | --- |
| POST | `/auth` | Login and get a token | — |
| GET | `/swagger/*any` | Swagger docs | — |
| GET | `/api/v1/userInfo` | Current user info (with permission list) | JWT |
| GET / POST / PUT / DELETE | `/api/v1/menus`, `/api/v1/menus/:id` | Menu (permission resource) management | JWT + Casbin |
| GET / POST / PUT / DELETE | `/api/v1/roles`, `/api/v1/roles/:id` | Role management | JWT + Casbin |
| GET / POST / PUT / DELETE | `/api/v1/users`, `/api/v1/users/:id` | User management | JWT + Casbin |

## 🔑 How Authorization Works

On startup, the project loads all Casbin policies derived from the `user`, `role` and `menu` associations. If anything changes, the corresponding policies are removed and reloaded:

- User → Role (`user_role`)
- Role → Menu (`role_menu`)

```
The permission relations are:
Role(role.name,  menu.path,  menu.method)
User(user.username,   role.name)

Example:
Operations    /api/v1/users    GET
hequan        Operations

When hequan sends a GET request to /api/v1/users, the permission is checked.
Because hequan belongs to the Operations role and that role has the matching
permission, the request passes.

The admin user has all permissions and skips policy matching.
The login endpoint /auth and /api/v1/userInfo are not checked.
```

The Casbin model is defined in `conf/rbac_model.conf` (keyMatch2 for paths, regexMatch for request methods).

## 📁 Directory Structure

```
go-admin/
├── conf/                 # Configuration files (app.ini, rbac_model.conf)
├── docs/                 # Docs
│   ├── sql/go.sql        # Table creation and seed data SQL
│   └── swagger/          # Swagger-generated API docs
├── logs/                 # Logs
├── middleware/           # Application middleware
│   ├── inject/           # Dependency injection, Casbin policy loading
│   ├── jwt/              # JWT validation
│   └── permission/       # Casbin authorization
├── models/               # Database models
├── pkg/                  # Shared packages (app/e/file/logging/setting/util)
├── routers/              # Routing and API handlers
├── service/              # Business logic
├── test/                 # Unit tests
├── main.go
└── Dockerfile
```

## 📸 Screenshot

![demo](docs/demo.jpg)

## 🧰 Development

```shell
# Regenerate the API docs (after editing swagger annotations)
swag init

# Hot compile (for development)
go get github.com/silenceper/gowatch
gowatch

# Run in the background
cd /opt/go-admin
nohup go run main.go >> /tmp/go-http.log 2>&1 &
```

> Graceful restart relies on fvbock/endless and is Unix-only; the corresponding code is kept commented in `main.go` and can be enabled if needed.

## 🔗 Related Projects

This project mainly referenced:

- [EDDYCJY/go-gin-example](https://github.com/EDDYCJY/go-gin-example) — contains more examples, such as file/image upload. This project extended and modified them.
- [LyricTian/gin-admin](https://github.com/LyricTian/gin-admin) — mainly a gin + casbin example.

## 📄 License

[MIT](LICENSE)

## 👤 Author

* He Quan ([hequan2017](https://github.com/hequan2017))
* QQ Group: 620176501
