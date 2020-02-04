package v1

import (
	"github.com/Anderson-Lu/gofasion/gofasion"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"go-admin/pkg/app"
	"go-admin/pkg/e"
	"go-admin/pkg/setting"
	"go-admin/pkg/util"
	"go-admin/service/menu_service"
	"io/ioutil"
	"net/http"
)

// @Summary   获取所有菜单
// @Tags menu
// @Accept json
// @Produce  json
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/menus  [GET]
func GetMenus(c *gin.Context) {
	appG := app.Gin{C: c}

	menuService := menu_service.Menu{
		ID:       com.StrTo(c.Query("id")).MustInt(),
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

// @Summary   增加菜单
// @Tags menu
// @Accept json
// @Produce  json
// @Param   body  body   models.Menu   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/menus  [POST]
func AddMenu(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	fsion := gofasion.NewFasion(string(dataByte))
	name := fsion.Get("name").ValueStr()
	type1 := fsion.Get("type").ValueStr()
	path := fsion.Get("path").ValueStr()
	method := fsion.Get("method").ValueStr()

	valid := validation.Validation{}
	valid.MaxSize(name, 100, "name").Message("最长为100字符")
	valid.MaxSize(type1, 100, "type").Message("最长为100字符")
	valid.MaxSize(path, 100, "path").Message("最长为100字符")
	valid.MaxSize(method, 100, "method").Message("最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}
	menuService := menu_service.Menu{
		Name:   name,
		Type:   type1,
		Path:   path,
		Method: method,
	}

	if err := menuService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

// @Summary   更新菜单
// @Tags menu
// @Accept json
// @Produce  json
// @Param  id  path  string true "id"
// @Param   body  body   models.Menu   true "body"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/menus/:id  [PUT]
func EditMenu(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
	)

	id := com.StrTo(c.Param("id")).MustInt()
	dataByte, _ := ioutil.ReadAll(c.Request.Body)
	fsion := gofasion.NewFasion(string(dataByte))
	name := fsion.Get("name").ValueStr()
	path := fsion.Get("path").ValueStr()
	method := fsion.Get("method").ValueStr()

	valid := validation.Validation{}
	valid.MaxSize(name, 100, "name").Message("最长为100字符")
	valid.MaxSize(path, 100, "path").Message("最长为100字符")
	valid.MaxSize(method, 100, "method").Message("最长为100字符")
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, nil)
		return
	}
	menuService := menu_service.Menu{
		Name:   name,
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

// @Summary   删除菜单
// @Tags menu
// @Accept json
// @Produce  json
// @Param  id  path  string true "id"
// @Success 200 {string} json "{ "code": 200, "data": {}, "msg": "ok" }"
// @Router /api/v1/menus/:id  [DELETE]
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

	menuService := menu_service.Menu{ID: id}
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
