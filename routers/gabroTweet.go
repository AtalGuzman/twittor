package routers

import (
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
