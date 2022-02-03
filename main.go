package main

import (
	"go.uber.org/fx"
	"rpolnx.com.br/golang-with-ci/src/server"
)

func main() {
	//serverContext, err := server.InitializeServer()
	//
	//if err != nil {
	//	log.Fatal("Server failed to initialize with error: ", err)
	//}

	fx.New(
		server.Module,
	).Run()

	//fatalErr := serverContext.Run(fmt.Sprintf("%s:%s", host, port))
	//
	//log.Fatal("Server crash with error: ", fatalErr)
}
