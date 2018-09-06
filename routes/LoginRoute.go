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

var LoginRoute = models.Route{
	ApplyRoute: func(application *iris.Application, con *database.Connection) {
		controller := controllers.LoginController{DataBase:con}

		var request loginRequest
		application.Handle("POST", "/login", func(ctx iris.Context) {
			err := ctx.ReadJSON(&request)
			if err != nil {
				ctx.JSON(models.Error{
					Message:fmt.Sprint(err),
					Code:505,
				})
			}

			data := controller.Login(request.Email, request.Password)
			ctx.JSON(data)
		})
	},
}
