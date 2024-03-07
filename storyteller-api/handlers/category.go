package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/storyteller-api/data"
	"github.com/rep-co/fablescope-backend/storyteller-api/util"
)

func HandleGetCategory(
	w http.ResponseWriter,
	_ *http.Request,
	_ httprouter.Params,
) {
	categories := data.CategoriesDefined

	err := util.WriteJSON(w, http.StatusOK, categories)
	if err != nil {
		log.Printf("An error occure at HandleGetCategory: %v.", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
