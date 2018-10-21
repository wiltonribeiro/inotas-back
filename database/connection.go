package Database

import (
	_ "github.com/lib/pq"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"fmt"
	"database/sql"
	"inotas-back/enviroment"
)

type Connection struct {
	dbHost string
	dbName string
	dbUser string
	dbPassword string
	con *sql.DB
}

func OpenConnection() (Connection,error) {
	c := Connection{enviroment.DbHost, enviroment.DbName, enviroment.DbUser, enviroment.DbPassword, nil}
	dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",c.dbHost, c.dbName, c.dbUser, c.dbPassword)
	db, err := sql.Open("cloudsqlpostgres", dsn)
	c.con = db
	return  c,err
}

func (c Connection) Close(){
	c.con.Close()
}

func (c Connection) GetDB() (*sql.DB){
	return c.con
}

//dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", c.dbUser, c.dbPassword, c.dbName)
//db, err := sql.Open("postgres", dbInfo)