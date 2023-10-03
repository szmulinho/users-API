package main

import (
	"github.com/szmulinho/prescription/cmd/server"
	"github.com/szmulinho/prescription/internal/database"
)

func main() {

	database.Connect()

	server.Run()
}
