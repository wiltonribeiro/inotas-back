package controllers

import (
	"inotas-back/models"
	"inotas-back/database"
	"strings"
)

type LocationController struct {
	DataBase* database.Connection
}

func (controller LocationController) GetStates() (states []models.State){
	query := "SELECT * FROM state"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	rows, err := stmt.Query()
	CheckFail(err)

	for rows.Next(){
		var state models.State
		rows.Scan(&state.Initials,&state.Name)
		states = append(states, state)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCities()(cities []models.City){
	query := "SELECT * FROM city LIMIT 10"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	rows, err := stmt.Query()
	CheckFail(err)

	for rows.Next(){
		var city models.City
		rows.Scan(&city.Id,&city.Name,&city.StateInitials)
		cities = append(cities, city)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCitiesByState(initials string)(cities []models.City){
	query := "SELECT * FROM city WHERE state_initials = $1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	rows, err := stmt.Query(initials)
	CheckFail(err)

	for rows.Next(){
		var city models.City
		rows.Scan(&city.Id,&city.StateInitials,&city.Name)
		cities = append(cities, city)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCityById(id string)(city models.City){
	query := "SELECT * FROM city where id = $1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	row := stmt.QueryRow(id)
	CheckFail(err)

	row.Scan(&city.Id,&city.Name,&city.StateInitials)
	return
}

func (controller LocationController) GetIdCityByStateAndName(initials string, city string)(id int){
	query := "SELECT id FROM city WHERE state_initials = $1 and name = $2"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	row := stmt.QueryRow(initials, strings.ToUpper(city))
	CheckFail(err)
	row.Scan(&id)
	return
}

