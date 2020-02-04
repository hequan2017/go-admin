package v1

import (
	"github.com/Anderson-Lu/gofasion/gofasion"
	"github.com/unknwon/com"
	"go-admin/middleware/inject"
	"go-admin/pkg/setting"
	"go-admin/pkg/util"
	"go-admin/service/role_service"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"go-admin/pkg/app"
	"go-admin/pkg/e"
)

// @Summary   获取所有角色
// @Tags role
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles  [GET]
func GetRoles(c *gin.Context) {
	appG := app.Gin{C: c}

	RoleService := Role_service.Role{
		ID:       com.StrTo(c.Query("id")).MustInt(),
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := RoleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_FAIL, nil)
		return
	}

	articles, err := RoleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_S_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

// @Summary   增加角色
// @Tags role
// @Accept json
// @Produce  json
// @Param   body  body   models.Role   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles  [POST]
func AddRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	fsion := gofasion.NewFasion(string(dataByte))
	name := fsion.Get("username").ValueStr()
	menuId := com.StrTo(fsion.Get("menu_id").ValueInt()).MustInt()

	valid := validation.Validation{}
	valid.MaxSize(name, 100, "path").Message("名称最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	RoleService := Role_service.Role{
		Name: name,
		Menu: menuId,
	}

	if id, err := RoleService.Add(); err != nil {

		err = inject.Obj.Common.RoleAPI.LoadPolicy(id)
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

// @Summary   更新角色
// @Tags role
// @Accept json
// @Produce  json
// @Param  id  path  string true "id"
// @Param   body  body   models.Role   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles/:id  [PUT]
func EditRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt()
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	fsion := gofasion.NewFasion(string(dataByte))
	name := fsion.Get("username").ValueStr()
	menuId := com.StrTo(fsion.Get("menu_id").ValueInt()).MustInt()

	valid := validation.Validation{}
	valid.MaxSize(name, 100, "path").Message("名称最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}
	RoleService := Role_service.Role{
		ID:   id,
		Name: name,
		Menu: menuId,
	}
	exists, err := RoleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_FAIL, nil)
		return
	}

	err = RoleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	err = inject.Obj.Common.RoleAPI.LoadPolicy(id)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// @Summary   删除角色
// @Tags role
// @Accept json
// @Produce  json
// @Param  id  path  string true "id"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/roles/:id  [DELETE]
func DeleteRole(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	RoleService := Role_service.Role{ID: id}
	exists, err := RoleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_FAIL, nil)
		return
	}
	role, err := RoleService.Get()
	err = RoleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	inject.Obj.Enforcer.DeleteUser(role.Name)

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
