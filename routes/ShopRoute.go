package routes

import (
	"inotas-back/models"
	"github.com/kataras/iris"
	"inotas-back/controllers"
	)

var ShopRoute = models.Route{
	ApplyRoute: func(application *iris.Application) {

		controller := controllers.ShopController{}

		application.Handle("POST", "/shop/products/update", func(ctx iris.Context) {
			var resultPost struct {
				Products []models.Product `json:"products"`
			}

			err := ctx.ReadJSON(&resultPost)
			token := ctx.GetHeader("Authorization")

			if err != nil {
				ctx.StatusCode(400)
				ctx.JSON(models.ErrorResponse(err,400))
			} else if result := controller.UpdateProductsCategories(token,resultPost.Products); result != (models.Error{}){
				ctx.StatusCode(result.Code)
				ctx.JSON(result)
			}
		})

		application.Handle("POST", "/shop/alias/update", func(ctx iris.Context) {
			var resultPost struct {
				Shop models.Shop `json:"shop"`
			}

			err := ctx.ReadJSON(&resultPost)
			token := ctx.GetHeader("Authorization")

			if err != nil {
				ctx.StatusCode(400)
				ctx.JSON(models.ErrorResponse(err,400))
			} else if err := controller.UpdateShopAlias(token,resultPost.Shop); err != (models.Error{}){
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

		application.Handle("GET", "/shop/items/{nfe}", func(ctx iris.Context){
			nfe := ctx.Params().Get("nfe")
			token := ctx.GetHeader("Authorization")
			result, err := controller.GetItems(token, nfe)
			if err != (models.Error{}){
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			} else{
				ctx.JSON(result)
			}
		})
	},
}
