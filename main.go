package main

import (
	delivery "github.com/ivchip/go-meli-filter-ip/interface/http"
	"github.com/ivchip/go-meli-filter-ip/repository/postgres"
	"github.com/ivchip/go-meli-filter-ip/usecases"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	postgres.New()
	usecases.New()
	err := usecases.Migrate()
	if err != nil {
		log.Fatalf("Migrate: %v", err)
	}
	delivery.StartWebServer(port)
}
