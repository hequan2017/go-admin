package permission

import (
	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	jwtGet "github.com/hequan2017/go-admin/pkg/util"
	"net/http"
)

func CasbinMiddleware(engine *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		t, _ := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
			return  nil,nil
		})
		if b, err := engine.EnforceSafe(jwtGet.GetIdFromClaims("username",t.Claims), c.Request.URL.Path, c.Request.Method); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "权限 判断错误",
				"msg":  "权限 判断错误",
				"data": "权限 判断错误",
			})
			c.Abort()
			return
		} else if !b {

			c.JSON(http.StatusUnauthorized, gin.H{
				"code": "没有权限",
				"msg":  "没有权限",
				"data": "没有权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
