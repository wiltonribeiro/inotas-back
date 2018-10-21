package DAOs

import (
	"inotas-back/database"
	"inotas-back/models"
)

type DAOCategory struct {}

func (dao *DAOCategory) ListCategories() (data []models.Category, error models.Error){
	con, err := Database.OpenConnection()

	query := "SELECT * FROM category"
	rows, err := con.GetDB().Query(query)
	if err != nil {
		return data, models.ErrorResponse(err, 505)
	}

	for rows.Next(){
		var category models.Category
		rows.Scan(&category.Id,&category.Name, &category.StateInitials)
		data = append(data, category)
	}

	rows.Close()
	con.Close()
	return
}

