package routers

import (
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/models"
)

func AltaRelacion(rw http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(rw, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(rw, "Error al insertar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(rw, "No se ha logrado hacer la inserción", http.StatusBadRequest)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
