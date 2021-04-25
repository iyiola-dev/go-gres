package router

import (
	"github.com/gorilla/mux"
	"github.com/iyiola-dev/go-gres/middleware"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", middleware.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}", middleware.GetBook).Methods("GET")
	router.HandleFunc("/updatebook", middleware.UpdateBook).Methods("PUT")
	router.HandleFunc("/addbook", middleware.AddBook).Methods("post")
	router.HandleFunc("/delbook{id}", middleware.DeleteBook).Methods("DELETE")
	return router
}
