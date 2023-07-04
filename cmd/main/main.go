package main

import (
	"log"
	"net/http"

	"github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/config"
	"github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/models"
	"github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/routes"
	"github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/utils"
	"github.com/gorilla/mux"
)

func main() {

	db, errDb := config.Connect()

	if errDb != nil {
		panic(errDb)
	}

	db.AutoMigrate(&models.Books{})

	router := mux.NewRouter()
	r := routes.RouterBook(router)

	servidor := &http.Server{
		Handler: r,
		Addr:    utils.GoDotEnv("PORT"),
		// Addr: ":5454",
	}

	err := servidor.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
