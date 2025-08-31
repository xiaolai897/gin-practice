package routers

import (
	"gin-practice/middleware"
	"gin-practice/pkg/api"

	"github.com/gin-gonic/gin"
)

type ArticleRouter struct{}

func (a *ArticleRouter) InitArticleRouter(Router *gin.RouterGroup) {
	article := Router.Group("/article")
	{
		article.Use(middleware.JwtAuthMiddleware())
		article.POST("/create", api.CreateArticle)
		article.GET("/list", api.SelectAllArticle)
		article.GET("/query/:post_id", api.QueryArticle)
		article.POST("/update/:post_id", api.UpdateArticle)
		article.POST("/del/:post_id", api.DelArticle)
	}
}
