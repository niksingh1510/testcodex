package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

// placeholder memory store using simple structs
type Tenant struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Instance struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	TenantID int    `json:"tenant_id"`
}

type Dataplane struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Region   string `json:"region"`
	TenantID int    `json:"tenant_id"`
}

var (
	tenants     = make([]Tenant, 0)
	instances   = make([]Instance, 0)
	dataplanes  = make([]Dataplane, 0)
	mu          sync.Mutex
	tenantID    = 1
	instanceID  = 1
	dataplaneID = 1
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/signup", signupHandler)
	mux.HandleFunc("/license", licenseHandler)
	mux.HandleFunc("/tenants", tenantsHandler)
	mux.HandleFunc("/instances", instancesHandler)
	mux.HandleFunc("/dataplanes", dataplanesHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// OIDC registration would happen here
	w.WriteHeader(http.StatusCreated)
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
		var in Instance
		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		in.ID = instanceID
		instanceID++
		instances = append(instances, in)
		mu.Unlock()
		w.WriteHeader(http.StatusAccepted)
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(instances)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func dataplanesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var dp Dataplane
		if err := json.NewDecoder(r.Body).Decode(&dp); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		dp.ID = dataplaneID
		dataplaneID++
		dataplanes = append(dataplanes, dp)
		mu.Unlock()
		w.WriteHeader(http.StatusAccepted)
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(dataplanes)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func tenantsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var t Tenant
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		mu.Lock()
		t.ID = tenantID
		tenantID++
		tenants = append(tenants, t)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
	case http.MethodGet:
		mu.Lock()
		defer mu.Unlock()
		json.NewEncoder(w).Encode(tenants)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// Example of OIDC verifier setup
// oidcVerifier would initialize an OIDC provider and verifier using authentik.
// In this skeleton implementation, it is left as a placeholder.
func oidcVerifier() {
	// TODO: integrate authentik OIDC verifier
}
