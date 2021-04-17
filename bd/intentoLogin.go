package bd

import (
	"github.com/AtalGuzman/twittor/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin: valida password ingresada*/
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usuario, encontrado, _ := ChequeoYaExisteUsuario(email)
	if !encontrado {
		return usuario, false
	}

	passwordBD := []byte(usuario.Password)
	passwordByte := []byte(password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordByte)

	if err != nil {
		return usuario, false
	}
	return usuario, true
}
