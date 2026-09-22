package main

import (
	"log"

	"github.com/alan-content-platform/alan-content-platform-api/internal/server"
)

func main() {
	srv := server.New()

	if err := srv.Run(":8080"); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
