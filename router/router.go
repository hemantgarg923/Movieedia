package router

import (
	"github.com/gorilla/mux"
	"github.com/hemantgarg925/mymongo/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/movies", controller.GetAllMyMovies).Methods("GET")
	router.HandleFunc("/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/movie/{id}", controller.DeleteAMovie).Methods("DELETE")
	router.HandleFunc("/deletemovies", controller.DeleteAllMovies).Methods("DELETE")

	return router
}
