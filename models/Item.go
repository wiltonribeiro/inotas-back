package models

type Item struct {
	Id int `json:"id"`
	ShopKey string `json:"shop_key"`
	ProductId string `json:"product_id"`
	TotalCost float64 `json:"total_cost"`
	Quantity float64 `json:"quantity"`
}

type ItemRequest struct {
	Id int `json:"id"`
	TotalCost float64 `json:"total_cost"`
	Quantity float64 `json:"quantity"`
	Product Product `json:"product"`
}




