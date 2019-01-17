package v1

import (
	"github.com/Unknwon/com"
	"github.com/hequan2017/go-admin/pkg/setting"
	"github.com/hequan2017/go-admin/service/role_service"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/hequan2017/go-admin/pkg/app"
	"github.com/hequan2017/go-admin/pkg/e"
	"github.com/hequan2017/go-admin/pkg/util"
)

type auth struct {
	Rolename string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func GetRole(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	RoleService := Role_service.Role{ID: id}
	exists, err := RoleService.ExistByID()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST, nil)
		return
	}

	article, err := RoleService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

func AddRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	name := c.Query("name")
	menu_id := com.StrTo(c.Query("menu_id")).MustInt()

	valid := validation.Validation{}
	valid.MaxSize(name, 100, "path").Message("名称最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	RoleService := Role_service.Role{
		Name: name,
		Menu: menu_id,
	}

	if err := RoleService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func GetRoles(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	valid := validation.Validation{}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	RoleService := Role_service.Role{
		Name:  name,
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

func EditRole(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	menu_id := com.StrTo(c.Query("menu_id")).MustInt()

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
		Menu: menu_id,
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

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

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

	err = RoleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
