package controllers

import (
	"inotas-back/models"
	"inotas-back/DAOs"
)

type SellerController struct {}


func (controller SellerController) GetSeller(token, key string) (seller models.Seller, error models.Error){
	authControl := AuthController{}
	DAOSeller := DAOs.DAOSeller{}


	_, error  = authControl.CheckAuth(token)
	if error != (models.Error{}) {
		return
	}

	return DAOSeller.GetSeller(key)

}
