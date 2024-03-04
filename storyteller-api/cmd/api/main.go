package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/handlers"
	"github.com/rep-co/fablescope-backend/storyteller-api/middlewares"
	"github.com/rep-co/fablescope-backend/storyteller-api/storygenerator"
	"github.com/rep-co/fablescope-backend/storyteller-api/util"
)

func main() {
	util.LoadEnv()

	apiKey := os.Getenv("YANDEX_GPT_API_KEY")
	catalogID := os.Getenv("YANDEX_CATALOG_ID")
	prompt := os.Getenv("PROMPT")

	storyGenerator := storygenerator.NewYandexStoryGeneratorWithAPIKey(apiKey, catalogID, prompt)

	router := httprouter.New()

	router.GET("/", handlers.HandleGetIndex)
	router.GET("/form/category", handlers.HandleGetCategory)
	router.POST(
		"/generate/story",
		middlewares.ValidateStoryParameters(
			middlewares.GenerateStory(
				handlers.HandleGetStory,
				storyGenerator,
			),
		),
	)

	log.Println("JSON API server is listening on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
