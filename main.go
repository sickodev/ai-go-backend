package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create Rouuter
	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/healthz", HealthHandler)
	r.HandleFunc("/error", ErrorHandler)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal("DID NOT LISTEN")
	} else {
		log.Default()
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page?")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
	}{"API Running..."}

	json, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	data := struct {
		Message string
	}{"API Failed..."}

	json, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(json)
}
