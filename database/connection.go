package database

import (
	_ "github.com/lib/pq"
	"fmt"
	"database/sql"
)

type Connection struct {
	dbUser string
	dbPassword string
	dbName string
	con *sql.DB
}

func CreateConnection(dbUser,dbPassword,dbName string) (Connection,error) {
	c := Connection{dbUser, dbPassword, dbName, nil}
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.dbUser, c.dbPassword, c.dbName)
	db, err := sql.Open("postgres", dbInfo)
	c.con = db
	return  c,err
}

func (c Connection) Close(){
	c.con.Close()
}

func (c Connection) GetDB() *sql.DB{
	return c.con
}