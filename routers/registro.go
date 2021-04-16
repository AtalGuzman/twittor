package routers

import (
	"encoding/json"
	"net/http"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/models"
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
	}

	w.WriteHeader(http.StatusCreated)
}
