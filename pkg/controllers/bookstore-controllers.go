package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/config"
	"github.com/MatiSierraDev/bookstore-add-db-iteracion-0/pkg/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm/clause"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", r.URL)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var book []models.Books
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	db.Find(&book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	var book models.Books
	params := mux.Vars(r)["id"]
	result := db.First(&book, params)

	if result.RowsAffected == 0 {
		panic(result.Error)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Books
	json.NewDecoder(r.Body).Decode(&book)

	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	db.Create(&book)
	json.NewEncoder(w).Encode(&book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {

	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	var book models.Books

	params := mux.Vars(r)["id"]
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		panic(err)
	}
	paramsId, err := strconv.Atoi(params)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	book.ID = uint(paramsId)

	if result := db.Where("id = ?", params).Updates(&book); result.RowsAffected == 0 {
		http.Error(w, "Error de solicitud", 404)
		return
	}
	json.NewEncoder(w).Encode(&book)
}

func DeletedBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db, err := config.Connect()

	if err != nil {
		panic(err)
	}

	var book []models.Books

	params := mux.Vars(r)["id"]
	paramsId, err := strconv.Atoi(params)
	if err != nil {
		panic(err)
	}

	// book.ID = uint(paramsId)
	result := db.Clauses(clause.Returning{}).Where("id = ?", paramsId).Delete(&book)
	if result.RowsAffected == 0 {
		http.Error(w, "Error de solicitud", 404)
		return
	}
}
