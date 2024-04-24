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
	connectionStr := "root:root@tcp(database:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
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

	user1 := model.User{
		Name:      "Tamara",
		Surname:   "Ranisovic",
		Picture:   "slika.png",
		Biography: "Opis",
		Moto:      "Ide gas",
	}
	account1 := model.Account{
		Username:  "tamara",
		Password:  "123",
		Email:     "tamara@email.com",
		Role:      0,
		IsBlocked: false,
		User:      user1,
	}
	database.Create(&account1)

	return database
}

func startServer(handler *handler.AccountHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/admin/all-accounts/", handler.GetAll).Methods("GET")
	router.HandleFunc("/admin/all-accounts/{id}/", handler.Get).Methods("GET")
	router.HandleFunc("/admin/block/{id}/", handler.Block).Methods("PUT")
	router.HandleFunc("/add-account/", handler.Create).Methods("POST")
	router.HandleFunc("/authenticate-guide/", handler.AuthenticateGuide).Methods("POST")

	router.HandleFunc("/accounts/get", handler.GetByUsernameAndPassword).Methods("POST")
	//cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8085", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repo.AccountRepository{DatabaseConnection: database}
	service := &service.AccountService{AccountRepo: repo}
	handler := &handler.AccountHandler{AccountService: service}

	startServer(handler)
}
