package routers

import (
	"gin-practice/middleware"
	"gin-practice/pkg/api"

	"github.com/gin-gonic/gin"
)

type CommentRouter struct{}

func (c *CommentRouter) InitCommentRouter(Router *gin.RouterGroup) {
	comment := Router.Group("/comment")
	{
		comment.Use(middleware.JwtAuthMiddleware())
		comment.POST("/create/:post_id", api.CreateRemark)
		comment.GET("/query/:post_id", api.QueryArticleRemark)
	}
}
