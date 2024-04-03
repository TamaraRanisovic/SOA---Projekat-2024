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
	blog := model.Blog{
		Title:       "Sample Blog",
		Description: "This is a sample blog description.",
		DateCreated: time.Now(),
	}

	// Save the blog to the database
	database.Create(&blog)

	// Create some pictures
	pictures := []model.Picture{
		{URL: "picture1.jpg", BlogID: blog.ID},
		{URL: "picture2.jpg", BlogID: blog.ID},
	}

	// Save the pictures to the database
	for _, picture := range pictures {
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
