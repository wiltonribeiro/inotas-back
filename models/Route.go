package models

import (
	"github.com/kataras/iris"
	"inotas/database"
)

type Route struct {
	ApplyRoute func(application* iris.Application, con* database.Connection)
}
