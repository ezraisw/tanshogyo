package model

import (
	"time"

	"gorm.io/gorm"
)

type Seller struct {
	ID        string `gorm:"primarykey"`
	UserID    string `gorm:"uniqueIndex,type:varchar(200)"`
	ShopName  string
	CreatedAt time.Time      `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:false"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
