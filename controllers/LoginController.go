package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"fmt"
	"inotas-back/enviroment"
)

type LoginController struct {
	DataBase* database.Connection
}

type Exists struct {
	Password string `json:"password"`
}

func (controller LoginController) Login(email, password string) (interface{}){
	query := "SELECT senha FROM usuario WHERE email=$1"
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
		return models.Error{
			Code:404,
			Message:"User not exist",
		}
	}

	return controller.checkLogin(search,email, password)
}

func (controller LoginController) checkLogin(search Exists, email string, password string) interface{}{
	encrypt := EncryptController{enviroment.SecretKey}
	ePass, err := encrypt.Decrypt(search.Password)

	if ePass == password && err == nil {
		authControl := AuthController{
			Email:email,
		}

		token,_ := authControl.GenerateAuth()
		return struct {
			Token string `json:"token"`
		}{
			Token:token,
		}
	}

	return models.Error{
		Code:404,
		Message:"Password incorrect",
	}
}
