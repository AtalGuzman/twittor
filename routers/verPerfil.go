package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
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
