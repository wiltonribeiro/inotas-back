package models

type Produto struct {
	Id string `json:"id"`
	Code string `json:"code"`
	CetegoriaId int `json:"cetegoria_id"`
	ValorUnidade float64 `json:"valor_unidade"`
	Descricao string `json:"descricao"`
	Unidade string `json:"unidade"`
}

