package routes

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"rpolnx.com.br/golang-with-ci/docs"
	"rpolnx.com.br/golang-with-ci/src/controller"
	"rpolnx.com.br/golang-with-ci/src/handler"
	"rpolnx.com.br/golang-with-ci/src/util"
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

func NewSwaggerRoutes(h *handler.Handler) {
	docs.SwaggerInfo_swagger.Title = ""
	util.Logger.Info("Generating docs from", docs.SwaggerInfo_swagger.Title)

	h.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
