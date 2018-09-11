package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"inotas-back/enviroment"
)

type UserController struct {
	DataBase* database.Connection
}

func (controller UserController) Register(user* models.User) (error models.Error){

	var result string
	encrypt := EncryptController{enviroment.SecretKey}
	result, error = encrypt.Encrypt([]byte(user.Password))
	if error != (models.Error{}) {
		return error
	}

	user.Password = result
	query := "INSERT INTO \"user\" (email,password,city_id,state_initials,name) VALUES ($1,$2,$3,$4,$5)"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		return  models.ErrorResponse(err, 403)
	}

	_,err = stmt.Exec(user.Email, user.Password, user.CityId, user.StateInitials, user.Name)
	if err != nil {
		return models.ErrorResponse(err, 409)
	}

	return
}

func (controller UserController) ChangePassword(token, newPassword string) (error models.Error){
	authControl := AuthController{}
	encryptControl := EncryptController{enviroment.SecretKey}

	email, err  := authControl.CheckAuth(token)
	if err != nil {
		return models.ErrorResponse(err, 403)
	} else {
		password, err := encryptControl.Encrypt([]byte(newPassword))
		if err != (models.Error{}){
			return err
		}
		return controller.updatePassword(password, email)
	}
}

func (controller UserController) updatePassword(encryptPass, email string) (error models.Error) {
	query := "UPDATE \"user\" SET password = $1 WHERE email = $2"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		return models.ErrorResponse(err, 505)
	}
	stmt.Exec(encryptPass,email)
	return
}
