package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/AtalGuzman/twittor/middleW"
	"github.com/AtalGuzman/twittor/routers"
	"github.com/gogearbox/gearbox"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores: Seteo de puerto, handler e inicio del servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middleW.ChequeoBd(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middleW.ChequeoBd(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middleW.ChequeoBd(middleW.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middleW.ChequeoBd(middleW.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}

func Manejadores2() {
	gb := gearbox.New()

	gb.Post("/registro", middleW.ChequeoBd2, routers.Registro2)
	gb.Post("/login", middleW.ChequeoBd2, routers.Login2)
	gb.Get("/verperfil", middleW.ChequeoBd2, middleW.ValidoJWT2, routers.VerPerfil2)
	gb.Put("/modificarPerfil", middleW.ChequeoBd2, middleW.ValidoJWT2, routers.ModificarPerfil2)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	gb.Start(":" + PORT)

}
