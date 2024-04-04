package handler

import (
	"blogservice.com/dto"
	"blogservice.com/model"
	"blogservice.com/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CommentHandler struct {
	CommentService *service.CommentService
}

func (handler *CommentHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	log.Printf("Comment with id %s", id)
	comment, err := handler.CommentService.FindComment(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(comment)
}

// Function for getting all Accounts
func (handler *CommentHandler) GetAll(writer http.ResponseWriter, req *http.Request) {

	comment, err := handler.CommentService.FindAllComments()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(comment)
}

func (handler *CommentHandler) GetCommenterInfo(w http.ResponseWriter, req *http.Request) {
	var commenterId string
	err := json.NewDecoder(req.Body).Decode(&commenterId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request body\n"))
		return
	}

	// Make a GET request to User Management microservice to get commenter details
	getCommenterbyId := "http://user-management-service:8081/admin/all-accounts/" + commenterId
	resp, err := http.Get(getCommenterbyId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to make GET request to User Management microservice\n"))
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get user details\n"))
		return
	}

	var user dto.UserDto
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to decode user data\n"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

// Function for creating new commment
// TODO append comment id to blog
func (handler *CommentHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var comment model.Comment
	err := json.NewDecoder(req.Body).Decode(&comment)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.CommentService.Create(&comment)
	if err != nil {
		println("Error while creating new comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *CommentHandler) UpdateComment(writer http.ResponseWriter, req *http.Request) {
	var updatedComment model.Comment
	err := json.NewDecoder(req.Body).Decode(&updatedComment)
	if err != nil {
		println("Error while parsing JSON")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	id := mux.Vars(req)["id"]

	existingComment, err := handler.CommentService.FindComment(id)
	if err != nil {
		println("Error while retrieving existing comment")
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// Update the existing comment with new content
	existingComment.Comment = updatedComment.Comment

	// Save the updated comment
	err = handler.CommentService.Update(existingComment)
	if err != nil {
		println("Error while updating comment")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	// Respond with the updated comment
	responseJSON, err := json.Marshal(existingComment)
	if err != nil {
		println("Error while encoding JSON response")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(responseJSON)
}
