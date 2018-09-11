package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"strings"
)

type ShopController struct {
	DataBase* database.Connection
}

func (controller ShopController) UpdateProductsCategories(products []models.Product) (error models.Error){
	for _, product := range products {
		if err :=controller.updateProductCategory(product); err != nil{
			return models.ErrorResponse(err, 505)
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

func (controller ShopController) UpdateShopAlias(shop models.Shop) (err error){
	query := "UPDATE shop SET alias = $1 WHERE nfe_key = $2;"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		return err
	}

	upperAlias := strings.ToUpper(shop.Alias)
	_, err = stmt.Exec(strings.TrimSpace(upperAlias), shop.NFeKey)
	return
}
