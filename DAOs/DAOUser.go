package DAOs

import (
	"inotas-back/models"
	"fmt"
	"inotas-back/database"
)

type DAOUser struct {}

func (dao *DAOUser) SaveUser(user* models.User) (error models.Error){
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "INSERT INTO \"user\" (email,password,city_id,state_initials,name) VALUES ($1,$2,$3,$4,$5)"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil {
		return  models.ErrorResponse(err, 401)
	}

	_,err = stmt.Exec(user.Email, user.Password, user.CityId, user.StateInitials, user.Name)
	if err != nil {
		fmt.Println(user.StateInitials)
		return models.ErrorResponse(err, 409)
	}
	return
}


func (dao *DAOUser) UpdatePassword(encryptPass, email string) (error models.Error) {
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "UPDATE \"user\" SET password = $1 WHERE email = $2"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil {
		return models.ErrorResponse(err, 500)
	}
	stmt.Exec(encryptPass,email)
	return
}

func (dao *DAOUser) GetUser(email string) (user models.User , error models.Error){
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "SELECT city_id,state_initials,name FROM  \"user\" WHERE email = $1"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil {
		error = models.ErrorResponse(err, 500)
		return
	}
	row := stmt.QueryRow(email)
	row.Scan(&user.CityId, &user.StateInitials, &user.Name)
	return
}

func (dao *DAOUser) GetUserPassword(email string) (password string, error models.Error) {
	con, err := Database.OpenConnection()
	defer con.Close()

	query := "SELECT password FROM \"user\" WHERE email=$1"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil {
		error = models.ErrorResponse(err, 500)
		return
	}
	value := stmt.QueryRow(email)
	value.Scan(&password)

	return
}