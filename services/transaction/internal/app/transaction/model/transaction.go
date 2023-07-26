package model

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	ID        string         `gorm:"primarykey"`
	UserID    string         `gorm:"index"`
	CreatedAt time.Time      `gorm:"autoCreateTime:false"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:false"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Details []TransactionDetail
}

type TransactionDetail struct {
	ID            string
	TransactionID string `gorm:"index"`
	ProductID     string `gorm:"index"`
	Price         int64
	Quantity      int
}
