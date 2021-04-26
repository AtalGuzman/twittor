package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/models"
)

func SubirAvatar(rw http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("avatar")
	fmt.Println(handler.Filename)
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/avatars/" + IDUsuario + "." + extension

	if err != nil {
		http.Error(rw, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(rw, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(rw, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario

	usuario.Avatar = IDUsuario + "." + extension
	status, err := bd.ModificoRegistro(usuario, IDUsuario)

	if err != nil || !status {
		http.Error(rw, "Error al grabar el avatar en la bd!"+err.Error(), http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusCreated)
}
