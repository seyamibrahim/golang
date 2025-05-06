package router

import (
	"github.com/gorilla/mux"
	controller "github.com/seyamibrahim/mongoapi/controllers"
)

func Router() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/api/movies",controller.GetMyAllMovies).Methods("GET")
	r.HandleFunc("/api/movie",controller.CreateMovie).Methods("POST")
	r.HandleFunc("/api/movie/{id}",controller.MarkAsWatched).Methods("PUT")
	r.HandleFunc("/api/movie/{id}",controller.DeleteOneMovie).Methods("DELETE")
	r.HandleFunc("/api/deleteall",controller.DeleteAllMovie).Methods("DELETE")

	return r
}