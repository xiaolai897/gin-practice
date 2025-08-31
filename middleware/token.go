package middleware

import (
	"gin-practice/config"
	"gin-practice/pkg/response"
	"gin-practice/utils"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			config.SELF_LOG.Error(err.Error())
			response.FailWithDetailedMessage("认证失败", c)
			c.Abort()
			return
		}
		c.Next()
	}
}
