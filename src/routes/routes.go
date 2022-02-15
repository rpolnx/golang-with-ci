package routes

import (
	"github.com/gin-gonic/gin"
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
	group := h.Gin.Group("/users")
	{
		group.GET("/", userController.GetAll)
		group.GET("/:id", userController.GetOne)
		group.POST("/", userController.Post)
		group.PUT("/:id", userController.Put)
		group.DELETE("/:id", userController.Delete)
	}

}

func NewSwaggerRoutes(h *handler.Handler) {
	util.Logger.Infof("Generating docs from %s", docs.SwaggerInfo_swagger.Title)

	h.Gin.GET("/swagger/*any", func(c *gin.Context) {
		if c.Request.URL.Path == "" {
			c.Redirect(301, "/swagger/index.html")
			return
		}

		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})
}
