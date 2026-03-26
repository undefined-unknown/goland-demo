package router

import (
	"net/http"

	"github.com/gorilla/mux"
)


func NewRouter() *mux.Router {
	router := mux.NewRouter()

	register := NewRegister(router)

	register.RegisterRoute("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet, http.MethodPost)

	prefixRouter := router.PathPrefix("/api/v1").Subrouter()

	prefixRouter.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}).Methods(http.MethodGet)

	return router
}