package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
}

func newAPIServer(listeningAddr string) *APIServer {
	return &APIServer{
		listenAddr: listeningAddr,
	}
}

func (s *APIServer) run() {
	router := mux.NewRouter()

	router.HandleFunc("/api/v1/create", makeHTTPHandleFunc(s.handleAccount))

	log.Println("JSON API server running on port: ", s.listenAddr)
	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {

	if r.Method == "POST" {
		createReq := new(CreateAccountRequest)
		if err := json.NewDecoder(r.Body).Decode(createReq); err != nil {
			return err
		}

		account := NewAccount(createReq.FirstName, createReq.LastName)

		log.Println("account created: ", account)

		return WriteJSON(w, http.StatusOK, account)

	}

	return fmt.Errorf("method not allowed %s", r.Method)

}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiError struct {
	Error string `json:"error"`
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}
