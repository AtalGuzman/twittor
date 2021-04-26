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
	router.HandleFunc("/tweet", middleW.ChequeoBd(middleW.ValidoJWT(routers.GraboTweet2))).Methods("POST")
	router.HandleFunc("/leotweet", middleW.ChequeoBd(middleW.ValidoJWT(routers.LeoTweets2))).Methods("GET")
	router.HandleFunc("/borrotweet", middleW.ChequeoBd(middleW.ValidoJWT(routers.EliminarTweet2))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middleW.ChequeoBd(middleW.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middleW.ChequeoBd(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middleW.ChequeoBd(middleW.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middleW.ChequeoBd(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middleW.ChequeoBd(middleW.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middleW.ChequeoBd(middleW.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middleW.ChequeoBd(middleW.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middleW.ChequeoBd(middleW.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middleW.ChequeoBd(middleW.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

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
	gb.Post("/tweet", middleW.ChequeoBd2, middleW.ValidoJWT2, routers.GraboTweet)
	gb.Get("/leotweet", middleW.ChequeoBd2, middleW.ValidoJWT2, routers.LeoTweets)
	gb.Get("/borrotweet", middleW.ChequeoBd2, middleW.ValidoJWT2, routers.EliminarTweet)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	gb.Start(":" + PORT)

}
