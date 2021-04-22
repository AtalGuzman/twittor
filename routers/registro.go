package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/models"
	gb "github.com/gogearbox/gearbox"
)

func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	error := json.NewDecoder(r.Body).Decode(&t) /*body se lee y luego explota*/

	if error != nil {
		http.Error(w, "ERROR DECODIFICACION JSON\n"+error.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "EMAIL VACÍO\n", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "PASSWORD DEMASIADO CORTO. DEBE SER AL MENOS DE 6 CARACTÉRES\n", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "EL EMAIL YA HA SIDO USADO", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "ERROR AL GUARDAR EN BD\n"+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "NO SE HA LOGRADO INSERTAR EL REGISTRO DEL USUARIO", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func Registro2(ctx gb.Context) {
	var t models.Usuario
	err := ctx.ParseBody(&t)

	if err != nil {
		ctx.Status(400)
		ctx.SendString("ERROR DECODIFICACION JSON\n" + err.Error())
		return
	}

	if len(t.Email) == 0 {
		ctx.Status(400)
		ctx.SendString("Email vacío")
		return
	}

	if len(t.Password) < 6 {
		ctx.Status(400)
		ctx.SendString("Password demasiado corto. Debe ser al menos de 6 caractéres")
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		ctx.Status(400)
		ctx.SendString("El email ya ha sido usado")
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		ctx.Status(500)
		ctx.SendString("Error al guardar en la bd")
		return
	}

	if !status {
		ctx.Status(500)
		ctx.SendString("Error al insertar en la bd")
		return
	}

	ctx.Status(http.StatusCreated)
}
