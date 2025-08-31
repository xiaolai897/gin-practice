package api

import (
	"gin-practice/config"
	"gin-practice/pkg/models"
	"gin-practice/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Remark struct {
	Content string `form:"content" json:"content" binding:"required"`
}

func CreateRemark(c *gin.Context) {
	var r Remark
	var p models.Post
	postIDStr := c.Param("post_id")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		response.FailWithDetailedMessage("ID格式错误", c)
		return
	}
	userID, exists := c.Get("userID")
	if !exists {
		response.OkWithMessage("not found user", c)
	}
	if err := c.ShouldBind(&r); err != nil {
		response.FailWithError(err, c)
		return
	}
	if err := config.SELF_DB.Where("id = ?", postID).First(&p).Error; err != nil {
		response.FailWithDetailedMessage("查无文章，请正确进行评论", c)
		return
	}

	comment := models.Comment{
		Content: r.Content,
		UserID:  userID.(int64),
		PostID:  postID,
		Base: models.Base{
			Createor: userID.(int64),
			Updater:  userID.(int64),
		},
	}
	config.SELF_DB.Create(&comment)
	response.OkWithMessage("发表评论成功", c)
}

func QueryArticleRemark(c *gin.Context) {
	var comment []models.Comment
	postIDStr := c.Param("post_id")
	config.SELF_DB.Where("post_id = ?", postIDStr).Find(&comment)
	response.OkWithData(comment, c)
}
