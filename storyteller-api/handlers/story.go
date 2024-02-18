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
		log.Printf("An errot occured at HandleGetStory: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	//TODO: why there is an error when using err :=?????
	//Потом зарефакторю, я обещаю
	err2 := util.WriteJSON(w, http.StatusOK, story)
	if err2 != nil {
		log.Printf("An errot occured at HandleGetStory: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
