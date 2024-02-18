package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/storygenerator"
)

func ValidateStoryParameters(
	next httprouter.Handle,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rawTags := r.URL.Query().Get("tags")
		tags := strings.Split(rawTags, ",")

		//Check if tags and categories are valid
		//Mb it's cool to write validator instead of doing it here or using regex
		ctx := context.WithValue(r.Context(), contextKeyTags, tags)
		next(w, r.WithContext(ctx), ps)
	}

}

func GenerateStory(
	next httprouter.Handle,
	storyGenerator *storygenerator.OpenAIStoryGenerator,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		tags, err := GetTagsKey(r.Context())
		if err != nil {
			log.Fatal(err)
		}

		story, err := storyGenerator.GenerateStory(r.Context(), tags)
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.WithValue(r.Context(), contextKeyStory, story)
		next(w, r.WithContext(ctx), ps)
	}
}
