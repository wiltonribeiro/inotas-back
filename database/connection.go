package database

import (
	_ "github.com/lib/pq"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"fmt"
	"database/sql"
)

type Connection struct {
	dbHost string
	dbName string
	dbUser string
	dbPassword string
	con *sql.DB
}

func CreateConnection(dbHost,dbName,dbUser,dbPassword string) (Connection,error) {

	//dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.dbUser, c.dbPassword, c.dbName)
	//db, err := sql.Open("postgres", dbInfo)

	c := Connection{dbHost, dbName, dbUser, dbPassword, nil}
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",c.dbHost, c.dbName, c.dbUser, c.dbPassword)
	db, err := sql.Open("cloudsqlpostgres", dsn)

	c.con = db
	return  c,err
}

func (c Connection) Close(){
	c.con.Close()
}

func (c Connection) GetDB() *sql.DB{
	return c.con
}