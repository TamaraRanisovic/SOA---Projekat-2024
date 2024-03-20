package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
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

func startServer(accountHandler *handler.AccountHandler, loginHandler *handler.LoginHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/accounts/{id}", accountHandler.Get).Methods("GET")
	router.HandleFunc("/accounts", accountHandler.Create).Methods("POST")
	router.HandleFunc("/accounts/log", accountHandler.GetByUsernameAndPassword).Methods("POST")

	router.HandleFunc("/login", loginHandler.Login).Methods("POST")

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
	accountRepo := &repo.AccountRepository{DatabaseConnection: database}
	accountService := &service.AccountService{AccountRepo: accountRepo}
	accountHandler := &handler.AccountHandler{AccountService: accountService}

	loginHandler := &handler.LoginHandler{AccountService: accountService}

	startServer(accountHandler, loginHandler)
}
