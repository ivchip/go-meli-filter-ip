package main

import (
	delivery "github.com/ivchip/go-meli-filter-ip/interface/http"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	delivery.StartWebServer(port)
}
