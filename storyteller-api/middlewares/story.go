package middlewares

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func ValidateStoryParameters(
	next httprouter.Handle,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//TODO: check if tags and categories are valid
		//TODO: add proper context key type

		tags := make([]int, 10)
		ctx := context.WithValue(r.Context(), "tags", tags)
		next(w, r.WithContext(ctx), ps)
	}
}

func MakeStoryPrompt(
	next httprouter.Handle,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		//TODO: construct prompt
		//TODO: add proper context key type
		prompt := "what is amogus?"
		ctx := context.WithValue(r.Context(), "promt", prompt)
		next(w, r.WithContext(ctx), ps)
	}
}
