package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/database"
	"inotas-back/controllers"
)

var NFeRoute = models.Route{
	func (application* iris.Application, con* database.Connection){

		controller := controllers.NFeController{DataBase:con}

		application.Handle("GET", "/nfe/{key}", func(ctx iris.Context) {
			data, err := controller.GetContent("willrcneto@gmail.com", ctx.Params().Get("key"))
			if err != nil {
				ctx.StatusCode(505)
				data = err
			}
			ctx.JSON(data)
		})
	},
}
