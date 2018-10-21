package models

import (
	"github.com/kataras/iris"
)

type Route struct {
	ApplyRoute func(application* iris.Application)
}
