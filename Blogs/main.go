package main

import (
	"database-example/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Start the HTTP server
func startServer(commentHandler *handler.CommentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	// Define routes
	router.HandleFunc("/comments", commentHandler.GetAll).Methods("GET")
	router.HandleFunc("/commnets/{id}", commentHandler.Get).Methods("GET")
	router.HandleFunc("/comments/add-comment", commentHandler.Create).Methods("POST")
	router.HandleFunc("/comments/{id}/edit", commentHandler.UpdateComment).Methods("PUT")
	router.HandleFunc("/comments/{id}/commenter-details", commentHandler.GetCommenterInfo).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")

	// Start the HTTP server on port 8082
	log.Fatal(http.ListenAndServe(":8084", router))
}

func main() {

	commentHandler := &handler.CommentHandler{}

	startServer(commentHandler)
}
