package main

import (
	"database-example/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Initialize the database connection
func initDB() *gorm.DB {
	connectionStr := "root:root@tcp(localhost:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	return database
}

// Start the HTTP server
func startServer(loginHandler *handler.LoginHandler) {
	router := mux.NewRouter().StrictSlash(true)

	// Define routes
	router.HandleFunc("/login", loginHandler.Login).Methods("POST")
	router.HandleFunc("/decode", loginHandler.DecodeToken).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")

	// Start the HTTP server on port 8082
	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	loginHandler := &handler.LoginHandler{}

	startServer(loginHandler)
}
