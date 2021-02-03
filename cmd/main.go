package main

import (
	"log"
	"os"

	"github.com/org-struct-hexagonal-architecture-/api"
)

const defaultPort = "8080"

func main() {
	log.Println("Comenzando API cmd")
	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}
	api.Start(port)
}
