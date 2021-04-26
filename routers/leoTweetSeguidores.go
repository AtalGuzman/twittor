package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AtalGuzman/twittor/bd"
)

func LeoTweetsSeguidores(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe agregar la pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe agregar la pagina", http.StatusBadRequest)
		return
	}

	respuest, status := bd.LeoTweetsSeguidores(IDUsuario, pagina)

	if !status {
		http.Error(w, "No se pudieron obtener los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuest)
}
