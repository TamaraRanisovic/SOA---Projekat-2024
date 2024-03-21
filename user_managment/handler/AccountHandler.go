package handler

import (
	"database-example/dto"
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type AccountHandler struct {
	AccountService *service.AccountService
	UserService    *service.UserService
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string     `json:"username"`
	Role     model.Role `json:"role"`
	jwt.StandardClaims
}

// Function for checking if a logged in account has role Administrator
func authenticate(writer http.ResponseWriter, req *http.Request) {
	var role model.Role = 0
	tokenString := req.Header.Get("Authorization")
	if tokenString == "" {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("No token provided\n"))
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		writer.Write([]byte("Failed to parse token: " + err.Error() + "\n"))
		return
	}

	// Check if the token is valid and the user's role allows access
	if !token.Valid || claims.Role != role {
		writer.WriteHeader(http.StatusForbidden)
		writer.Write([]byte("Access denied\n"))
		return
	}
}

// Function for getting Account by given id
// Printing into terminal
// Returning json object
func (handler *AccountHandler) Get(writer http.ResponseWriter, req *http.Request) {
	authenticate(writer, req)
	id := mux.Vars(req)["id"]
	log.Printf("Account with id %s", id)
	account, err := handler.AccountService.FindAccount(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(account)
	writer.Write([]byte("Function executed successfully\n"))
}

// Function for getting all Accounts
func (handler *AccountHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	authenticate(writer, req)
	accounts, err := handler.AccountService.FindAllAccounts()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(accounts)
}

// Function for creating a new account
func (handler *AccountHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var account model.Account
	err := json.NewDecoder(req.Body).Decode(&account)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.AccountService.Create(&account)
	if err != nil {
		println("Error while creating a new account")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

// Function for blocking an account
func (handler AccountHandler) Block(writer http.ResponseWriter, req *http.Request) {
	authenticate(writer, req)
	id := mux.Vars(req)["id"]
	log.Printf("Account with id %s", id)

	err := handler.AccountService.BlockAccount(id)

	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (handler *AccountHandler) GetByUsernameAndPassword(writer http.ResponseWriter, req *http.Request) {
	// Decode the request body to get credentials
	var creds dto.Credentials
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	// Call the service to find the account by username and password
	account, err := handler.AccountService.FindAccountByUsernameAndPassword(creds.Username, creds.Password)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	// Set response content type and status code
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	// Encode the account into JSON and send the response
	json.NewEncoder(writer).Encode(account)
}
