package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"fmt"
	"inotas-back/enviroment"
	"errors"
)

type LoginController struct {
	DataBase* database.Connection
}

type Exists struct {
	Password string `json:"password"`
}

func (controller LoginController) Login(email, password string) (interface{}){
	query := "SELECT password FROM \"user\" WHERE email=$1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		return models.Error{
			Code:505,
			Message:fmt.Sprint(err),
		}
	}
	value := stmt.QueryRow(email)
	var search Exists
	value.Scan(&search.Password)

	if search.Password == "" {
		return  models.ErrorResponse(errors.New("user not exist"), 404)
	}

	return controller.checkLogin(search,email, password)
}

func (controller LoginController) checkLogin(search Exists, email string, password string) interface{}{
	encrypt := EncryptController{enviroment.SecretKey}
	ePass, err := encrypt.Decrypt(search.Password)

	if ePass == password && err == nil {
		authControl := AuthController{}

		token,_ := authControl.GenerateAuth(email)
		return struct {
			Token string `json:"token"`
		}{
			Token:token,
		}
	}

	return models.ErrorResponse(errors.New("password incorrect"), 404)
}
