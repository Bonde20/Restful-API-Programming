

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var books []Book

func createNewBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	var book Book
	json.Unmarshal(reqBody, &book)
	books = append(books, book)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)

}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintln(w, "Endpoint Hit: returnAllBooks")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	Vars := mux.Vars(r)
	Id := Vars["id"]

	for _, book := range books {
		if book.Id == Id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(book)
			break

		}

	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)

	id := vars["id"]

	for index, book := range books {

		if book.Id == id {

			books = append(books[:index], books[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break

		}
	}
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	inputId := vars["id"]
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()

	for i, book := range books {

		if book.Id == inputId {

			books = append(books[:i], books[i+1:]...)
			var updatedBook Book
			json.Unmarshal(reqBody, &updatedBook)
			books = append(books, updatedBook)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break

		}
	}
}

func main() {

	books = []Book{
		{Id: "1", Title: "Basic"},
		{Id: "2", Title: "Comics"},
		{Id: "3", Title: "Robinhood"},
	}

	mx := mux.NewRouter().StrictSlash(true)
  mx.HandleFunc("/books", returnAllBooks).Methods("GET")
	mx.HandleFunc("/books/{id}", returnSingleBook).Methods("GET")
	mx.HandleFunc("/book", createNewBook).Methods("POST")
	mx.HandleFunc("/books/{id}", deleteBook).Methods("Delete")
	mx.HandleFunc("/books/{id}", updateBook).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", mx))

}

type Book struct {
	Id    string `json:"Id"`
	Title string `json:"Title"`
}
