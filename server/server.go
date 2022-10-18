package server

import (
	"github.com/AuthSystemJWT/deyki/v2/controller"
	"github.com/AuthSystemJWT/deyki/v2/database"
)

func Run() {

	database.LoadEnvVariables()
	database.ConnectDB()
	controller.GinRouter()
}