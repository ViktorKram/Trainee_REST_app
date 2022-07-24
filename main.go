package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	book1 := Book{
		ID:     "1",
		Title:  "Война и Мир",
		Author: &Author{Firsname: "Лев", Lastname: "Толстой"}}

	book2 := Book{
		ID:     "2",
		Title:  "Преступление и наказание",
		Author: &Author{Firsname: "Фёдор", Lastname: "Достоевский"}}

	booksToAdd := []Book{book1, book2}

	books = append(books, booksToAdd...)

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
