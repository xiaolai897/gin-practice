package api

import (
	"gin-practice/config"
	"gin-practice/pkg/models"
	"gin-practice/pkg/response"
	"gin-practice/utils"

	"github.com/gin-gonic/gin"
)

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Email    string `form:"email" json:"email" binding:"required"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func UserRegister(c *gin.Context) {
	var r Register
	if err := c.ShouldBind(&r); err != nil {
		response.FailWithError(err, c)
		return
	}
	r.Password = utils.StringSha256(r.Password)
	u := models.User{
		Username: r.Username,
		Password: r.Password,
		Name:     r.Name,
		Email:    r.Email,
	}
	config.SELF_DB.Create(&u)
	response.OkWithMessage("register success!", c)
}

func UserLogin(c *gin.Context) {
	var l Login
	var u models.User
	if err := c.ShouldBind(&l); err != nil {
		response.FailWithError(err, c)
		return
	}
	config.SELF_DB.Where("username = ?", l.Username).Find(&u)
	crypt := utils.StringSha256(l.Password)
	if u.Password == crypt {
		token, err := utils.GenerateToken(u.ID)
		if err != nil {
			response.FailWithDetailedMessage("暂无数据", c)
			return
		}
		data := map[string]string{
			"token": token,
		}
		response.OkWithData(data, c)
		return
	}
	response.FailWithDetailedMessage("用户名或密码错误", c)
}
