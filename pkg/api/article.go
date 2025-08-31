package api

import (
	"gin-practice/config"
	"gin-practice/pkg/models"
	"gin-practice/pkg/response"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"content" binding:"required"`
}

type UpgradeArticle struct {
	Title   string `form:"title" json:"title" binding:"-"`
	Content string `form:"content" json:"content" binding:"-"`
}

func CreateArticle(c *gin.Context) {
	var a Article
	userID, exists := c.Get("userID")
	if !exists {
		response.OkWithMessage("not found user", c)
	}
	if err := c.ShouldBind(&a); err != nil {
		response.FailWithError(err, c)
		return
	}
	p := models.Post{
		UserID:  userID.(int64),
		Title:   a.Title,
		Content: a.Content,
	}
	config.SELF_DB.Create(&p)
	response.OkWithMessage("文章创建成功", c)
}

func SelectAllArticle(c *gin.Context) {
	var p []models.Post
	config.SELF_DB.Find(&p)
	response.OkWithData(p, c)
}

func QueryArticle(c *gin.Context) {
	var p models.Post
	postID := c.Param("post_id")
	config.SELF_DB.Where("id = ?", postID).Take(&p)
	response.OkWithData(p, c)
}

func UpdateArticle(c *gin.Context) {
	var a UpgradeArticle
	var p models.Post
	postID := c.Param("post_id")
	userID, exists := c.Get("userID")
	if !exists {
		response.OkWithMessage("not found user", c)
	}
	if err := c.ShouldBind(&a); err != nil {
		response.FailWithError(err, c)
		return
	}
	if err := config.SELF_DB.Where("user_id = ? and id = ?", userID, postID).First(&p).Error; err != nil {
		response.FailWithDetailedMessage("无权更改文章", c)
		return
	}
	if a.Title == "" && a.Content == "" {
		response.FailWithDetailedMessage("请输入修改标题或者内容", c)
	}
	if a.Title != "" {
		p.Title = a.Title
	}
	if a.Content != "" {
		p.Content = a.Content
	}
	config.SELF_DB.Save(&p)
	response.OkWithMessage("文章更新成功", c)
}

func DelArticle(c *gin.Context) {
	var p models.Post
	postID := c.Param("post_id")
	userID, exists := c.Get("userID")
	if !exists {
		response.OkWithMessage("not found user", c)
	}
	if err := config.SELF_DB.Where("user_id = ? and id = ?", userID, postID).First(&p).Error; err != nil {
		response.OkWithMessage("查无此文章", c)
		return
	}
	if err := config.SELF_DB.Where("user_id = ? and id = ?", userID, postID).Delete(&p).Error; err != nil {
		response.FailWithDetailedMessage("无权删除文章", c)
		return
	}
	response.OkWithMessage("文章删除成功", c)
}
