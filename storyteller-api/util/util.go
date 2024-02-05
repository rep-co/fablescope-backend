package util

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// TODO: Dotenv does not provide functionality for auto searching in all
// parental directories. This code is just from stackOverFlow, better to rewrite
func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + "fablescope-backend" + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
