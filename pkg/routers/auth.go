package routers

import (
	"gin-practice/pkg/api"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

func (a *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	auth := Router.Group("")
	{
		auth.POST("/register", api.UserRegister)
		auth.POST("/login", api.UserLogin)
	}
}
