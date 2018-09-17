package models

type User struct {
	Email string `json:"email"`
	Name string `json:"name"`
	CityId int `json:"city_id"`
	StateInitials string `json:"state_initials"`
	Password string `json:"password"`
} 
