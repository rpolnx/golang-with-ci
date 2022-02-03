package routes

import (
	"rpolnx.com.br/golang-with-ci/src/controller"
	"rpolnx.com.br/golang-with-ci/src/handler"
)

func NewHealthcheckRoutes(h *handler.Handler, healthcheckController controller.HealthcheckController) {
	h.Gin.GET("/healthcheck", healthcheckController.GetHealthStatus)
}

func NewUserRoutes(h *handler.Handler, userController controller.UserController) {
	h.Gin.GET("/users", userController.GetAll)
	h.Gin.GET("/users/:id", userController.GetOne)
	h.Gin.POST("/users", userController.Post)
	h.Gin.PUT("/users/:id", userController.Put)
	h.Gin.DELETE("/users/:id", userController.Delete)
}
