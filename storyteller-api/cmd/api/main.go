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

	key := os.Getenv("CHAT_GPT_KEY")
	prompt := os.Getenv("PROMPT")

	var storyGenerator = storygenerator.NewOpenAIStoryGenerator(key, prompt)

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
