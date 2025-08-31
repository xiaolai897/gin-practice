package models

import (
	"errors"
	"gin-practice/pkg/idgen"

	"gorm.io/gorm"
)

type Comment struct {
	Base
	UserID  int64  `json:"user_id"`
	PostID  int64  `json:"post_id"`
	Content string `gorm:"not null"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	if c.PostID == 0 || c.UserID == 0 {
		return errors.New("评论失败，没有关联的文章或者用户")
	}
	if err := tx.Model(&Post{}).Select("id").Where("id = ?", c.PostID).Take(&Post{}).Error; err != nil {
		return err
	}
	if err := tx.Model(&User{}).Select("id").Where("id = ?", c.UserID).Take(&User{}).Error; err != nil {
		return err
	}
	c.ID = idgen.GenerateID()
	return
}

func (c *Comment) AfterCreate(tx *gorm.DB) (err error) {
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_number", gorm.Expr("comment_number + 1")).Error; err != nil {
		return err
	}
	if err := tx.Model(&User{}).Where("id = ?", c.UserID).Update("post_number", gorm.Expr("post_number + 1")).Error; err != nil {
		return err
	}
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	if err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_number", gorm.Expr("comment_number - 1")).Error; err != nil {
		return err
	}
	if err := tx.Model(&User{}).Where("id = ?", c.UserID).Update("post_number", gorm.Expr("post_number - 1")).Error; err != nil {
		return err
	}
	return
}
