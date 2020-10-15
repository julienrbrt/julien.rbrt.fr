package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

// NewRouter creates a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// default route
	router.HandleFunc("/", handleHome)
	// route public folder
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	// not found route defaults to home
	router.NotFoundHandler = http.HandlerFunc(handleHome)
	// add logger middleware
	router.Use(Logger)

	return router
}
