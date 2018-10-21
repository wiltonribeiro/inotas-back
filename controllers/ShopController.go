package controllers

import (
	"inotas-back/models"
	"inotas-back/DAOs"
)

type ShopController struct {}

func (controller ShopController) GetShop(token string) ([]models.ShopComplete , models.Error){
	var email string
	authControl := AuthController{}
	email, err  := authControl.CheckAuth(token)
	if err != (models.Error{}) {
		return nil, err
	}

	DAOShop := DAOs.DAOShop{}
	return DAOShop.GetShop(email)
}

func (controller ShopController) UpdateProductsCategories(products []models.Product) models.Error{
	DAOShop := DAOs.DAOShop{}
	return DAOShop.UpdateProductsCategories(products)
}

func (controller ShopController) UpdateShopAlias(shop models.Shop) models.Error{
	DAOShop := DAOs.DAOShop{}
	return DAOShop.UpdateShopAlias(shop)
}
