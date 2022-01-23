package server

import (
	"github.com/gin-gonic/gin"
	"rpolnx.com.br/golang-with-ci/src/routes"
)

func InitializeServer() (*gin.Engine, error) {

	routesHandler := routes.InitializeRoutes()

	return routesHandler, nil
}
