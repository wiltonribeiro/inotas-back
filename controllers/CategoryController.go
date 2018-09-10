package controllers

import (
	"inotas-back/database"
	"inotas-back/models"
	"fmt"
)

type CategoryController struct {
	DataBase* database.Connection
}

func (controller CategoryController) GetCategories() (data []models.Category, error models.Error){
	query := "SELECT * FROM category"
	rows, err := controller.DataBase.GetDB().Query(query)
	if err != nil {
		return data,models.Error{
			Code:505,
			Message:fmt.Sprint(err),
		}
	}

	for rows.Next(){
		var category models.Category
		rows.Scan(&category.Id,&category.Name, &category.StateInitials)
		data = append(data, category)
	}

	return
}
