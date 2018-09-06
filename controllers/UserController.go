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

func (controller UserController) Register(user* models.Usuario) (error models.Error){

	var result string
	encrypt := EncryptController{enviroment.SecretKey}
	result, error = encrypt.Encrypt([]byte(user.Senha))
	if error != (models.Error{}) {
		return error
	}

	user.Senha = result
	query := "INSERT INTO usuario(email,senha,id_cidade,sigla_estado,nome) VALUES ($1,$2,$3,$4,$5)"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil {
		return models.Error{
			Code:403,
			Message:fmt.Sprint(err),
		}
	}

	_,err = stmt.Exec(user.Email, user.Senha, user.IdCidade, user.SiglaEstado, user.Nome)
	if err != nil {
		return models.Error{
			Code:409,
			Message:fmt.Sprint(err),
		}
	}

	return
}
