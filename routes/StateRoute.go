package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/controllers"
)

var StateRoute = models.Route{
	func (application* iris.Application){

		controller := controllers.LocationController{}

		application.Handle("GET", "/states", func(ctx iris.Context) {
			data, err := controller.GetStates()
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
			ctx.JSON(data)
		})
	},
}
