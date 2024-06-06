package service

import (
	"encoding/json"
	"exoplanet-service/dto"
	"exoplanet-service/validate"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

var exoplanets = make(map[string]dto.Exoplanet)

func AddExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	var exoplanet dto.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the input
	if err := validate.ValidateExoplanet(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	exoplanet.ID = uuid.New().String()
	exoplanets[exoplanet.ID] = exoplanet
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ListExoplanetsHandler(w http.ResponseWriter, r *http.Request) {
	var list []dto.Exoplanet
	for _, exoplanet := range exoplanets {
		list = append(list, exoplanet)
	}

	err := json.NewEncoder(w).Encode(list)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	exoplanet, exists := exoplanets[id]
	if !exists {
		http.Error(w, dto.ErrExoplanetNotFound.Error(), http.StatusNotFound)
		return
	}

	err := json.NewEncoder(w).Encode(exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
