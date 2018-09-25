package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"strings"
)

type ShopController struct {
	DataBase* database.Connection
}

func (controller ShopController) GetShop(token string) (shops []models.ShopComplete ,error models.Error){
	var email string
	authControl := AuthController{}
	email, error  = authControl.CheckAuth(token)
	if error != (models.Error{}) {
		return
	}

	query := "SELECT shop.*, seller.* FROM shop INNER JOIN seller ON shop.seller_cnpj = seller.cnpj WHERE shop.user_email = $1;"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		error = models.ErrorResponse(err,500)
		return
	}

	rows, err := stmt.Query(email)

	if err != nil {
		error = models.ErrorResponse(err,500)
		return
	}
	for rows.Next(){
		var shop models.ShopComplete
		rows.Scan(&shop.NFeKey, &shop.TotalCost, &shop.Payment, &shop.Date, &shop.UserEmail, &shop.SellerCnpj, &shop.Alias,
			&shop.Seller.Cnpj, &shop.Seller.Name, &shop.Seller.Street, &shop.Seller.Number, &shop.Seller.PostalCode, &shop.Seller.OtherInfo, &shop.Seller.District, &shop.Seller.CityId, &shop.Seller.StateInitials)
		shops = append(shops,shop)
	}
	return
}

func (controller ShopController) UpdateProductsCategories(products []models.Product) (error models.Error){
	for _, product := range products {
		if err :=controller.updateProductCategory(product); err != nil{
			return models.ErrorResponse(err, 500)
		}
	}
	return
}

func (controller ShopController) updateProductCategory(product models.Product) (err error){
	query := "UPDATE product SET category_id = $1 WHERE id = $2;"
	stmt, err := controller.DataBase.GetDB().Prepare(query)

	if err != nil {
		return err
	}

	_,err = stmt.Exec(product.CategoryId, product.Id)
	return err
}

func (controller ShopController) UpdateShopAlias(shop models.Shop) (error models.Error){
	query := "UPDATE shop SET alias = $1 WHERE nfe_key = $2;"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		return models.ErrorResponse(err, 500)
	}

	upperAlias := strings.ToUpper(shop.Alias)
	_, err = stmt.Exec(strings.TrimSpace(upperAlias), shop.NFeKey)
	return
}
