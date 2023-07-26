package model

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          string `gorm:"primarykey"`
	SellerID    string `gorm:"index"`
	Name        string
	Description string
	Price       int64
	Quantity    int
	CreatedAt   time.Time      `gorm:"autoCreateTime:false"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:false"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
