package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/KenethSandoval/goTwitter/middlewares"
	"github.com/KenethSandoval/goTwitter/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Handlers set my port, the handler liste to the server  */
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.DBCheck(routers.Registry)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}