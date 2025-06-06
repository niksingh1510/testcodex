package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// placeholder memory store
var instances = make([]string, 0)
var dataplanes = make([]string, 0)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/license", licenseHandler)
	mux.HandleFunc("/instances", instancesHandler)
	mux.HandleFunc("/dataplanes", dataplanesHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func licenseHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// validate license key (placeholder)
	w.WriteHeader(http.StatusOK)
}

func instancesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// create instance placeholder
		instances = append(instances, "instance")
		w.WriteHeader(http.StatusAccepted)
	case http.MethodGet:
		json.NewEncoder(w).Encode(instances)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func dataplanesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// create dataplane placeholder
		dataplanes = append(dataplanes, "dataplane")
		w.WriteHeader(http.StatusAccepted)
	case http.MethodGet:
		json.NewEncoder(w).Encode(dataplanes)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
