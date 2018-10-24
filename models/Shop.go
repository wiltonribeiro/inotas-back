package models

type Shop struct {
	NFeKey string `json:"nfe_key"`
	Alias string `json:"alias"`
	UserEmail string `json:"user_email"`
	SellerCnpj string `json:"seller_cnpj"`
	Date string `json:"date"`
	Payment string `json:"payment"`
	TotalCost float64 `json:"total_cost"`
}

type ShopRequest struct {
	NFeKey string `json:"nfe_key"`
	Alias string `json:"alias"`
	Date string `json:"date"`
	Payment string `json:"payment"`
	TotalCost float64 `json:"total_cost"`
	Seller SellerMinimum `json:"seller"`
}



