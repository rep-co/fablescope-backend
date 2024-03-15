package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/handlers"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/util"
)

func main() {
	// TODO: use separate .env
	util.LoadEnv()

	port := os.Getenv("PORT")

	router := httprouter.New()

	router.POST("/sing-up", handlers.HandleSingUp)
	router.POST("/sing-in", handlers.HandleSingIn)
	router.POST("/refresh", handlers.HandleRefresh)

	log.Printf("JSON API server is listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
