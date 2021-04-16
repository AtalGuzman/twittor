package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/AtalGuzman/twittor/middleW"
	"github.com/AtalGuzman/twittor/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores: Seteo de puerto, handler e inicio del servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleW.ChequeoBd(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
