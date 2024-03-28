package middlewares

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/data"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/database"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/services"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/util"
	"golang.org/x/crypto/bcrypt"
)

const (
	ydbRequestTTL = time.Second * 5
	tokenTTL      = time.Minute * 15
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

		ctxRequestValue := context.WithValue(
			r.Context(),
			contextKeyAccountRequest,
			&request,
		)
		next(w, r.WithContext(ctxRequestValue), ps)
	}
}

func SingUp(
	ctx context.Context,
	next httprouter.Handle,
	accountService *services.AccountService,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		request, err := GetAccountRequestKey(r.Context())
		if err != nil {
			log.Printf("An error occure at SingUp: %v.", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		err = accountService.CreateNewAccount(ctx, request)
		if err != nil {
			log.Printf("An error occure at SingUp: %v.", err)
			switch {
			case errors.Is(err, &database.RequestTimeoutError):
				http.Error(w, http.StatusText(http.StatusGatewayTimeout), http.StatusGatewayTimeout)
			case errors.Is(err, &database.TransactionError):
				http.Error(w, "An account with the given email already exists", http.StatusConflict)
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
	s database.AccountStorage,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		request, err := GetAccountRequestKey(r.Context())
		if err != nil {
			log.Printf("An error occure at SingIn: %v.", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctxTxDeadline, cancel := context.WithDeadline(
			ctx,
			time.Now().Add(ydbRequestTTL),
		)
		defer cancel()

		// TODO: Mb it's a good idea to create some sort of a service
		// and then refactor this, moving into it's dedicated service
		account, err := s.GetAccount(ctxTxDeadline, request.Email)
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

		// TODO: JWT cooking should also be it's own service
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(tokenTTL)},
			Subject:   account.ID.String(),
		})
		tokenStr, err := token.SignedString([]byte("amogus"))
		if err != nil {
			log.Printf("An error occure at SingIn: %v.", err)
			return
		}
		refreshTokenStr := "amogus" //Cook refresh token

		ctxRequestValue := context.WithValue(
			r.Context(),
			contextKeyTokens,
			&data.Tokens{
				JWTToken:     tokenStr,
				RefreshToken: refreshTokenStr,
			},
		)
		next(w, r.WithContext(ctxRequestValue), ps)
	}
}

func Refresh(
	ctx context.Context,
	next httprouter.Handle,
	s database.AccountStorage,
) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		next(w, r, ps)
	}
}
