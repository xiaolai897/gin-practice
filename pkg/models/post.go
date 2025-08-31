package models

import (
	"errors"
	"gin-practice/pkg/idgen"

	"gorm.io/gorm"
)

type Post struct {
	Base
	Title         string `json:"title" gorm:"type:varchar(256);not null"`
	Content       string `json:"content" gorm:"not null"`
	CommentNumber int64  `json:"comment_number"`
	UserID        int64  `json:"user_id"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	if p.UserID == 0 {
		return errors.New("文章创建失败，没有关联的用户")
	}
	if err := tx.Model(&User{}).Select("id").Where("id = ?", p.UserID).Take(&User{}).Error; err != nil {
		return err
	}
	p.ID = idgen.GenerateID()
	p.Createor = p.UserID
	p.Updater = p.UserID
	return
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_number", gorm.Expr("post_number + 1")).Error
}

func (p *Post) AfterDelete(tx *gorm.DB) (err error) {
	return tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_number", gorm.Expr("post_number - 1")).Error
}
