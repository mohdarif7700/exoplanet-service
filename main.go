package main

import (
	"exoplanet-service/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/exoplanet/add", service.AddExoplanetHandler).Methods("POST")
	r.HandleFunc("/exoplanet/list", service.ListExoplanetsHandler).Methods("GET")
	r.HandleFunc("/exoplanet/{id}", service.GetExoplanetHandler).Methods("GET")
	r.HandleFunc("/exoplanet/update/{id}", service.UpdateExoplanetHandler).Methods("PUT")
	r.HandleFunc("/exoplanet/remove/{id}", service.DeleteExoplanetHandler).Methods("DELETE")
	r.HandleFunc("/exoplanet/fuel/{id}", service.FuelEstimationHandler).Methods("GET")

	log.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
