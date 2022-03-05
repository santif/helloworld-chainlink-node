package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func NewServer() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/price", a.fetchPrice).Methods(http.MethodPost)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchPrice(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching price")
	w.Header().Set("Content-Type", "application/json")

	// var req Request
	var resp PriceResponse

	// defer r.Body.Close()
	// err := json.NewDecoder(r.Body).Decode(&req)

	// if err != nil {
	// 	log.Printf("Error decoding request: %s\n", err)
	// 	w.WriteHeader(http.StatusBadRequest)

	// 	resp = PriceResponse{
	// 		Error: "Error decoding request",
	// 	}
	// 	json.NewEncoder(w).Encode(resp)

	// 	return
	// }

	resp.Data.Price = 123
	resp.Data.Timestamp = 123456789

	json.NewEncoder(w).Encode(resp)
}
