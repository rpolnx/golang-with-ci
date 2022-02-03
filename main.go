package main

import (
	"go.uber.org/fx"
	"rpolnx.com.br/golang-with-ci/src/server"
)

func main() {
	fx.New(
		server.Module,
	).Run()
}
