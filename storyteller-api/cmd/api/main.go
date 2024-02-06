package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/handlers"
	"github.com/rep-co/fablescope-backend/storyteller-api/middlewares"
	"github.com/rep-co/fablescope-backend/storyteller-api/util"
)

func main() {
	util.LoadEnv()

	router := httprouter.New()

	router.GET("/", handlers.HandleGetIndex)
	router.GET("/form/category", handlers.HandleGetCategory)
	router.GET(
		"/generate/story",
		middlewares.ValidateStoryParameters(
			middlewares.MakeStoryPrompt(
				handlers.HandleGetStory,
			),
		),
	)

	log.Println("JSON API server is listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
