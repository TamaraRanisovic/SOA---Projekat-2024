package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

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

/*
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
*/
/*
func (handler *BlogHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB is the maximum size of the file
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Get form values
	title := r.FormValue("title")
	description := r.FormValue("description")
	dateCreated := r.FormValue("date_created")
	status := r.FormValue("status")

	// Do something with the data, such as printing it
	fmt.Println("Title:", title)
	fmt.Println("Description:", description)
	fmt.Println("Date Created:", dateCreated)
	fmt.Println("Status:", status)

	// Send response to client
	fmt.Fprintf(w, "Blog added successfully!")
}
*/

// Function for creating a new blog
func (handler *BlogHandler) Create(writer http.ResponseWriter, req *http.Request) {
	// Parse form data
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(writer, "Failed to parse form data", http.StatusInternalServerError)
		return
	}

	// Extract form fields
	title := req.Form.Get("title")
	description := req.Form.Get("description")
	dateCreatedStr := req.Form.Get("date_created")
	statusStr := req.Form.Get("status")
	pictures := req.MultipartForm.File["pictures"] // Get slice of picture files

	// Validate form data
	if title == "" || description == "" || dateCreatedStr == "" || statusStr == "" || len(pictures) == 0 {
		http.Error(writer, "All fields are required", http.StatusBadRequest)
		return
	}

	// Parse date created
	dateCreated, err := time.Parse("2006-01-02", dateCreatedStr)
	if err != nil {
		http.Error(writer, "Invalid date format", http.StatusBadRequest)
		return
	}

	// Parse status
	var status model.Status
	switch statusStr {
	case "draft":
		status = model.Draft
	case "published":
		status = model.Published
	case "closed":
		status = model.Closed
	default:
		http.Error(writer, "Invalid status", http.StatusBadRequest)
		return
	}

	// Create blog object
	blog := model.Blog{
		Title:       title,
		Description: description,
		DateCreated: dateCreated,
		Status:      status,
	}

	// Process each uploaded picture
	for _, file := range pictures {
		// Open uploaded file
		uploadedFile, err := file.Open()
		if err != nil {
			http.Error(writer, "Failed to open uploaded file", http.StatusInternalServerError)
			return
		}
		defer uploadedFile.Close()

		// Perform operations with the file, such as saving it to disk or uploading to a cloud storage service
		// For this example, we'll simply generate a URL representing the uploaded file
		// This is just a placeholder; you'll need to replace it with your actual logic
		pictureURL := file.Filename

		// Append the picture URL to the blog
		blog.Pictures = append(blog.Pictures, model.Picture{URL: pictureURL})
	}

	// Create blog using service
	err = handler.BlogService.Create(&blog)
	if err != nil {
		http.Error(writer, "Failed to create blog", http.StatusInternalServerError)
		return
	}

	// Return success response
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}
