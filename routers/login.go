package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/jwt"
	"github.com/AtalGuzman/twittor/models"
	"github.com/gogearbox/gearbox"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	var u models.Usuario

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Usuario o contraseña equivocada", 400)
		return
	}

	if len(u.Email) == 0 {
		http.Error(w, "No se ingresó email", 400)
		return
	}

	documento, existe := bd.IntentoLogin(u.Email, u.Password)
	if !existe {
		http.Error(w, "Usuario o contraseña equivocada", 400)
	}
	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Hubo un error "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{Token: jwtKey}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{Name: "token", Value: jwtKey, Expires: expirationTime})
}

func Login2(ctx gearbox.Context) {

	ctx.Set("content-type", "application/json")
	var u models.Usuario

	err := ctx.ParseBody(&u)

	if err != nil {
		ctx.Status(400)
		ctx.SendString("Usuario o contraseña equivocada")
		return
	}

	if len(u.Email) == 0 {
		ctx.Status(400)
		ctx.SendString("No se ingresó el mail")
		return
	}

	documento, existe := bd.IntentoLogin(u.Email, u.Password)
	if !existe {
		ctx.Status(400)
		ctx.SendString("Usuario o contraseña equivocada")
		return
	}
	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		ctx.Status(400)
		ctx.SendString("Hubo un error " + err.Error())
		return
	}

	resp := models.RespuestaLogin{Token: jwtKey}

	ctx.Status(http.StatusCreated)
	ctx.Set("Content-type", "application/json")
	ctx.SendJSON(resp)

}
