package model

import (
	"time"

	"gorm.io/gorm"
)

type Authentication struct {
	Token     string         `gorm:"primarykey"`
	UserID    string         `gorm:"index"`
	CreatedAt time.Time      `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:false"`
	ExpiredAt time.Time      `gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
