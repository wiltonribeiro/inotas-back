package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/controllers"
)

var NFeRoute = models.Route{
	func (application* iris.Application){

		controller := controllers.NFeController{}

		application.Handle("GET", "/nfe/{key}", func(ctx iris.Context) {
			token := ctx.GetHeader("Authorization")
			err := controller.GetContent(token, ctx.Params().Get("key"))
			if err != (models.Error{}) {
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
		})
	},
}
