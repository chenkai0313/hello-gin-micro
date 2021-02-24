package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"hello-gin-micro/app"
	"net/http"
)

//中间件
func AuthPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无权访问",
				"data": "",
			})
			c.Abort()
			return
		}
		//判断token是否过期
		redis := app.RedisDB
		val, _ := redis.Get(token).Result()
		if val == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "登陆过期，请重新登陆",
				"data": "",
			})
			c.Abort()
			return
		}
		//解析获取token信息
		_ = json.Unmarshal([]byte(val), &loginInfo)
		urlPath := c.Request.URL.Path
		if !loginInfo.existPermission(c, urlPath) {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无权限",
				"data": "",
			})
			c.Abort()
			return
		}
		//更新token过期时间
		c.Set("admin_id", loginInfo.AdminId)
		loginInfo.UpdateExpired()
	}
}

//是否有权限
func (u LoginInfo) existPermission(c *gin.Context, urlPath string) bool {

	return true
}
