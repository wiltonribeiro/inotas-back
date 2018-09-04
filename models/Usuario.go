package models

type Usuario struct {
	Email string `json:"email"`
	Nome string `json:"nome"`
	IdCidade string `json:"id_cidade"`
	SiglaEstado string `json:"sigla_estado"`
	Senha string `json:"senha"`
} 
