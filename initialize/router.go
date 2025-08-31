package initialize

import (
	"gin-practice/config"
	"gin-practice/middleware"
	"gin-practice/pkg/routers"
	"gin-practice/utils"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	if err := utils.Translator("zh"); err != nil {
		config.SELF_LOG.Error(err.Error())
		return nil
	}

	Router := gin.Default()
	Router.Use(middleware.Recovery())
	PrivateGroup := Router.Group("")
	var authRouter routers.AuthRouter
	var aticleRouter routers.ArticleRouter
	var commentRouter routers.CommentRouter
	authRouter.InitAuthRouter(PrivateGroup)
	aticleRouter.InitArticleRouter(PrivateGroup)
	commentRouter.InitCommentRouter(PrivateGroup)
	return Router
}
