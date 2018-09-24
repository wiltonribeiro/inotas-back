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
			token := ctx.GetHeader("Authorization")
			data, err := controller.GetContent(token, ctx.Params().Get("key"))
			if err != nil {
				ctx.StatusCode(500)
				data = err
			}
			ctx.JSON(data)

		})
	},
}
