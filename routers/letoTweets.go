package routers

import (
	"strconv"

	"github.com/AtalGuzman/twittor/bd"
	"github.com/gogearbox/gearbox"
)

func LeoTweets(gb gearbox.Context) {

	ID := gb.Query("id")
	if len(ID) < 1 {
		gb.Status(400)
		gb.SendString("Debe enviar el parámetro ID")
		return
	}
	pagina := gb.Query("pagina")
	if len(pagina) < 1 {
		gb.Status(400)
		gb.SendString("Debe enviar el parámetro de número de página")
		return
	}
	n_pagina, err := strconv.Atoi(pagina)
	if err != nil {
		gb.Status(500)
		gb.SendString("Error la página escogida")
		return
	}
	resultados, status := bd.LeoTweets(ID, int64(n_pagina))
	if !status {
		gb.Status(500)
		gb.SendString("Error al leer los tweets")
		return
	}

	gb.Set("content-type", "application/json")
	gb.Status(gearbox.StatusCreated)
	gb.SendJSON(resultados)

}
