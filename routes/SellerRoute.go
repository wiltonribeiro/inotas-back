package routes

import (
	"github.com/kataras/iris"
	"inotas-back/models"
	"inotas-back/controllers"
)

var SellerRoute = models.Route{
	func (application* iris.Application){

		controller := controllers.SellerController{}

		application.Handle("GET", "/seller/{cnpj}", func(ctx iris.Context) {
			key := ctx.Params().Get("cnpj")
			token := ctx.GetHeader("Authorization")

			data, err := controller.GetSeller(token, key)
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
			ctx.JSON(data)
		})
	},
}
