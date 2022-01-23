package main

import (
	"fmt"
	"log"
	"rpolnx.com.br/golang-with-ci/src/server"
)

const host = "localhost"
const port = "8080"

func main() {
	serverContext, err := server.InitializeServer()

	if err != nil {
		log.Fatal("Server failed to initialize with error: ", err)
	}

	fatalErr := serverContext.Run(fmt.Sprintf("%s:%s", host, port))

	log.Fatal("Server crash with error: ", fatalErr)
}
