package DAOs

import (
	"inotas-back/database"
	"inotas-back/models"
)

type DAOState struct {}

func (dao *DAOState) ListStates() (states []models.State, error models.Error){
	con, err := Database.OpenConnection()

	query := "SELECT * FROM state"
	stmt, err := con.GetDB().Prepare(query)
	if err != nil { return states, models.ErrorResponse(err, 500)}

	rows, err := stmt.Query()
	if err != nil { return states, models.ErrorResponse(err, 500)}

	for rows.Next(){
		var state models.State
		rows.Scan(&state.Initials,&state.Name)
		states = append(states, state)
	}

	rows.Close()
	con.Close()
	return
}

