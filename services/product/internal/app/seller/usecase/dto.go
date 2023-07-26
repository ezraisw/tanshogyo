package usecase

import "time"

type Seller struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	ShopName  string    `json:"shopName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type SellerForm struct {
	ShopName string `json:"shopName"`
}

type Field struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}
