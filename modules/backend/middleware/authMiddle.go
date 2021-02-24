package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hello-gin-micro/app"
	"net/http"
	"time"
)

var loginInfo LoginInfo

type LoginInfo struct {
	Token   string `json:"token"`
	AdminId int32  `json:"admin_id"`
}

const Is_Admin_true int32 = 1

//中间件
func AuthMiddleware() gin.HandlerFunc {
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

		_ = json.Unmarshal([]byte(val), &loginInfo)
		c.Set("admin_id", loginInfo.AdminId)
		loginInfo.UpdateExpired()
	}
}

//更新token过期时间
func (u LoginInfo) UpdateExpired() {
	if _, err := app.RedisDB.Expire(u.Token, 24*3600*time.Second).Result(); err != nil {
		fmt.Println("更新token失败")
	}
}
