package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/handlers"
	"github.com/rep-co/fablescope-backend/storyteller-api/iamgenerator"
	"github.com/rep-co/fablescope-backend/storyteller-api/middlewares"
	"github.com/rep-co/fablescope-backend/storyteller-api/storygenerator"
	"github.com/rep-co/fablescope-backend/storyteller-api/util"
)

func main() {
	util.LoadEnv()

	port := os.Getenv("PORT")
	catalogID := os.Getenv("YANDEX_CATALOG_ID")
	prompt := os.Getenv("PROMPT")

	iamTokenGenerator := iamgenerator.NewIAMTokenServerless()
	iamToken, err := iamTokenGenerator.GenerateToken()
	if err != nil {
		log.Printf("Can't get token %s", err)
		return
	}

	storyGenerator := storygenerator.NewYandexStoryGeneratorWithIAMToken(iamToken, catalogID, prompt)

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

	log.Printf("JSON API server is listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
