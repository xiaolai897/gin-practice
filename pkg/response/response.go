package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ERROR   = 500
	SUCCESS = 200
)

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data any, c *gin.Context) {
	Result(SUCCESS, data, "", c)
}

func FailWithError(err error, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, err.Error(), c)
}

func FailWithDetailedMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data any, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
