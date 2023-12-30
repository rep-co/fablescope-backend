package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/handlers"
)

func main() {
	router := httprouter.New()

	router.GET("/", handlers.HandleGetIndex)
	router.GET("/form/category", handlers.HandleGetCategory)

	log.Println("JSON API server is listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
