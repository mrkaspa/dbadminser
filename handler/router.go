package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type fakeRouter struct {
	mux *mux.Router
}

// NewRouter inits a router
func NewRouter() http.Handler {
	n := negroni.Classic()
	router := &fakeRouter{mux: mux.NewRouter()}
	connHandler := connHandler{}
	router.mux.HandleFunc("/conns", connHandler.storeHandler).Methods("POST")
	router.mux.HandleFunc("/conns", connHandler.storeHandler).Methods("GET")
	router.mux.HandleFunc("/conns/{id}", connHandler.storeHandler).Methods("DELETE")
	n.Use(router)
	return n
}

func (router *fakeRouter) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	router.mux.ServeHTTP(w, r)
	next(w, r)
}
