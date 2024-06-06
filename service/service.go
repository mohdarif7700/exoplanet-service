package service

import (
	"encoding/json"
	"exoplanet-service/dto"
	"exoplanet-service/validate"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
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

func UpdateExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var exoplanet dto.Exoplanet
	if err := json.NewDecoder(r.Body).Decode(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := validate.ValidateExoplanet(&exoplanet); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, exists := exoplanets[id]
	if !exists {
		http.Error(w, dto.ErrExoplanetNotFound.Error(), http.StatusNotFound)
		return
	}

	exoplanet.ID = id
	exoplanets[id] = exoplanet
	err := json.NewEncoder(w).Encode(exoplanet)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteExoplanetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, exists := exoplanets[id]
	if !exists {
		http.Error(w, dto.ErrExoplanetNotFound.Error(), http.StatusNotFound)
		return
	}

	delete(exoplanets, id)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w)
}

func FuelEstimationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	crewCapacity, err := strconv.Atoi(r.URL.Query().Get("crewCapacity"))
	if err != nil || crewCapacity <= 0 {
		http.Error(w, "Invalid crew capacity", http.StatusBadRequest)
		return
	}

	exoplanet, exists := exoplanets[id]
	if !exists {
		http.Error(w, dto.ErrExoplanetNotFound.Error(), http.StatusNotFound)
		return
	}

	var gravity float64
	switch exoplanet.Type {
	case dto.GasGiant:
		gravity = 0.5 / math.Pow(exoplanet.Radius, 2)
	case dto.Terrestrial:
		gravity = exoplanet.Mass / math.Pow(exoplanet.Radius, 2)
	}

	fuel := float64(exoplanet.Distance) / math.Pow(gravity, 2) * float64(crewCapacity)
	err = json.NewEncoder(w).Encode(map[string]float64{"fuel": fuel})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
