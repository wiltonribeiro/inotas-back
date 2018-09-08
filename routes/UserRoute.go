package routes

import (
	"inotas-back/models"
	"github.com/kataras/iris"
	"inotas-back/database"
	"inotas-back/controllers"
	"fmt"
)

type loginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

var UserRoute = models.Route{
	ApplyRoute: func(application *iris.Application, con *database.Connection) {

		userController := controllers.UserController{DataBase:con}

		application.Handle("POST", "/login", func(ctx iris.Context) {

			loginController := controllers.LoginController{DataBase:con}

			var request loginRequest
			err := ctx.ReadJSON(&request)

			if err != nil {
				ctx.StatusCode(505)
				ctx.JSON(models.Error{
					Message:fmt.Sprint(err),
					Code:505,
				})
			} else {
				data := loginController.Login(request.Email, request.Password)
				switch data.(type) {
				case models.Error:
					ctx.StatusCode(data.(models.Error).Code)
				}
				ctx.JSON(data)
			}
		})

		application.Handle("POST", "/register", func(ctx iris.Context){
			var user models.Usuario
			err := ctx.ReadJSON(&user)

			if err != nil {
				ctx.StatusCode(505)
				ctx.JSON(models.Error{
					Message:fmt.Sprint(err),
					Code:505,
				})
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
				Token string `json:"token"`
				NewPassword string `json:"password"`
			}{}

			err := ctx.ReadJSON(&request)
			if err != nil {
				ctx.StatusCode(505)
				ctx.JSON(models.Error{
					Message:fmt.Sprint(err),
					Code:505,
				})
			} else {
				err := userController.ChangePassword(request.Token, request.NewPassword)
				if err != (models.Error{}){
					ctx.StatusCode(err.Code)
					ctx.JSON(err)
				}
				ctx.StatusCode(200)
			}
		})
	},
}
