package product

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

type ProductForm struct {
	SellerID    string `json:"sellerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Quantity    int    `json:"quantity"`
}
