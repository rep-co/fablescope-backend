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
	"github.com/rep-co/fablescope-backend/wardrobe-auth/util"
)

func main() {
	ctx := context.Background()
	// TODO: use separate .env
	util.LoadEnv()

	port := os.Getenv("PORT")

	storage, err := database.NewYDBStorage()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer storage.Close(ctx)

	if err := storage.Init(ctx); err != nil {
		log.Fatal(err)
		return
	}

	router := httprouter.New()

	router.POST(
		"/sing-up",
		middlewares.ValidateAccountCredentials(
			ctx,
			middlewares.SingUp(
				ctx,
				handlers.HandleSingUp,
				storage,
			),
		),
	)
	router.POST("/sing-in", handlers.HandleSingIn)
	router.POST("/refresh", handlers.HandleRefresh)

	log.Printf("JSON API server is listening on port: %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
