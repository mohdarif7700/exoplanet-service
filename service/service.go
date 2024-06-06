package service

import (
	"encoding/json"
	"exoplanet-service/dto"
	"exoplanet-service/validate"
	"github.com/google/uuid"
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
