package routes

import (
	"github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/controllers"
	"github.com/gorilla/mux"
)

var RouterBook = func(routes *mux.Router) *mux.Router {
	routes.HandleFunc("/api/index", controllers.Index).Methods("GET")
	routes.HandleFunc("/api/books", controllers.GetAllBooks).Methods("GET")
	routes.HandleFunc("/api/books/{id}", controllers.GetBook).Methods("GET")
	routes.HandleFunc("/api/books", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/api/books/{id}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/api/books/{id}", controllers.DeletedBook).Methods("DELETE")

	return routes
}
