package main

import (
	"database-example/handler"
	"database-example/model"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "root:root@tcp(localhost:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Account{})
	database.AutoMigrate(&model.User{})

	user := model.User{
		Name:      "Andjela",
		Surname:   "Radojevic",
		Picture:   "slika.png",
		Biography: "Opsi",
		Moto:      "Ide gas",
	}
	account := model.Account{
		Username:  "aya",
		Password:  "123",
		Email:     "aya@email.com",
		Role:      0,
		IsBlocked: false,
		User:      user,
	}
	database.Create(&account)

	return database
}

func startServer(loginHandler *handler.LoginHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/login", loginHandler.Login).Methods("POST")
	router.HandleFunc("/decode", loginHandler.DecodeToken).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
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
