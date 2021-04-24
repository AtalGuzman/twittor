package routers

import (
	"github.com/AtalGuzman/twittor/bd"
	"github.com/gogearbox/gearbox"
)

func EliminarTweet(gb gearbox.Context) {
	id := gb.Query("id")
	if len(id) < 1 {
		gb.Status(gearbox.StatusBadRequest)
		gb.SendString("Debe espeficiar un id del tweet")
		return
	}

	userid := gb.Query("userid")
	if len(userid) < 1 {
		gb.Status(gearbox.StatusBadRequest)
		gb.SendString("Debe especificar un id del usuario dueño del tweet")
		return
	}

	err := bd.BorroTweet(id, userid)

	if err != nil {
		gb.Status(gearbox.StatusInternalServerError)
		gb.SendString("Error mientra se borraba tweet " + err.Error())
		return

	}

	gb.Status(gearbox.StatusOK)
	gb.Set("content-type", "application/json")
}
