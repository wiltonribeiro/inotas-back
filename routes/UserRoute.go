package routes

import (
	"inotas-back/models"
	"github.com/kataras/iris"
	"inotas-back/controllers"
)

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var UserRoute = models.Route{
	ApplyRoute: func(application *iris.Application) {

		userController := controllers.UserController{}

		application.Handle("POST", "/login", func(ctx iris.Context) {

			loginController := controllers.LoginController{}

			var request loginRequest
			err := ctx.ReadJSON(&request)

			if err != nil {
				ctx.StatusCode(500)
				ctx.JSON(models.ErrorResponse(err, 500))
			} else {
				data, err := loginController.Login(request.Email, request.Password)
				if err != (models.Error{}){
					ctx.StatusCode(err.Code)
					ctx.JSON(err)
				} else {
					ctx.JSON(data)
				}

			}
		})

		application.Handle("GET", "/user", func(ctx iris.Context){
			user, err := userController.GetUser(ctx.GetHeader("Authorization"))
			if err != (models.Error{}) {
				ctx.StatusCode(err.Code)
				ctx.JSON(err)
			} else{
				ctx.JSON(user)
			}
		})

		application.Handle("POST", "/register", func(ctx iris.Context){
			var user models.User
			err := ctx.ReadJSON(&user)
			if err != nil {
				ctx.StatusCode(500)
				ctx.JSON(models.ErrorResponse(err, 500))
			} else {
				err := userController.Register(&user)
				if err != (models.Error{}){
					ctx.StatusCode(err.Code)
					ctx.JSON(err)
				}
				ctx.StatusCode(200)
			}
		})

		application.Handle("POST", "/changePassword", func(ctx iris.Context){
			request := struct {
				NewPassword string `json:"password"`
			}{}

			err := ctx.ReadJSON(&request)
			if err != nil {
				ctx.StatusCode(500)
				ctx.JSON(models.ErrorResponse(err, 500))
			} else {
				err := userController.ChangePassword(ctx.GetHeader("Authorization"), request.NewPassword)
				if err != (models.Error{}){
					ctx.StatusCode(err.Code)
					ctx.JSON(err)
				}
				ctx.StatusCode(200)
			}
		})
	},
}
