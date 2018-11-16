package controllers

import (
	"inotas-back/models"
	"inotas-back/DAOs"
)

type ShopController struct {}

func (controller ShopController) GetShopList(token string) ([]models.ShopRequest, models.Error){
	var email string
	authControl := AuthController{}
	email, err  := authControl.CheckAuth(token)
	if err != (models.Error{}) {
		return nil, err
	}

	DAOShop := DAOs.DAOShop{}
	return DAOShop.GetShopList(email)
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

	DAOItem := DAOs.DAOItem{}
	return DAOItem.GetItems(nfe)
}

func (controller ShopController) DeleteShop(token string,nfe string) (models.Error){
	authControl := AuthController{}
	email, err  := authControl.CheckAuth(token)

	if err != (models.Error{}) {
		return err
	}

	DAOShop := DAOs.DAOShop{}
	return DAOShop.DeleteShop(nfe, email)
}
