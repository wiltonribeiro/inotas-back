package controllers

import (
	"inotas-back/models"
	"inotas-back/enviroment"
	"errors"
	"inotas-back/DAOs"
)

type LoginController struct {}

type Exists struct {
	Password string `json:"password"`
}

func (controller LoginController) Login(email, password string) (token interface{}, err models.Error){
	DAOUser := DAOs.DAOUser{}
	passwordDB, err := DAOUser.GetUserPassword(email)
	if err != (models.Error{}){
		return
	}

	search := Exists{passwordDB}
	if search.Password == "" {
		err = models.ErrorResponse(errors.New("user not exist"), 403)
		return
	}

	token, err = controller.checkLogin(search,email, password)
	return
}

func (controller LoginController) checkLogin(search Exists, email string, password string) (dataToken interface{}, error models.Error){
	encrypt := EncryptController{enviroment.SecretKey}
	ePass, err := encrypt.Decrypt(search.Password)

	if ePass == password && err == (models.Error{}) {
		authControl := AuthController{}

		token,_ := authControl.GenerateAuth(email)
		dataToken = struct {Token string `json:"token"`}{Token:token}
		return
	}

	error = models.ErrorResponse(errors.New("password incorrect"), 403)
	return
}
