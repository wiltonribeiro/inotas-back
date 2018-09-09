package models

type Seller struct {
	Cnpj string `json:"cnpj"`
	Name string `json:"name"`
	Street string `json:"street"`
	Number string `json:"number"`
	PostalCode string `json:"postal_code"`
	OtherInfo string `json:"other_info"`
	District string `json:"district"`
	CityId int `json:"city_id"`
	City string `json:"city"`
	StateInitials string `json:"state_initials"`
}



