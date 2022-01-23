package routes

import (
	"github.com/gin-gonic/gin"
	"rpolnx.com.br/golang-with-ci/src/controller"
)

func InitializeRoutes() *gin.Engine {

	r := gin.Default()

	healthcheckController := controller.InitilizeHealthcheckController()

	r.GET("/healthcheck", healthcheckController.GetHealthStatus)

	return r
}
