package usecase

import "time"

type Cart struct {
	Details []CartDetail `json:"details"`
}

type CartDetail struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type CartInfo struct {
	TotalPrice int              `json:"totalPrice"`
	Details    []CartInfoDetail `json:"details"`
}

type CartInfoDetail struct {
	SubTotal  int    `json:"subTotal"`
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type Transaction struct {
	ID         string              `json:"id"`
	UserID     string              `json:"userId"`
	Details    []TransactionDetail `json:"details"`
	TotalPrice int64               `json:"totalPrice"`
	CreatedAt  time.Time           `json:"createdAt"`
	UpdatedAt  time.Time           `json:"updatedAt"`
}

type TransactionDetail struct {
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	Price     int64  `json:"price"`
	Quantity  int    `json:"quantity"`
}

type TransactionList struct {
	Count  int           `json:"count"`
	Limit  int           `json:"limit"`
	Offset int           `json:"offset"`
	Data   []Transaction `json:"data"`
}
