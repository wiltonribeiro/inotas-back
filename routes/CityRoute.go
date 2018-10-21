package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/controllers"
)

var CityRoute = models.Route{
	func (application* iris.Application){

		controller := controllers.LocationController{}

		application.Handle("GET", "/cities", func(ctx iris.Context) {
			data, err := controller.GetCities()
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
			ctx.JSON(data)
		})

		application.Handle("GET", "/cities/id/{id}", func(ctx iris.Context) {
			id := ctx.Params().Get("id")
			data, err := controller.GetCityById(id)
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
			ctx.JSON(data)
		})

		application.Handle("GET", "/cities/{state}", func(ctx iris.Context) {
			data, err := controller.GetCitiesByState(ctx.Params().Get("state"))
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
			ctx.JSON(data)
		})

		application.Handle("GET", "/cities/{state}/{city}", func(ctx iris.Context) {
			var state, city string
			state = ctx.Params().Get("state")
			city = ctx.Params().Get("city")

			data, err := controller.GetCityIdByStateAndName(state,city)

			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}

			var result = struct {
				Id int `json:"id"`
			}{data}
			ctx.JSON(result)
		})
	},
}
