package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"os"
)

type Handler struct {
	Gin *gin.Engine
}

// NewHandler returns a new gin router
func NewHandler() *Handler {
	r := gin.Default()

	handler := Handler{Gin: r}

	return &handler
}

func RegisterHooks(lifecycle fx.Lifecycle, h *Handler) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				serverHost := fmt.Sprintf("%s:%s", os.Getenv("APP_HOST"), os.Getenv("APP_PORT"))

				fmt.Println("Starting application in ", serverHost)

				return h.Gin.Run(serverHost)
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
