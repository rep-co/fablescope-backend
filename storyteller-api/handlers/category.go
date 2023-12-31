package handlers

import (
	"fmt"
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
	categories := data.Categories

	err := util.WriteJSON(w, http.StatusOK, categories)
	if err != nil {
		fmt.Printf("An error occure at HandleGetCategory: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
