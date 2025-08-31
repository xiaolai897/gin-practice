package middleware

import (
	"gin-practice/config"
	"gin-practice/pkg/response"
	"gin-practice/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Recovery() gin.HandlerFunc {
	return gin.CustomRecovery(func(c *gin.Context, recovered any) {
		if errs, ok := recovered.(validator.ValidationErrors); ok {
			response.FailWithDetailed(errs.Translate(utils.Trans), "校验失败", c)
			return
		}
		if err, ok := recovered.(error); ok {
			config.SELF_LOG.Error(err.Error())
			response.FailWithError(err, c)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	})
}
