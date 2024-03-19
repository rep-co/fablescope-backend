package handlers

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/util"
)

func HandleSingUp(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params,
) {
	util.WriteJSON(w, http.StatusOK, "success")
}

func HandleSingIn(
	w http.ResponseWriter,
	r *http.Request,
	_ httprouter.Params,
) {

}

func HandleRefresh(
	w http.ResponseWriter,
	_ *http.Request,
	_ httprouter.Params,
) {
	err := util.WriteJSON(w, http.StatusOK, "AMOGUS Refresh")
	if err != nil {
		log.Printf("An error occure at HandleGetCategory: %v.", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
