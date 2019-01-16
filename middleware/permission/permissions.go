package permission

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"github.com/hequan2017/go-admin/models"
	"github.com/hequan2017/go-admin/pkg/logging"
	"net/http"
	"strings"
)

func CasbinMiddleware(engine *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Query("token")
		var user  models.User
		token_header_payload_signature := strings.Split(token, ".")
		token_decode, err := base64.URLEncoding.DecodeString(token_header_payload_signature[1])
		if err != nil {
			logging.Info(err)
		}
		_ = json.Unmarshal([]byte(string(token_decode)), &user)
		fmt.Println(user.Username)
		if b, err := engine.EnforceSafe(user.Username, c.Request.URL.Path, c.Request.Method); err != nil {
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
