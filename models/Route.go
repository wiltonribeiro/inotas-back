package models

import (
	"github.com/kataras/iris"
	"inotas-back/database"
)

type Route struct {
	ApplyRoute func(application* iris.Application, con* database.Connection)
}
