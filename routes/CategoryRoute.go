package routes

import (
	"inotas-back/models"
	"github.com/kataras/iris"
	"inotas-back/database"
	"inotas-back/controllers"
)

var CategoryRoute = models.Route{
	ApplyRoute: func(application *iris.Application, con *database.Connection) {

		controller := controllers.CategoryController{DataBase:con}
		application.Handle("GET", "/categories", func(ctx iris.Context) {
			data, err := controller.GetCategories()
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
			ctx.JSON(data)
		})
	},
}
