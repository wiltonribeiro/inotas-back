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

		application.Handle("GET", "/cidades", func(ctx iris.Context) {
			data := controller.GetCities()
			ctx.JSON(data)
		})

		application.Handle("GET", "/cidades/id/{id}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			data := controller.GetCityById(id)
			ctx.JSON(data)
		})

		application.Handle("GET", "/cidades/{estado}", func(ctx iris.Context) {
			data := controller.GetCitiesByState(ctx.Params().Get("estado"))
			ctx.JSON(data)
		})

		application.Handle("GET", "/cidades/{estado}/{cidade}", func(ctx iris.Context) {
			estado := ctx.Params().Get("estado")
			cidade := ctx.Params().Get("cidade")
			data := controller.GetIdCityByStateAndName(estado,cidade)
			ctx.JSON(struct {
				Id int `json:"id"`
			}{data})
		})
	},
}
