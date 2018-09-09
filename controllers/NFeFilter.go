package controllers

import (
	"inotas-back/models"
	"github.com/nu7hatch/gouuid"
)

type NFeFilter struct {

}

func (filter NFeFilter) getItems(NFe models.NFeRequest) ([]models.Item, []models.Product){
	var items []models.Item
	var products []models.Product
	for _, data := range NFe.Items {
		idGenerate,_ := uuid.NewV4()

		product := models.Product{
			Description:data.Description,
			Id: idGenerate.String(),
			Code: data.Code,
			CategoryId: 0,
			Unity: data.Unit,
			UnityCost:data.UnitAmount,
		}
		products = append(products, product)

		item := models.Item{
			ShopKey:NFe.Protocol.AccessKey,
			ProductId: idGenerate.String(),
			Quantity:data.Quantity,
			TotalCost:data.TotalAmount,
		}

		items = append(items, item)
	}

	return items, products
}

func (filter NFeFilter) getShop(NFe models.NFeRequest, email string) (shop models.Shop){

	payment := NFe.PaymentType
	if len(NFe.Payment) != 0 {
		payment = NFe.Payment[0].PaymentDetail[0].Method
	}

	shop = models.Shop{
		NFeKey: NFe.Protocol.AccessKey,
		Payment: payment,
		UserEmail:email,
		Date: NFe.IssuedOn,
		SellerCnpj: NFe.Issuer.FederalTaxNumber,
		TotalCost:NFe.Totals.Icms.ProductAmount,
	}
	return
}

func (filter NFeFilter) getSeller(NFe models.NFeRequest) (seller models.Seller){
	seller = models.Seller{
		Cnpj: NFe.Issuer.FederalTaxNumber,
		Name: NFe.Issuer.Name,
		PostalCode: NFe.Issuer.Address.PostalCode,
		StateInitials:NFe.Issuer.Address.State,
		District:NFe.Issuer.Address.District,
		Street:NFe.Issuer.Address.Street,
		OtherInfo:NFe.Issuer.Address.AdditionalInformation,
		Number:NFe.Issuer.Address.Number,
		City: NFe.Issuer.Address.City.Name,
		CityId: 0,
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



