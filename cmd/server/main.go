package main

import (
	"log"
	"os"

	"github.com/ritu-p/go-rest/internal/api"
	"github.com/ritu-p/go-rest/internal/db"
	logger "github.com/ritu-p/go-rest/pkg"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("could not connect to db: %v", err)
	}

	router := api.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Infof("starting server on :%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
