package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/controllers"
)

var NFeRoute = models.Route{
	func (application* iris.Application){

		controller := controllers.NFeController{}

		application.Handle("POST", "/nfe", func(ctx iris.Context) {
			var resultPost struct {
				NfeKey string `json:"nfe_key"`
			}

			err := ctx.ReadJSON(&resultPost)
			token := ctx.GetHeader("Authorization")
			var errorR models.Error

			if err != nil {
				errorR = models.ErrorResponse(err, 500)
				ctx.StatusCode(errorR.Code)
				ctx.JSON(errorR)
			}

			errorR = controller.GetContent(token, resultPost.NfeKey)
			if errorR != (models.Error{}) {
				ctx.StatusCode(errorR.Code)
				ctx.JSON(err)
			}
		})
	},
}
