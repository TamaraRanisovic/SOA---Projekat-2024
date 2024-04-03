package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"blogservice.com/model"
	"blogservice.com/service"
	"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

// Function for getting blog by given id
func (handler *BlogHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Blog with id %s", id)
	blog, err := handler.BlogService.FindBlog(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blog)
}

// Function for getting all blogs
func (handler *BlogHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	blogs, err := handler.BlogService.FindAllBlogs()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(blogs)
}

// Function for creating a new blog
func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var blog model.Blog
	err := json.NewDecoder(req.Body).Decode(&blog)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Validate the JSON data
	err = model.ValidateJsonBlog(blog)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.BlogService.Create(&blog)
	if err != nil {
		println("Error while creating a new blog")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
