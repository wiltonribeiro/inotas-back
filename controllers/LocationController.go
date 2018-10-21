package controllers

import (
	"inotas-back/models"
	"inotas-back/DAOs"
)

type LocationController struct {}

func (controller LocationController) GetStates() ([]models.State, models.Error){
	DAOStates := DAOs.DAOState{}
	return DAOStates.ListStates()
}

func (controller LocationController) GetCities()([]models.City, models.Error){
	DAOCity := DAOs.DAOCity{}
	return DAOCity.ListCities()
}

func (controller LocationController) GetCitiesByState(initials string)([]models.City, models.Error){
	DAOCity := DAOs.DAOCity{}
	return DAOCity.ListCitiesByState(initials)
}

func (controller LocationController) GetCityById(id string)(models.City, models.Error){
	DAOCity := DAOs.DAOCity{}
	return DAOCity.GetCityById(id)
}

func (controller LocationController) GetCityIdByStateAndName(initials string, city string)(int, models.Error){
	DAOCity := DAOs.DAOCity{}
	return DAOCity.GetCityIdByStateAndName(initials, city)
}

