package main

import (
	"encoding/json"
	"net/http"
)

type Negocio struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Tipo   string `json:"tipo"`
	Ciudad string `json:"ciudad"`
}

func main() {
	http.HandleFunc("/negocios", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET"{
			w.Header().Set("Allow", "GET")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
		negocios := []Negocio{
			{ID: 1, Nombre: "Hotel Oro Verde", Tipo: "Tipo 1", Ciudad: "Cuidad 1"},
			{ID: 2, Nombre: "Restaurante el Rincon", Tipo: "Tipo 2", Ciudad: "Cuidad 2"},
			{ID: 3, Nombre: "Mall del Pacifico", Tipo: "Tipo 3", Ciudad: "Cuidad 3"},
		}
		w.Header().Set("Content-Type", "application/jon")
		json.NewEncoder(w).Encode(negocios)
	})
	http.ListenAndServe(":8080", nil)
}
