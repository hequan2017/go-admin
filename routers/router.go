package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/hequan2017/go-admin/middleware/inject"
	"github.com/hequan2017/go-admin/middleware/permission"
	"github.com/hequan2017/go-admin/middleware/jwt"
	"github.com/hequan2017/go-admin/pkg/setting"
	"github.com/hequan2017/go-admin/routers/api"
	"github.com/hequan2017/go-admin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	obj := inject.Init()

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)

	err := loadCasbinPolicyData(obj)
	if err != nil {
		panic("加载casbin策略数据发生错误：" + err.Error())
	}

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(permission.CasbinMiddleware(obj.Enforcer))

	apiv1.Use(jwt.JWT())
	{

		apiv1.GET("/menus", v1.GetMenus)
		apiv1.GET("/menus/:id", v1.GetMenu)
		apiv1.POST("/menus", v1.AddMenu)
		apiv1.PUT("/menus/:id", v1.EditMenu)
		apiv1.DELETE("/menus/:id", v1.DeleteMenu)

		apiv1.GET("/roles", v1.GetRoles)
		apiv1.GET("/roles/:id", v1.GetRole)
		apiv1.POST("/roles", v1.AddRole)
		apiv1.PUT("/roles/:id", v1.EditRole)
		apiv1.DELETE("/roles/:id", v1.DeleteRole)

		apiv1.GET("/users", api.GetUsers)
		apiv1.GET("/users/:id", api.GetUser)
		apiv1.POST("/users", api.AddUser)
		apiv1.PUT("/users/:id", api.EditUser)
		apiv1.DELETE("/users/:id", api.DeleteUser)
	}

	return r
}

// 加载casbin策略数据，包括角色权限数据、用户角色数据
func loadCasbinPolicyData(obj *inject.Object) error {
	c := obj.Common

	err := c.RoleAPI.LoadAllPolicy()
	if err != nil {
		return err
	}
	err = c.UserAPI.LoadAllPolicy()
	if err != nil {
		return err
	}

	return nil
}
