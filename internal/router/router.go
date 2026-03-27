package router

import (
	"net/http"

	"github.com/gorilla/mux"

	"template-backend/internal/controller"
)


func NewRouter() *mux.Router {
	router := mux.NewRouter()

	register := NewRegister(router)

	templateController := controller.NewTemplateController()

	register.RegisterRoute("/ping", false, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	}).Methods(http.MethodGet, http.MethodPost)

	register.RegisterRoute("/template/list", false, templateController.List).Methods(http.MethodGet)

	register.RegisterRoute("/template/favorite", true, templateController.Favorite).Methods(http.MethodPost)


	prefixRouter := router.PathPrefix("/api/v1").Subrouter()

	prefixRouter.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}).Methods(http.MethodGet)

	return router
}