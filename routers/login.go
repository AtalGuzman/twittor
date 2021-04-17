package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/jwt"
	"github.com/AtalGuzman/twittor/models"
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
