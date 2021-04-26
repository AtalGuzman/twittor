package routers

import (
	"encoding/json"
	"net/http"
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

func LeoTweets2(rw http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(rw, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	pagina := r.URL.Query().Get("pagina")

	if len(pagina) < 1 {
		http.Error(rw, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	n_pagina, err := strconv.Atoi(pagina)

	if err != nil {
		http.Error(rw, "Error página escogida", http.StatusBadRequest)
		return
	}

	resultados, status := bd.LeoTweets(ID, int64(n_pagina))

	if !status {
		http.Error(rw, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.WriteHeader(http.StatusAccepted)
	json.NewEncoder(rw).Encode(resultados)

}
