package models

type Shop struct {
	NFeKey string `json:"nfe_key"`
	UserEmail string `json:"user_email"`
	SellerCnpj string `json:"seller_cnpj"`
	Date string `json:"date"`
	Payment string `json:"payment"`
	TotalCost float64 `json:"total_cost"`
}



