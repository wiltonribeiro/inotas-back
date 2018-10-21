package server

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/routes"
	"github.com/iris-contrib/middleware/cors"
)

func initRoutes(routes []models.Route){

	app := iris.Default()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.Use(crs)

	for _, item := range routes {
		 item.ApplyRoute(app)
	}
	app.Run(iris.Addr(":8080"))
}

func InitAll(){
	r := []models.Route{
		routes.StateRoute,
		routes.CityRoute,
		routes.NFeRoute,
		routes.UserRoute,
		routes.CategoryRoute,
		routes.ShopRoute,
		routes.WelcomeRoute,
	}
	initRoutes(r)
}
