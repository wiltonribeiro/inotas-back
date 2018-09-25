package controllers

import (
	"inotas-back/models"
	"inotas-back/database"
	"strings"
)

type LocationController struct {
	DataBase* database.Connection
}

func (controller LocationController) GetStates() (states []models.State, error models.Error){
	query := "SELECT * FROM state"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil { return states, models.ErrorResponse(err, 500)}

	rows, err := stmt.Query()
	if err != nil { return states, models.ErrorResponse(err, 500)}

	for rows.Next(){
		var state models.State
		rows.Scan(&state.Initials,&state.Name)
		states = append(states, state)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCities()(cities []models.City, error models.Error){
	query := "SELECT * FROM city LIMIT 10"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil { return cities, models.ErrorResponse(err, 500)}

	rows, err := stmt.Query()
	if err != nil { return cities, models.ErrorResponse(err, 500)}


	for rows.Next(){
		var city models.City
		rows.Scan(&city.Id,&city.Name,&city.StateInitials)
		cities = append(cities, city)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCitiesByState(initials string)(cities []models.City, error models.Error){
	query := "SELECT * FROM city WHERE state_initials = $1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil { return cities, models.ErrorResponse(err, 500)}

	rows, err := stmt.Query(initials)
	if err != nil { return cities, models.ErrorResponse(err, 500)}

	for rows.Next(){
		var city models.City
		rows.Scan(&city.Id,&city.StateInitials,&city.Name)
		cities = append(cities, city)
	}

	rows.Close()
	return
}

func (controller LocationController) GetCityById(id string)(city models.City, error models.Error){
	query := "SELECT * FROM city where id = $1"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil { return city, models.ErrorResponse(err, 500)}

	row := stmt.QueryRow(id)
	if err != nil { return city, models.ErrorResponse(err, 500)}

	row.Scan(&city.Id,&city.Name,&city.StateInitials)
	return
}

func (controller LocationController) GetIdCityByStateAndName(initials string, city string)(id int, error models.Error){
	query := "SELECT id FROM city WHERE state_initials = $1 and name = $2"
	stmt, err := controller.DataBase.GetDB().Prepare(query)
	if err != nil { return 0, models.ErrorResponse(err, 500)}

	row := stmt.QueryRow(initials, strings.ToUpper(city))
	if err != nil { return 0, models.ErrorResponse(err, 500)}
	row.Scan(&id)
	return
}

