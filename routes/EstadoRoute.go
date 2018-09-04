package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/database"
	"inotas-back/controllers"
)

var EstadoRoute = models.Route{
	func (application* iris.Application, con* database.Connection){

		controller := controllers.LocationController{con}

		application.Handle("GET", "/estados", func(ctx iris.Context) {
			data := controller.GetStates()
			ctx.JSON(data)
		})
	},
}
