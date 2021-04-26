package routers

import (
	"net/http"

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

func EliminarTweet2(rw http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(rw, "Debe espeficiar un id del tweet", http.StatusBadRequest)
		return
	}

	userid := r.URL.Query().Get("userid")
	if len(userid) < 1 {
		http.Error(rw, "Debe especificar un id del usuario dueño del tweet", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(id, userid)

	if err != nil {
		http.Error(rw, "Error mientra se borraba tweet "+err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("content-type", "application/json")
}
