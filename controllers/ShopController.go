package controllers

import (
	"inotas-back/models"
	"inotas-back/DAOs"
)

type ShopController struct {}

func (controller ShopController) GetShop(token string) ([]models.ShopRequest, models.Error){
	var email string
	authControl := AuthController{}
	email, err  := authControl.CheckAuth(token)
	if err != (models.Error{}) {
		return nil, err
	}

	DAOShop := DAOs.DAOShop{}
	return DAOShop.GetShop(email)
}

func (controller ShopController) UpdateProductsCategories(token string, products []models.Product) models.Error{
	authControl := AuthController{}
	_, err  := authControl.CheckAuth(token)

	if err != (models.Error{}) {
		return err
	}

	DAOShop := DAOs.DAOShop{}
	return DAOShop.UpdateProductsCategories(products)
}

func (controller ShopController) UpdateShopAlias(token string,shop models.Shop) models.Error{
	authControl := AuthController{}
	_, err  := authControl.CheckAuth(token)

	if err != (models.Error{}) {
		return err
	}

	DAOShop := DAOs.DAOShop{}
	return DAOShop.UpdateShopAlias(shop)
}

func (controller ShopController) GetItems(token string,nfe string) ([]models.ItemRequest, models.Error){
	authControl := AuthController{}
	_, err  := authControl.CheckAuth(token)

	if err != (models.Error{}) {
		return nil, err
	}

	DAOShop := DAOs.DAOItem{}
	return DAOShop.GetItems(nfe)
}
