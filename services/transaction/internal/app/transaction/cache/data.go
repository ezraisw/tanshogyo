package cache

type Cart struct {
	Details []CartDetail `json:"details"`
}

type CartDetail struct {
	ProductID string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
