package handler

import (
	"encoding/json"
	"errors"
	"log"
	"mime/multipart"
	"net/http"
	"time"

	"blogservice.com/model"
	"blogservice.com/service"
	"github.com/gorilla/mux"
)

type BlogHandler struct {
	BlogService *service.BlogService
}

// FormData represents the parsed form data
type FormData struct {
	Title       string
	Description string
	DateCreated string
	Status      string
	Pictures    []*multipart.FileHeader
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
	formData, err := parseFormData(req)
	if err != nil {
		handleError(writer, err, http.StatusInternalServerError)
		return
	}

	// Validate form data
	err = validateFormData(formData)
	if err != nil {
		handleError(writer, err, http.StatusBadRequest)
		return
	}

	// Create blog object
	blog, err := createBlog(formData)
	if err != nil {
		handleError(writer, err, http.StatusInternalServerError)
		return
	}

	// Process uploaded pictures
	err = uploadPictures(req, blog)
	if err != nil {
		handleError(writer, err, http.StatusInternalServerError)
		return
	}

	// Create blog using service
	err = handler.BlogService.Create(blog)
	if err != nil {
		handleError(writer, err, http.StatusInternalServerError)
		return
	}

	// Return success response
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

// Parse form data from the request
func parseFormData(req *http.Request) (FormData, error) {
	err := req.ParseMultipartForm(10 << 20)
	if err != nil {
		return FormData{}, errors.New("Failed to parse form data")
	}

	formData := FormData{
		Title:       req.Form.Get("title"),
		Description: req.Form.Get("description"),
		DateCreated: req.Form.Get("date_created"),
		Status:      req.Form.Get("status"),
		Pictures:    req.MultipartForm.File["pictures"],
	}

	return formData, nil
}

// Validate form data
func validateFormData(formData FormData) error {
	if formData.Title == "" || formData.Description == "" || formData.DateCreated == "" || formData.Status == "" {
		return errors.New("All fields are required")
	}

	return nil
}

// Create blog object
func createBlog(formData FormData) (*model.Blog, error) {
	dateCreated, err := time.Parse("2006-01-02", formData.DateCreated)
	if err != nil {
		return nil, errors.New("Invalid date format")
	}

	status, err := model.GetStatus(formData.Status)
	if err != nil {
		return nil, err // Return the error if status retrieval fails
	}

	blog := &model.Blog{
		Title:       formData.Title,
		Description: formData.Description,
		DateCreated: dateCreated,
		Status:      status,
	}

	return blog, nil
}

// Process uploaded pictures
func uploadPictures(req *http.Request, blog *model.Blog) error {
	for _, file := range req.MultipartForm.File["pictures"] {
		pictureURL := file.Filename
		blog.Pictures = append(blog.Pictures, model.Picture{URL: pictureURL})
	}

	return nil
}

// Handle HTTP errors
func handleError(writer http.ResponseWriter, err error, status int) {
	http.Error(writer, err.Error(), status)
}
