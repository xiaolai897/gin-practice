package models

import (
	"gin-practice/pkg/idgen"
	"time"

	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type Base struct {
	ID        int64     `json:"id" gorm:"primaryKey;autoIncrement:false"`
	CreatedAt time.Time `json:"createAt"`
	UpdatedAt time.Time `json:"updateAt"`
	Createor  int64     `json:"createor"`
	Updater   int64     `json:"updater"`
	DeletedAt soft_delete.DeletedAt
}

func (u *Base) BeforeCreate(tx *gorm.DB) (err error) {
	snowID := idgen.GenerateID()
	u.ID = snowID
	return
}
