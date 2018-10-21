package controllers

import (
	"inotas-back/models"
	"inotas-back/DAOs"
)

type CategoryController struct {}

func (controller CategoryController) GetCategories() ([]models.Category, models.Error){
	DAOCategory := DAOs.DAOCategory{}
	return DAOCategory.ListCategories()
}