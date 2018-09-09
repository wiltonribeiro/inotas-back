package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"fmt"
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
		return models.Error{
			Code:403,
			Message:fmt.Sprint(err),
		}
	}

	_,err = stmt.Exec(user.Email, user.Password, user.CityId, user.StateInitials, user.Name)
	if err != nil {
		return models.Error{
			Code:409,
			Message:fmt.Sprint(err),
		}
	}

	return
}

func (controller UserController) ChangePassword(auth, newPassword string) (error models.Error){
	authControl := AuthController{}
	encryptControl := EncryptController{enviroment.SecretKey}

	email, err  := authControl.CheckAuth(auth)
	if err != nil {
		return models.Error{
			Code:403,
			Message:fmt.Sprint(err),
		}
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
		return models.Error{
			Code:505,
			Message:fmt.Sprint(err),
		}
	}
	stmt.Exec(encryptPass,email)
	return
}
