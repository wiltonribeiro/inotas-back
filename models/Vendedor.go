package models

type Vendedor struct {
	Cnpj string `json:"cnpj"`
	Nome string `json:"nome"`
	Rua string `json:"rua"`
	Numero string `json:"numero"`
	Cep string `json:"cep"`
	InfoAdicional string `json:"info_adicional"`
	Distrito string `json:"distrito"`
	IdCidade int `json:"id_cidade"`
	Cidade string `json:"cidade"`
	SiglaEstado string `json:"sigla_estado"`
}



