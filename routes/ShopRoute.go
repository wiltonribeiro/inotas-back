package routes

import (
	"inotas-back/models"
	"github.com/kataras/iris"
	"inotas-back/database"
	"inotas-back/controllers"
	)

var ShopRoute = models.Route{
	ApplyRoute: func(application *iris.Application, con *database.Connection) {

		controller := controllers.ShopController{DataBase:con}
		controlAuth := controllers.AuthController{}

		application.Handle("POST", "/shop/products/update", func(ctx iris.Context) {
			var resultPost struct {
				Products []models.Product `json:"products"`
			}

			err := ctx.ReadJSON(&resultPost)
			_ , errAuth := controlAuth.CheckAuth(ctx.GetHeader("Authorization"))

			if errAuth != (models.Error{}){
				ctx.StatusCode(errAuth.Code)
			} else if err != nil {
				ctx.StatusCode(505)
				ctx.JSON(models.ErrorResponse(err,505))
			} else if result := controller.UpdateProductsCategories(resultPost.Products); result != (models.Error{}){
				ctx.StatusCode(result.Code)
				ctx.JSON(result)
			}
		})

		application.Handle("POST", "/shop/alias/update", func(ctx iris.Context) {
			var resultPost struct {
				Shop models.Shop `json:"shop"`
			}

			err := ctx.ReadJSON(&resultPost)
			_ , errAuth := controlAuth.CheckAuth(ctx.GetHeader("Authorization"))

			if errAuth != (models.Error{}){
				ctx.StatusCode(errAuth.Code)
			} else if err != nil {
				ctx.StatusCode(505)
				ctx.JSON(models.ErrorResponse(err,505))
			} else if err := controller.UpdateShopAlias(resultPost.Shop); err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			}
		})

		application.Handle("GET", "/shop/all", func(ctx iris.Context){
			token := ctx.GetHeader("Authorization")
			result, err := controller.GetShop(token)
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			} else{
				ctx.JSON(result)
			}
		})
	},
}
