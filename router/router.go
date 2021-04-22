package router

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", middleware.getBooks).Methods("GET")
	router.HandleFunc("/book{id}", middleware.getBook).Methods("GET")
	router.HandleFunc("/book{id}", middleware.updateBook).Methods("PUT")
	router.HandleFunc("/book", middleware.addBook()).Methods("post")
	router.HandleFunc("/book{id}", middleware.deleteBook).Methods("DELETE")
	return router
}
