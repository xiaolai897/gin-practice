package models

type User struct {
	Base
	Name          string `json:"name" gorm:"type:varchar(256);index:,unique,where:deleted_at = 0;not null"` // A regular string field
	Email         string `json:"email" gorm:"type:varchar(256);index:,unique,where:deleted_at = 0;not null"`
	Username      string `json:"username" gorm:"type:varchar(256);index:,unique,where:deleted_at = 0;not null"`
	Password      string `gorm:"type:varchar(1024);index:,unique,where:deleted_at = 0;not null"`
	PostNumber    int64  `json:"post_number"`
	CommentNumber int64  `json:"comment_number"`
}
