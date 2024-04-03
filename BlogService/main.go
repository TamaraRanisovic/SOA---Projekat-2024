package main

import (
	"log"
	"net/http"
	"time"

	"blogservice.com/handler"
	"blogservice.com/model"
	"blogservice.com/repo"
	"blogservice.com/service"

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
	database.AutoMigrate(&model.Picture{})

	database.AutoMigrate(&model.Blog{})

	// Create a new blog
	blog1 := model.Blog{
		Title:       "First Blog",
		Description: "This is a first blog description.",
		DateCreated: time.Now(),
	}

	// Save the blog to the database
	database.Create(&blog1)

	// Create some pictures
	pictures1 := []model.Picture{
		{URL: "picture1.jpg", BlogID: blog1.ID},
		{URL: "picture2.jpg", BlogID: blog1.ID},
	}

	// Save the pictures to the database
	for _, picture := range pictures1 {
		database.Create(&picture)
	}

	blog2 := model.Blog{
		Title:       "Second Blog",
		Description: "This is a second blog description.",
		DateCreated: time.Now(),
	}

	// Save the blog to the database
	database.Create(&blog2)

	// Create some pictures
	pictures2 := []model.Picture{
		{URL: "picture3.jpg", BlogID: blog2.ID},
		{URL: "picture4.jpg", BlogID: blog2.ID},
	}

	// Save the pictures to the database
	for _, picture := range pictures2 {
		database.Create(&picture)
	}

	return database
}

func startServer(handler *handler.BlogHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/blog/all-blogs/", handler.GetAll).Methods("GET")

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
	repo := &repo.BlogRepository{DatabaseConnection: database}
	service := &service.BlogService{BlogRepo: repo}
	handler := &handler.BlogHandler{BlogService: service}

	startServer(handler)
}
