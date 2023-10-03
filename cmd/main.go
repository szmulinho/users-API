package main

import (
	"github.com/szmulinho/users/cmd/server"
	"github.com/szmulinho/users/internal/database"
)

func main() {

	database.Connect()

	server.Run()
}
