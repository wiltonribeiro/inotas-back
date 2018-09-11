package server

import (
	"log"
	"inotas-back/database"
	"github.com/kataras/iris"
	"inotas-back/enviroment"
	"inotas-back/models"
	"inotas-back/routes"
	"github.com/iris-contrib/middleware/cors"
)

func checkFail(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func initDB() (con database.Connection){
	con, err := database.CreateConnection(enviroment.DbUser,enviroment.DbPassword,enviroment.DbName)
	checkFail(err)
	return
}

func initRoutes(routes []models.Route, db* database.Connection){

	app := iris.Default()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
	})
	app.Use(crs)

	for _, item := range routes {
		 item.ApplyRoute(app, db)
	}
	app.Run(iris.Addr(":8080"))
}

func InitAll(){
	db := initDB()
	r := []models.Route{
		routes.StateRoute,
		routes.CityRoute,
		routes.NFeRoute,
		routes.UserRoute,
		routes.CategoryRoute,
		routes.ShopRoute,
	}
	initRoutes(r, &db)
	defer db.Close()
}
