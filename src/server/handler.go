package server

import (
	"github.com/gin-gonic/gin"
	"rpolnx.com.br/golang-with-ci/src/adapter"
	"rpolnx.com.br/golang-with-ci/src/controller"
	"rpolnx.com.br/golang-with-ci/src/repository"
	"rpolnx.com.br/golang-with-ci/src/routes"
	"rpolnx.com.br/golang-with-ci/src/service"
)

func InitializeServer() (*gin.Engine, error) {
	r := gin.Default()

	userRepo, err := repository.InitializeUserDatabaseClient()

	if err != nil {
		return nil, err
	}

	userController := getUserController(userRepo)

	healthcheckController := controller.InitializeHealthcheckController()

	routes.AppendHealthcheckRoutes(r, healthcheckController)
	routes.AppendUserRoutes(r, userController)

	return r, nil
}

func getUserController(userRepo repository.UserDBRepository) controller.UserController {
	userAdapter := adapter.InitializeUserAdapter(userRepo)
	userService := service.InitializeUserService(userAdapter)
	userController := controller.InitializeUserController(userService)
	return userController
}
