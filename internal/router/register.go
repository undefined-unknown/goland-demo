package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Register struct {
	router *mux.Router
}

func NewRegister(router *mux.Router) *Register {
	return &Register { router }
}

func (r *Register) RegisterRoute(path string, handler http.HandlerFunc) * RouteItem {
	route := r.router.HandleFunc(path, handler)

	return &RouteItem{ route }
}

type RouteItem struct {
	route *mux.Route
}

func (ri *RouteItem) Methods(methods ...string) {
	ri.route.Methods(methods...)
}