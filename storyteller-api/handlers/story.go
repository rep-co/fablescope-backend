package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/middlewares"
	"github.com/rep-co/fablescope-backend/storyteller-api/util"
)

func HandleGetStory(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params,
) {
	story, err := middlewares.GetStoryKey(r.Context())
	if err != nil {
		log.Printf("An error occured at HandleGetStory: %v.", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	errJSON := util.WriteJSON(w, http.StatusOK, story)
	if errJSON != nil {
		log.Printf("An error occured at HandleGetStory: %v.", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
