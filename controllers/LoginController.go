package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"inotas-back/enviroment"
	"errors"
)

type LoginController struct {
	DataBase* database.Connection
}

type Exists struct {
	Password string `json:"password"`
}

func (controller LoginController) Login(email, password string) (token interface{}, error models.Error){
	query := "SELECT password FROM \"user\" WHERE email=$1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		error = models.ErrorResponse(err, 500)
		return
	}

	value := stmt.QueryRow(email)
	var search Exists
	value.Scan(&search.Password)

	if search.Password == "" {
		error = models.ErrorResponse(errors.New("user not exist"), 403)
		return
	}

	token, error = controller.checkLogin(search,email, password)
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
