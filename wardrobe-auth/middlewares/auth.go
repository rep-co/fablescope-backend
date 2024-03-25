package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/database"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/util"
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordHashingCost = bcrypt.DefaultCost // 10
	ydbRequestTTL       = time.Second * 5
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

		ctx, cancel := context.WithDeadline(ctx, time.Now().Add(ydbRequestTTL))
		defer cancel()

		// TODO: Mb it's a good idea to create some sort of a service
		// and then refactor this, moving into it's dedicated service
		account := data.NewAccount(request)
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(account.Password),
			passwordHashingCost,
		)
		if err != nil {
			log.Printf("An error occure at SingUp: %v.", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		account.Password = string(hashedPassword)

		if err := s.CreateAccount(ctx, account); err != nil {
			log.Printf("An error occure at SingUp: %v.", err)
			switch {
			case errors.Is(err, &database.RequestTimeoutError):
				http.Error(w, http.StatusText(http.StatusGatewayTimeout), http.StatusGatewayTimeout)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		next(w, r, ps)
	}
}

func SingIn(
	ctx context.Context,
	next httprouter.Handle,
	s database.Storage,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		request, err := GetAccountRequestKey(r.Context())
		if err != nil {
			log.Printf("An error occure at SingIn: %v.", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctx, cancel := context.WithDeadline(ctx, time.Now().Add(ydbRequestTTL))
		defer cancel()

		// TODO: Mb it's a good idea to create some sort of a service
		// and then refactor this, moving into it's dedicated service
		account, err := s.GetAccount(ctx, request.Email)
		if err != nil {
			log.Printf("An error occure at SingIn: %v.", err)
			switch {
			case errors.Is(err, &database.RequestTimeoutError):
				http.Error(w, http.StatusText(http.StatusGatewayTimeout), http.StatusGatewayTimeout)
			case errors.Is(err, &database.NoResultError):
				http.Error(w, "Wrong Email or Password", http.StatusUnauthorized)
			default:
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		if err = bcrypt.CompareHashAndPassword(
			[]byte(account.Password),
			[]byte(request.Password),
		); err != nil {
			log.Printf("An error occure at SingIn: %v.", err)
			http.Error(w, "Wrong Email or Password", http.StatusUnauthorized)
			return
		}

		// TODO:
		// 1. Generate JWT
		// 2. Add JWT to JWT db
		// 3. Add JWT to context
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
