package controllers

import (
	"inotas-back/models"
	"inotas-back/enviroment"
	"inotas-back/DAOs"
)

type UserController struct {}

func (controller UserController) Register(user* models.User) (models.Error){

	encrypt := EncryptController{enviroment.SecretKey}
	DAOUser := DAOs.DAOUser{}
	var result string

	result, err := encrypt.Encrypt([]byte(user.Password))
	if err != (models.Error{}) {
		return err
	}
	user.Password = result


	return DAOUser.SaveUser(user)
}

func (controller UserController) ChangePassword(token, newPassword string) (models.Error){
	authControl := AuthController{}
	encryptControl := EncryptController{enviroment.SecretKey}
	DAOUser := DAOs.DAOUser{}

	var email, password string
	email, err  := authControl.CheckAuth(token)
	if err != (models.Error{}) {
		return err
	} else {
		password, err = encryptControl.Encrypt([]byte(newPassword))
		if err != (models.Error{}){
			return err
		}

		return DAOUser.UpdatePassword(password, email)
	}
}

func (controller UserController) GetUser(token string) (user models.User , error models.Error){
	var email string
	authControl := AuthController{}
	DAOUser := DAOs.DAOUser{}

	email, error  = authControl.CheckAuth(token)
	if error != (models.Error{}) {
		return
	}

	return DAOUser.GetUser(email)
}
