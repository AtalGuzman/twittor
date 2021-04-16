package main

import (
	"log"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/AtalGuzman/twittor/handlers"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la BD :C")
		return
	}

	handlers.Manejadores()
}
