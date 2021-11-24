
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

	defer r.Body.Close()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var book Book
	json.Unmarshal(reqBody, &book)
	books = append(books, book)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)

}

func returnAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Endpoint Hit: returnAllBooks")

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func returnSingleBook(w http.ResponseWriter, r *http.Request) {
	Vars := mux.Vars(r)
	Id := Vars["id"]
	for _, book := range books {
		if book.Id == Id {

			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(book)
			if err != nil {

				log.Println("Error Getting required book:", err)

			}

		}
	}

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

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


func main() {

	books = []Book{
		{Id: "1", Title: "Basic"},
		{Id: "2", Title: "Comics"},
	}

	mx := mux.NewRouter().StrictSlash(true)

	mx.HandleFunc("/books", returnAllBooks).Methods("GET")
	mx.HandleFunc("/book/{id}", returnSingleBook).Methods("GET")
	mx.HandleFunc("/book", createNewBook).Methods("POST")
	mx.HandleFunc("/book/{id}", deleteBook).Methods("Delete")
	log.Fatal(http.ListenAndServe(":8080", mx))

}

type Book struct {
	Id    string `json:"Id"`
	Title string `json:"Title"`
}
