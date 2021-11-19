
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

	reqBody, _ := ioutil.ReadAll(r.Body)
	var book Book
	json.Unmarshal(reqBody, &book)
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint Hit: returnAllBooks")

	json.NewEncoder(w).Encode(books)
}

func main() {

	books = []Book{
		{Id: "1", Title: "Basic"},
		{Id: "2", Title: "Comics"},
	}

	mx := mux.NewRouter().StrictSlash(true)

	mx.HandleFunc("/books", returnAllBooks).Methods("GET")
	mx.HandleFunc("/book", createNewBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", mx))
}

type Book struct {
	Id    string `json:"Id"`
	Title string `json:"Title"`
}
