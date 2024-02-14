package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func ValidateStoryParameters(
	next httprouter.Handle,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		rawTags := r.URL.Query().Get("tags")
		tags := strings.Split(rawTags, ",")

		//Check if tags and categories are valid
		//TODO: add proper context key type

		ctx := context.WithValue(r.Context(), "tags", tags)
		next(w, r.WithContext(ctx), ps)
	}

}
