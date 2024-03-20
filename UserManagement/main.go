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

	database.AutoMigrate(&model.Student{})
	database.Exec("INSERT IGNORE INTO students VALUES ('aec7e123-233d-4a09-a289-75308ea5b7e6', 'Marko Markovic', 'Graficki dizajn')")

	database.AutoMigrate(&model.Account{})
	database.Exec("INSERT IGNORE INTO `students`.`accounts`  VALUES ('1','aya', '123', 'aya@email.com', 'Administrator', '0')")

	database.AutoMigrate(&model.User{})
	database.Exec("INSERT INTO `students`.`users` (`id`, `username`, `password`, `email`, `role`, `is_blocked`, `name`, `surname`, `picture`, `biography`, `moto`) VALUES ('1', 'aya', '123', 'aya@email.com', 'Administrator', '0', 'Andjela', 'Radojevic', 'slika.png', 'Opsi', 'Ide gas')")

	database.AutoMigrate(&model.Rating{})
	database.AutoMigrate(&model.Blog{})
	
	return database
}

func startServer(handler *handler.StudentHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/students/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/students", handler.Create).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repo.StudentRepository{DatabaseConnection: database}
	service := &service.StudentService{StudentRepo: repo}
	handler := &handler.StudentHandler{StudentService: service}

	startServer(handler)
}
