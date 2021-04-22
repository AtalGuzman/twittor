package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/models"
	"github.com/gogearbox/gearbox"
)

func ModificarPerfil(rw http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(rw, "Datos incorrectos "+err.Error(), 400)
		return
	}

	status, err := bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	if !status {
		http.Error(rw, "No se modificó el registro del usuario", 400)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func ModificarPerfil2(ctx gearbox.Context) {
	var t models.Usuario
	err := ctx.ParseBody(&t)
	fmt.Print("*")
	fmt.Print(t)
	if err != nil {
		ctx.Status(400)
		ctx.SendString("Datos incorrectos " + err.Error())
		return
	}

	status, err := bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		ctx.Status(400)
		ctx.SendString(err.Error())
		return
	}

	if !status {
		ctx.Status(400)
		ctx.SendString("No se modificó el registro de usuario")
		return
	}

	ctx.Status(gearbox.StatusCreated)
}
