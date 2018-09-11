package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
)

type CategoryController struct {
	DataBase* database.Connection
}

func (controller CategoryController) GetCategories() (data []models.Category, error models.Error){
	query := "SELECT * FROM category"
	rows, err := controller.DataBase.GetDB().Query(query)
	if err != nil {
		return data, models.ErrorResponse(err, 505)
	}

	for rows.Next(){
		var category models.Category
		rows.Scan(&category.Id,&category.Name, &category.StateInitials)
		data = append(data, category)
	}
	rows.Close()
	return
}