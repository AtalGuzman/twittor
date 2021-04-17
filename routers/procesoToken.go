package routers

import (
	"errors"
	"strings"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/models"
	jwt "github.com/dgrijalva/jwt-go"
)

var Email string
var IDUsuario string

func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MasterDelDesarrolloGrupodeFB")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err != nil {
		return claims, false, string(""), err
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

	if !encontrado {
		return claims, false, string(""), errors.New("no se encontro usuario")
	}

	Email = claims.Email
	IDUsuario = claims.ID.Hex()

	return claims, encontrado, IDUsuario, nil

}
