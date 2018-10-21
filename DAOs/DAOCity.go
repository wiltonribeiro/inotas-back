package DAOs

import (
	"inotas-back/database"
	"inotas-back/models"
	"strings"
)

type DAOCity struct {}

func (dao *DAOCity) ListCities() (cities []models.City, error models.Error){
	con, err := Database.OpenConnection()

	query := "SELECT * FROM city LIMIT 10"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil { return cities, models.ErrorResponse(err, 500)}

	rows, err := stmt.Query()
	if err != nil { return cities, models.ErrorResponse(err, 500)}


	for rows.Next(){
		var city models.City
		rows.Scan(&city.Id,&city.Name,&city.StateInitials)
		cities = append(cities, city)
	}

	rows.Close()
	con.Close()
	return
}

func (dao *DAOCity) ListCitiesByState(initials string) (cities []models.City, error models.Error){
	con, err := Database.OpenConnection()

	query := "SELECT * FROM city WHERE state_initials = $1"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil { return cities, models.ErrorResponse(err, 500)}

	rows, err := stmt.Query(initials)
	if err != nil { return cities, models.ErrorResponse(err, 500)}

	for rows.Next(){
		var city models.City
		rows.Scan(&city.Id,&city.StateInitials,&city.Name)
		cities = append(cities, city)
	}

	rows.Close()
	con.Close()
	return
}

func (dao *DAOCity) GetCityById(id string) (city models.City, error models.Error){
	con, err := Database.OpenConnection()

	query := "SELECT * FROM city where id = $1"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil { return city, models.ErrorResponse(err, 500)}

	row := stmt.QueryRow(id)
	if err != nil { return city, models.ErrorResponse(err, 500)}

	row.Scan(&city.Id,&city.Name,&city.StateInitials)

	con.Close()
	return
}

func (dao *DAOCity) GetCityIdByStateAndName(initials string, city string)(id int, error models.Error){
	con, err := Database.OpenConnection()

	query := "SELECT id FROM city WHERE state_initials = $1 and name = $2"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil { return 0, models.ErrorResponse(err, 500)}

	row := stmt.QueryRow(initials, strings.ToUpper(city))
	if err != nil { return 0, models.ErrorResponse(err, 500)}
	row.Scan(&id)

	con.Close()
	return
}

