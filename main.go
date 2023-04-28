package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorila/mux"
)

type Book struct {
	ID            string  `json:"id"`
	Name          string  `json:"name"`
	Isbn          string  `json:"isbn"`
	Discount      string  `json:"discount"`
	TypeOfTheBook string  `json:"typeofthebook"`
	Price         *Price  `json:"price"`
	Author        *Author `json:"author"`
}

type Author struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	NoOfBooksWritten string `json:"noofbookswritten"`
	Place            string `json:"place"`
	Language         string `json:"language"`
}

type Price struct {
	OnlinePrice  string `json:"onlineprice"`
	OfflinePrice string `json:"offlineprice"`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["ID"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode("Item not found")
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.Id == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			json.NewDecoder(r.Body).Decode(&book)
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["ID"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {
	r := mux.NewRouter()

	books: = append(books, Book{ID: "1", Name: "Kavya", TypeOfTheBook: "comedy", Isbn: "4", Discount: "15", Author: &Author{Id: "SS", Name: "Rajamoule", NoOfBooksWritten: "2", Place: "Chennai", Language: "Tamil"}, Price: &Price{OnlinePrice: "700", OfflinePrice: "1000"}})

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/book/{ID}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{ID}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{ID}", deleteBook).Methods("DELETE")

	fmt.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
