package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"log"
)

const host = ""
const port = "8080"

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
				serverHost := fmt.Sprintf("%s:%s", host, port)

				fmt.Println("Starting application in ", serverHost)
				go func() {
					log.Fatal(h.Gin.Run(serverHost))
				}()
				return nil
			},
			OnStop: func(context.Context) error {
				fmt.Println("Stopping application")
				return nil
			},
		},
	)
}
