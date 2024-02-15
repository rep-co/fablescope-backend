package handlers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/util"
)

func HandleGetStory(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params,
) {
	fmt.Println("AMOGUS")
	story := r.Context().Value("story").(string)
	fmt.Println(story)

	err := util.WriteJSON(w, http.StatusOK, story)
	if err != nil {
		fmt.Printf("An error occure at HandleGetCategory: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
