package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/data"
	"github.com/rep-co/fablescope-backend/storyteller-api/storygenerator"
	"github.com/rep-co/fablescope-backend/storyteller-api/util"
)

func ValidateStoryParameters(
	next httprouter.Handle,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		var rawTags data.TagNames
		err := util.ReadJSON(r, &rawTags)
		if err != nil {
			log.Printf("An errot occured at ValidateStory: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		//Check if tags and categories are valid
		//Mb it's cool to write validator instead of doing it here or using regex
		tags := rawTags.TagNames
		ctx := context.WithValue(r.Context(), contextKeyTags, tags)
		next(w, r.WithContext(ctx), ps)
	}

}

func GenerateStory(
	next httprouter.Handle,
	storyGenerator storygenerator.StoryGenerator,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tags, err := GetTagsKey(r.Context())
		if err != nil {
			log.Printf("An errot occured at GenerateStory: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		story, err := storyGenerator.GenerateStory(r.Context(), tags)
		if err != nil {
			log.Printf("An errot occured at GenerateStory: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyStory, story)
		next(w, r.WithContext(ctx), ps)
	}
}
