package handler

import (
	"context"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"os"
	"rpolnx.com.br/golang-with-ci/src/util"
	"time"
)

type Handler struct {
	Gin *gin.Engine
}

// NewHandler returns a new gin router
func NewHandler() *Handler {
	engine := gin.New()

	isLocal := os.Getenv("IS_LOCAL")

	if isLocal != "" {
		engine.Use(ginzap.Ginzap(util.NamedLogger, time.RFC3339, true))
		engine.Use(ginzap.RecoveryWithZap(util.NamedLogger, true))
	} else {
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
	}

	handler := Handler{Gin: engine}

	return &handler
}

func RegisterHooks(lifecycle fx.Lifecycle, h *Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				host := os.Getenv("APP_HOST")
				port := os.Getenv("APP_PORT")

				if port == "" {
					port = "8080"
				}
				serverHost := fmt.Sprintf("%s:%s", host, port)

				util.Logger.Info("Starting application in ", serverHost)

				go h.Gin.Run(serverHost)

				return nil
			},
			OnStop: func(context.Context) error {
				util.Logger.Info("Stopping application")
				return nil
			},
		},
	)
}
