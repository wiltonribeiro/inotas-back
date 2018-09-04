package controllers

import (
	"inotas-back/models"
	"github.com/nu7hatch/gouuid"
)

type NFeFilter struct {

}

func (filter NFeFilter) getItems(NFe models.NFeRequest) ([]models.Item, []models.Produto){
	var items []models.Item
	var products []models.Produto
	for _, data := range NFe.Items {
		idGenerate,_ := uuid.NewV4()

		product := models.Produto{
			Descricao:data.Description,
			Id: idGenerate.String(),
			Code: data.Code,
			CetegoriaId: 0,
			Unidade: data.Unit,
			ValorUnidade:data.UnitAmount,
		}
		products = append(products, product)

		item := models.Item{
			CompraNFeKey:NFe.Protocol.AccessKey,
			ProdutoId: idGenerate.String(),
			Qntd:data.Quantity,
			ValorTotal:data.TotalAmount,
		}

		items = append(items, item)
	}

	return items, products
}

func (filter NFeFilter) getShop(NFe models.NFeRequest, email string) (shop models.Compra){

	payment := NFe.PaymentType
	if len(NFe.Payment) != 0 {
		payment = NFe.Payment[0].PaymentDetail[0].Method
	}

	shop = models.Compra{
		NFeKey: NFe.Protocol.AccessKey,
		FormaPagamento: payment,
		UsuarioEmail:email,
		DataCompra: NFe.IssuedOn,
		VendedorCnpj: NFe.Issuer.FederalTaxNumber,
		Total:NFe.Totals.Icms.ProductAmount,
	}
	return
}

func (filter NFeFilter) getSeller(NFe models.NFeRequest) (seller models.Vendedor){
	seller = models.Vendedor{
		Cnpj: NFe.Issuer.FederalTaxNumber,
		Nome: NFe.Issuer.Name,
		Cep: NFe.Issuer.Address.PostalCode,
		SiglaEstado:NFe.Issuer.Address.State,
		Distrito:NFe.Issuer.Address.District,
		Rua:NFe.Issuer.Address.Street,
		InfoAdicional:NFe.Issuer.Address.AdditionalInformation,
		Numero:NFe.Issuer.Address.Number,
		Cidade: NFe.Issuer.Address.City.Name,
		IdCidade: 0,
	}

	return
}

func (filter NFeFilter) FilterData(email string, NFe models.NFeRequest) (chan interface{}){
	c := make(chan interface{},4)
	go func() {
		c <- filter.getSeller(NFe)
	}()
	go func() {
		c <- filter.getShop(NFe,email)
	}()
	go func() {
		product, item :=  filter.getItems(NFe)
		c <- product
		c <- item
	}()
	return c
}



