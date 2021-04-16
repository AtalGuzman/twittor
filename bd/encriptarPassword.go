package bd

import (
	"golang.org/x/crypto/bcrypt"
)

func EncriptarPassword(password string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), costo)
	return string(bytes), err
}
