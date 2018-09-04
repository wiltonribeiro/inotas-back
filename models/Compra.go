package models

type Compra struct {
	NFeKey string `json:"nfe_key"`
	UsuarioEmail string `json:"usuario_email"`
	VendedorCnpj string `json:"vendedor_cnpj"`
	DataCompra string `json:"data_compra"`
	FormaPagamento string `json:"forma_pagamento"`
	Total float64 `json:"total"`
}



