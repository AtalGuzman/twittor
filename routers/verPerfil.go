package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/gogearbox/gearbox"
)

func VerPerfil(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	usuario, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(rw, "Ocurrió error al buscar el registro "+err.Error(), 400)
		return
	}

	rw.Header().Set("content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(usuario)
}

func VerPerfil2(ctx gearbox.Context) {
	fmt.Print("*")
	ID := ctx.Query("id")
	if len(ID) < 1 {
		ctx.Status(gearbox.StatusBadRequest)
		ctx.SendString("Debe enviar el parámetro ID")
		return
	}

	usuario, err := bd.BuscoPerfil(ID)
	if err != nil {
		ctx.Status(gearbox.StatusBadRequest)
		ctx.SendString("Ocurrió error al buscar el registro " + err.Error())
		return
	}

	ctx.Status(http.StatusCreated)
	ctx.Set("Content-type", "application/json")
	ctx.SendJSON(usuario)

}
