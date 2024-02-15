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

		ctx := context.WithValue(r.Context(), contextKeyTags, tags)
		next(w, r.WithContext(ctx), ps)
	}

}

func GenerateStory(
	next httprouter.Handle,
	storyGenerator *storygenerator.OpenAIStoryGenerator,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//TODO: unsafe code, need to be fixed
		tags := r.Context().Value(contextKeyTags).([]string)
		story, err := storyGenerator.GenerateStory(r.Context(), tags)
		if err != nil {
			log.Fatal(err)
		}

		//TODO: re-think how to properly use keys
		ctx := context.WithValue(r.Context(), "story", story)
		next(w, r.WithContext(ctx), ps)
	}
}
