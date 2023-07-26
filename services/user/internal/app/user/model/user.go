package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"primarykey"`
	Username  string `gorm:"uniqueIndex,type:varchar(200)"`
	Password  string
	Email     string `gorm:"uniqueIndex,type:varchar(200)"`
	Name      string
	CreatedAt time.Time      `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:false"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
