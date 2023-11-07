package main

import (
	"fmt"
	"github.com/szmulinho/common/utils"
	"github.com/szmulinho/users/internal/database"
	"github.com/szmulinho/users/internal/server"
	"log"
)

func main() {
	fmt.Println("Staring the application...")
	defer fmt.Println("Closing the application...")

	db, err := database.Connect()
	if err != nil {
		log.Fatalf("connecting to dataabse: %v", err)
	}

	ctx, _, wait := utils.Gracefully()

	server.Run(ctx, db)

	wait()
}
