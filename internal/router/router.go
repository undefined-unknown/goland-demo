package router

import (
	"net/http"

	"github.com/gorilla/mux"
)


func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	prefixRouter := router.PathPrefix("/api/v1").Subrouter()

	prefixRouter.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}).Methods(http.MethodGet)

	return router
}