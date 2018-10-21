package routes

import (
	"inotas-back/models"
	"github.com/kataras/iris"
)

var WelcomeRoute = models.Route{
	ApplyRoute: func(application *iris.Application) {
		application.Handle("GET", "/", func(ctx iris.Context) {
			hello := "Welcome to wawona.inotas API"
			ctx.HTML(hello)
		})
	},
}
