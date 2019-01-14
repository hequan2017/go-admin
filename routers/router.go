package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/hequan2017/go-admin/pkg/setting"
	"github.com/hequan2017/go-admin/routers/api"
	"github.com/hequan2017/go-admin/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	
	
	r.GET("/auth", api.GetAuth)


	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		
		apiv1.GET("/menus", v1.GetMenus)

		apiv1.GET("/menus/:id", v1.GetMenu)

		apiv1.POST("/menus", v1.AddMenu)

		apiv1.PUT("/menus/:id", v1.EditMenu)

		apiv1.DELETE("/menus/:id", v1.DeleteMenu)

	}

	return r
}
