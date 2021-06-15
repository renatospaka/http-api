package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port = ":8000"
	router := mux.NewRouter()
	
	router.HandleFunc("/", upAndRunning).Methods("GET") 
	router.HandleFunc("/posts", getPosts).Methods("GET")
	//http.Handle("/", router)

	log.Println("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}