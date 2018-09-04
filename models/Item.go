package models

type Item struct {
	Id int `json:"id"`
	CompraNFeKey string `json:"compra_nfe_key"`
	ProdutoId string `json:"produto_id"`
	ValorTotal float64 `json:"valor_total"`
	Qntd float64 `json:"qntd"`
}



