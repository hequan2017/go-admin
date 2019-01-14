package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hequan2017/go-admin/pkg/app"
	"github.com/hequan2017/go-admin/pkg/e"
	"github.com/hequan2017/go-admin/pkg/setting"
	"github.com/hequan2017/go-admin/pkg/util"
	"github.com/hequan2017/go-admin/service/menu_service"
	"net/http"
)

func GetMenu(c *gin.Context) {
	appG := app.Gin{C: c}
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if !valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	menuService := menu_service.Menus{ID: id}
	exists, err := menuService.ExistByID()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST, nil)
		return
	}

	article, err := menuService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

func GetMenus(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	menuService := menu_service.Menus{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := menuService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_FAIL, nil)
		return
	}

	articles, err := menuService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_S_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = articles
	data["total"] = total

	appG.Response(http.StatusOK, e.SUCCESS, data)
}

func AddMenu(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	path := c.Query("path")
	method := c.Query("method")

	valid := validation.Validation{}
	valid.MaxSize(path, 100, "path").Message("名称最长为100字符")
	valid.MaxSize(path, 100, "method").Message("名称最长为100字符")

	if !valid.HasErrors() {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	menuService := menu_service.Menus{
		Path:   path,
		Method: method,
	}
	if err := menuService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func EditMenu(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()
	path := c.Query("path")
	method := c.Query("method")

	valid := validation.Validation{}
	valid.MaxSize(path, 100, "path").Message("名称最长为100字符")
	valid.MaxSize(path, 100, "method").Message("名称最长为100字符")
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if !valid.HasErrors() {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}
	menuService := menu_service.Menus{
		Path:   path,
		Method: method,
	}
	exists, err := menuService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_FAIL, nil)
		return
	}

	err = menuService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteMenu(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	menuService := menu_service.Menus{ID: id}
	exists, err := menuService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_FAIL, nil)
		return
	}

	err = menuService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
