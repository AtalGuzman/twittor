package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/models"
	"github.com/gogearbox/gearbox"
)

func GraboTweet(gb gearbox.Context) {
	var mensaje models.Tweet
	err := gb.ParseBody(&mensaje)
	if err != nil {
		gb.Status(400)
		gb.SendString("Error en la decodificacion " + err.Error())
		return
	}

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		gb.Status(400)
		gb.SendString(err.Error())
		return
	}
	if !status {
		gb.Status(400)
		gb.SendString("No se pudo insertar a la bd")
		return
	}
}

func GraboTweet2(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "No se pudo decodificar el mensaje", http.StatusBadRequest)
		return
	}

	registro := models.GraboTweet{
		UserId:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se pudo insertar la cosita", http.StatusBadRequest)
		return
	}
}
