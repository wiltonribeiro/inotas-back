package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/database"
	"inotas-back/controllers"
)

var CityRoute = models.Route{
	func (application* iris.Application, con* database.Connection){

		controller := controllers.LocationController{DataBase:con}

		application.Handle("GET", "/cities", func(ctx iris.Context) {
			data := controller.GetCities()
			ctx.JSON(data)
		})

		application.Handle("GET", "/cities/id/{id}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			data := controller.GetCityById(id)
			ctx.JSON(data)
		})

		application.Handle("GET", "/cities/{state}", func(ctx iris.Context) {
			data := controller.GetCitiesByState(ctx.Params().Get("state"))
			ctx.JSON(data)
		})

		application.Handle("GET", "/cities/{state}/{city}", func(ctx iris.Context) {
			state := ctx.Params().Get("state")
			city := ctx.Params().Get("city")
			data := controller.GetIdCityByStateAndName(state,city)
			ctx.JSON(struct {
				Id int `json:"id"`
			}{data})
		})
	},
}
