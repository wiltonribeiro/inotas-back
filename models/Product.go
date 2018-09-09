package models

type Product struct {
	Id string `json:"id"`
	Code string `json:"code"`
	CategoryId int `json:"category_id"`
	UnityCost float64 `json:"unity_cost"`
	Description string `json:"description"`
	Unity string `json:"unity"`
}

