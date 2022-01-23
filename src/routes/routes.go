package routes

import (
	"github.com/gin-gonic/gin"
	"rpolnx.com.br/golang-with-ci/src/controller"
)

func AppendHealthcheckRoutes(r *gin.Engine, healthcheckController controller.HealthcheckController) {
	r.GET("/healthcheck", healthcheckController.GetHealthStatus)
}

func AppendUserRoutes(r *gin.Engine, userController controller.UserController) {
	r.GET("/users", userController.GetAll)
	r.GET("/users/:id", userController.GetOne)
	r.POST("/users", userController.Post)
	r.PUT("/users/:id", userController.Put)
	r.DELETE("/users/:id", userController.Delete)
}
