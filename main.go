package main

import (
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
)


func main() {
	//Init Router
	r := mux.NewRouter();

	r.HandleFunc("/api/books",getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/books",createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}",deleteBook).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":6000",r))
}
