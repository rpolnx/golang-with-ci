package server

import (
	"go.uber.org/fx"
	"rpolnx.com.br/golang-with-ci/src/adapter"
	"rpolnx.com.br/golang-with-ci/src/controller"
	"rpolnx.com.br/golang-with-ci/src/handler"
	"rpolnx.com.br/golang-with-ci/src/repository"
	"rpolnx.com.br/golang-with-ci/src/routes"
	"rpolnx.com.br/golang-with-ci/src/service"
	"rpolnx.com.br/golang-with-ci/src/util"
)

var RepositoryModule = fx.Options(fx.Provide(repository.InitializeUserDatabaseClient))
var AdapterModule = fx.Options(RepositoryModule, fx.Provide(adapter.InitializeUserAdapter))

var ServiceModule = fx.Options(AdapterModule, fx.Provide(service.InitializeUserService))
var ControllerModule = fx.Options(ServiceModule,
	fx.Provide(controller.InitializeUserController, controller.InitializeHealthcheckController),
)

var RoutesModule = fx.Options(ControllerModule,
	fx.Invoke(routes.NewHealthcheckRoutes),
	fx.Invoke(routes.NewUserRoutes),
)

var HandlerModule = fx.Options(fx.Provide(handler.NewHandler))

var Module = fx.Options(
	HandlerModule,
	RoutesModule,
	fx.Invoke(util.NewLoggers),
	fx.Invoke(handler.RegisterHooks),
)
