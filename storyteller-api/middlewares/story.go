package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ValidateStoryParameters(
	next httprouter.Handle,
	r *http.Request,
	ps httprouter.Params,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//TODO: check if tags and categories are valid

		next(w, r, ps)
	}
}

func MakeStoryPromt(
	next httprouter.Handle,
	r *http.Request,
	ps httprouter.Params,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//TODO: construct promt

		next(w, r, ps)
	}
}
