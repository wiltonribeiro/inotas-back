package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/database"
	"inotas-back/controllers"
)

var StateRoute = models.Route{
	func (application* iris.Application, con* database.Connection){

		controller := controllers.LocationController{DataBase:con}

		application.Handle("GET", "/states", func(ctx iris.Context) {
			data := controller.GetStates()
			ctx.JSON(data)
		})
	},
}
