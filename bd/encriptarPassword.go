package bd

import (
	"golang.org/x/crypto/bcrypt"
)

/*EncriptarPassword: permita revisar que la cosita funcione*/
func EncriptarPassword(password string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	return string(bytes), err
}
