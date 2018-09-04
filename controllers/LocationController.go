package controllers

import (
	"inotas-back/models"
	"inotas-back/database"
)

type LocationController struct {
	DataBase* database.Connection
}

func (controller LocationController) GetStates() (states []models.Estado){
	query := "SELECT * FROM estado"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	rows, err := stmt.Query()
	CheckFail(err)

	for rows.Next(){
		var state models.Estado
		rows.Scan(&state.Sigla,&state.Nome)
		states = append(states, state)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCities()(cities []models.Cidade){
	query := "SELECT nome,sigla_estado FROM cidade LIMIT 10"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	rows, err := stmt.Query()
	CheckFail(err)

	for rows.Next(){
		var city models.Cidade
		rows.Scan(&city.Nome,&city.SiglaEstado)
		cities = append(cities, city)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCitiesByState(initials string)(cities []models.Cidade){
	query := "SELECT * FROM cidade WHERE sigla_estado = $1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	rows, err := stmt.Query(initials)
	CheckFail(err)

	for rows.Next(){
		var city models.Cidade
		rows.Scan(&city.Id,&city.SiglaEstado,&city.Nome)
		cities = append(cities, city)
	}

	rows.Close()

	return
}

func (controller LocationController) GetCityById(id string)(city models.Cidade){
	query := "SELECT * FROM cidade where id = $1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	row := stmt.QueryRow(id)
	CheckFail(err)

	row.Scan(&city.Id,&city.Nome,&city.SiglaEstado)
	return
}

func (controller LocationController) GetIdCityByStateAndName(initials string, city string)(id int){
	query := "SELECT id FROM cidade where sigla_estado = $1 and nome = $2"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	row := stmt.QueryRow(initials, city)
	CheckFail(err)

	row.Scan(&id)
	return
}

