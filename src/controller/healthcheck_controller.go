package controller

import "github.com/gin-gonic/gin"

type HealthcheckController interface {
	GetHealthStatus(c *gin.Context)
}

type healthcheckController struct{}

func (*healthcheckController) GetHealthStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"STATUS": "UP",
	})
}

func InitializeHealthcheckController() HealthcheckController {
	return &healthcheckController{}
}
