package api

import "C"
import (
	"github.com/astaxie/beego/validation"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-admin/middleware/inject"
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	"go-admin/pkg/setting"
	"go-admin/pkg/util"
	jwtGet "go-admin/pkg/util"
	Role_service "go-admin/service/role_service"
	"go-admin/service/user_service"
	"net/http"
	"strings"
)

type auth struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role_id"`
}

// @Summary   获取登录token 信息
// @Tags auth
// @Accept json
// @Produce  json
// @Param   body  body   models.AuthSwag   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {string} json
// @Router /auth  [POST]
func Auth(c *gin.Context) {

	appG := app.Gin{C: c}
	var reqInfo auth
	err := c.BindJSON(&reqInfo)
	//dataByte, _ := ioutil.ReadAll(c.Request.Body)
	//fsion := gofasion.NewFasion(string(dataByte))
	//fmt.Println(fsion.Get("username").ValueStr())

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	valid := validation.Validation{}
	valid.MaxSize(reqInfo.Username, 100, "username").Message("最长为100字符")
	valid.MaxSize(reqInfo.Password, 100, "password").Message("最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
		return
	}

	authService := user_service.User{Username: reqInfo.Username, Password: reqInfo.Password}
	isExist, err := authService.Check()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	user, err := authService.Get()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	token, err := util.GenerateToken(user.ID, reqInfo.Username, reqInfo.Password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary   获取登录用户信息
// @Tags  users
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {"lists":""}, "msg": "ok" }"
// @Failure 400 {string} json
// @Router /api/v1/userInfo  [GET]
func GetUserInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	Authorization := c.GetHeader("Authorization")
	token := strings.Split(Authorization, " ")
	t, err := jwt.Parse(token[1], func(*jwt.Token) (interface{}, error) {
		return jwtGet.JwtSecret, nil
	})

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH, nil)
		return
	}
	u := jwtGet.GetIdFromClaims("username", t.Claims)
	userService := user_service.User{
		Username: u,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	user, err := userService.Get()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_S_FAIL, nil)
		return
	}

	menus := make([]string, 1)
	if u == "admin" {
		menus = append(menus, "admin")
	}

	for _, v := range user.Role {
		RoleService := Role_service.Role{ID: v.ID}
		exists, err := RoleService.ExistByID()

		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
			return
		}
		if !exists {
			appG.Response(http.StatusOK, e.ERROR_NOT_EXIST, nil)
			return
		}

		r, err := RoleService.Get()
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
			return
		}
		for _, v2 := range r.Menu {
			menus = append(menus, v2.Name)
		}
		menus = util.RemoveRepByMap(menus)

	}

	user.Password = strings.Join(menus, ",")

	data := make(map[string]interface{})
	data["lists"] = user

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary   获取所有用户
// @Tags  users
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {string} json
// @Router /api/v1/users  [GET]
func GetUsers(c *gin.Context) {
	appG := app.Gin{C: c}

	userService := user_service.User{
		ID:       com.StrTo(c.Query("id")).MustInt(),
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := userService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_FAIL, nil)
		return
	}

	user, err := userService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_S_FAIL, nil)
		return
	}
	for _, v := range user {
		v.Password = ""
	}

	data := make(map[string]interface{})
	data["lists"] = user
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary   增加用户
// @Tags  users
// @Accept json
// @Produce  json
// @Param   body  body   models.User   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {string} json
// @Router /api/v1/users  [POST]
func AddUser(c *gin.Context) {

	appG := app.Gin{C: c}
	var reqInfo auth
	err := c.BindJSON(&reqInfo)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	valid := validation.Validation{}
	valid.MaxSize(reqInfo.Username, 100, "username").Message("最长为100字符")
	valid.MaxSize(reqInfo.Password, 100, "password").Message("最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
		return
	}

	userService := user_service.User{
		Username: reqInfo.Username,
		Password: reqInfo.Password,
		Role:     reqInfo.Role,
	}
	if id, err := userService.Add(); err == nil {
		err = inject.Obj.Common.UserAPI.LoadPolicy(id)
		if err != nil {
			appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
			return
		}
		appG.Response(http.StatusOK, e.SUCCESS, nil)
	} else {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}
}

// @Summary   更新用户
// @Tags  users
// @Accept json
// @Produce  json
// @Param   body  body   models.User   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Failure 400 {string} json
// @Router /api/v1/users/:id  [PUT]
func EditUser(c *gin.Context) {

	appG := app.Gin{C: c}
	var reqInfo auth
	err := c.BindJSON(&reqInfo)

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(reqInfo.Username, 100, "username").Message("最长为100字符")
	valid.MaxSize(reqInfo.Password, 100, "password").Message("最长为100字符")

	if err != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
		return
	}

	userService := user_service.User{
		ID:       id,
		Username: reqInfo.Username,
		Password: reqInfo.Password,
		Role:     reqInfo.Role,
	}
	exists, err := userService.ExistByID()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_FAIL, nil)
		return
	}

	err = userService.Edit()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	err = inject.Obj.Common.UserAPI.LoadPolicy(id)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary   删除用户
// @Tags  users
// @Accept json
// @Produce  json
// @Param  id  path  int true "id"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/users/:id  [DELETE]
func DeleteUser(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	userService := user_service.User{ID: id}
	exists, err := userService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_FAIL, nil)
		return
	}
	user, err := userService.Get()

	err = userService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	inject.Obj.Enforcer.DeleteUser(user.Username)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
