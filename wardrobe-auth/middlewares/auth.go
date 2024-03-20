package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/database"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/util"
)

func ValidateAccountCredentials(
	ctx context.Context,
	next httprouter.Handle,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		request := data.AccountRequest{}
		err := util.ReadJSON(r, &request)
		if err != nil {
			log.Printf("An error occure at ValidateAccountCredentials: %v.", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyAccountRequest, &request)
		next(w, r.WithContext(ctx), ps)
	}
}

func SingUp(
	ctx context.Context,
	next httprouter.Handle,
	s database.Storage,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		request, err := GetAccountRequestKey(r.Context())
		if err != nil {
			log.Printf("An error occure at SingUp: %v.", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		account := data.NewAccount(request)
		if err := s.CreateAccount(ctx, account); err != nil {
			log.Printf("An error occure at SingUp: %v.", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// TODO:
		// 1. Generate JWT
		// 2. Add JWT to JWT db
		// 3. Add JWT to context

		next(w, r, ps)
	}
}

func SingIn(
	ctx context.Context,
	next httprouter.Handle,
	s database.Storage,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		_, err := GetAccountRequestKey(r.Context())
		if err != nil {
			log.Printf("An error occure at SingIn: %v.", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		log.Println("success")
		next(w, r, ps)
	}
}

func Refresh(
	ctx context.Context,
	next httprouter.Handle,
	s database.Storage,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		next(w, r, ps)
	}
}
