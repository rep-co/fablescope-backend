package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/database"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/handlers"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/middlewares"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/services"
	"github.com/rep-co/fablescope-backend/wardrobe-auth/util"
)

func main() {
	util.LoadEnv()

	port := os.Getenv("PORT")
	ydbConnString := os.Getenv("YDB_CONN_STRING")

	// TODO: get token by api call, will do it later
	token := os.Getenv("TOKEN")

	ctx := context.Background()

	accountStorage, err := database.NewYDBStorage(ctx, ydbConnString, token)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer accountStorage.Close(ctx)

	if err := accountStorage.Init(ctx); err != nil {
		log.Fatal(err)
		return
	}

	accountService := services.NewAccountService(accountStorage)

	router := httprouter.New()

	router.POST(
		"/sing-up",
		middlewares.ValidateAccountCredentials(
			ctx,
			middlewares.SingUp(
				ctx,
				handlers.HandleSingUp,
				accountService,
			),
		),
	)
	router.POST(
		"/sing-in",
		middlewares.ValidateAccountCredentials(
			ctx,
			middlewares.SingIn(
				ctx,
				handlers.HandleSingIn,
				accountStorage,
			),
		),
	)
	router.POST("/refresh", handlers.HandleRefresh)

	log.Printf("JSON API server is listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
