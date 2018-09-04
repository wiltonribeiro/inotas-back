package controllers

import (
	"inotas-back/models"
	"inotas-back/database"
	"fmt"
)

type LocationController struct {
	DataBase* database.Connection
}

func (controller LocationController) GetStates() (estados []models.Estado){
	query := "SELECT * FROM estado"
	rows, err := controller.DataBase.GetDB().Query(query)
	CheckFail(err)

	for rows.Next(){
		var estado models.Estado
		rows.Scan(&estado.Sigla,&estado.Nome)
		estados = append(estados, estado)
	}

	rows.Close()

	return
}

func (controller LocationController) GetCities()(cidades []models.Cidade){
	query := "SELECT nome,sigla_estado FROM cidade"
	rows, err := controller.DataBase.GetDB().Query(query)
	CheckFail(err)

	for rows.Next(){
		var cidade models.Cidade
		rows.Scan(&cidade.Nome,&cidade.SiglaEstado)
		cidades = append(cidades, cidade)
	}

	rows.Close()

	return
}

func (controller LocationController) GetCitiesByState(sigla string)(cidades []models.Cidade){
	query := fmt.Sprintf("SELECT * FROM cidade WHERE sigla_estado = '%v'", sigla)
	rows, err := controller.DataBase.GetDB().Query(query)
	CheckFail(err)

	for rows.Next(){
		var cidade models.Cidade
		rows.Scan(&cidade.Id,&cidade.SiglaEstado,&cidade.Nome)
		cidades = append(cidades, cidade)
	}

	rows.Close()

	return
}

func (controller LocationController) GetCityById(id string)(cidade models.Cidade){
	query := fmt.Sprintf("SELECT * FROM cidade where id = '%v'",id)
	row := controller.DataBase.GetDB().QueryRow(query)
	row.Scan(&cidade.Id,&cidade.Nome,&cidade.SiglaEstado)
	return
}

func (controller LocationController) GetIdCityByStateAndName(sigla string, cidade string)(id int){
	query := fmt.Sprintf("SELECT id FROM cidade where sigla_estado = '%v' and nome = '%v'", sigla, cidade)
	row := controller.DataBase.GetDB().QueryRow(query)
	row.Scan(&id)
	return
}

