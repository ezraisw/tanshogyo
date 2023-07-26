package usecase

import "time"

type Product struct {
	ID          string    `json:"id"`
	SellerID    string    `json:"sellerId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int64     `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AuthedProductForm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int    `json:"quantity"`
}

type ProductForm struct {
	SellerID    string `json:"sellerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int    `json:"quantity"`
}

type ProductList struct {
	Count  int       `json:"count"`
	Limit  int       `json:"limit"`
	Offset int       `json:"offset"`
	Data   []Product `json:"data"`
}
