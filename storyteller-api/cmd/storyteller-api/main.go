package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rep-co/fablescope-backend/storyteller-api/internal/config"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
