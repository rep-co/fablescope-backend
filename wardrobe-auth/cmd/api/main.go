package main

import (
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
	// TODO: use separate .env
	util.LoadEnv()

	port := os.Getenv("PORT")

	storage, err := database.NewYDBStorage()
	if err != nil {
		log.Fatal(err)
		return
	}
	//if err := storage.Init(); err != nil {
	//	log.Fatal(err)
	//	return
	//}

	router := httprouter.New()

	router.POST(
		"/sing-up",
		middlewares.ValidateAccountCredentials(
			middlewares.SingUp(
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
