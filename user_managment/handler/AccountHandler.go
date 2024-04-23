package handler

import (
	"bytes"
	"database-example/dto"
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"io"
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

func (handler *AccountHandler) AuthenticateGuide(w http.ResponseWriter, req *http.Request) {
	var role model.Role = 1

	var tokenBody struct {
		Token string `json:"token"`
	}

	tokenString := req.Header.Get("Authorization")
	log.Println(tokenString)

	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("No token provided\n"))
		return
	}
	tokenBody.Token = tokenString

	tokenBodyJSON, err := json.Marshal(tokenBody)
	if err != nil {
		log.Println("Failed to marshal tokenBody:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal tokenBody\n"))
		return
	}

	decodeToken := "http://localhost:8082/decode"
	resp, err := http.Post(decodeToken, "application/json", bytes.NewBuffer(tokenBodyJSON))
	if err != nil {
		log.Println("Failed to make POST request to Auth microservice:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to make POST request to Auth microservice\n"))
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Response body:", string(body))
		log.Println("Failed to read response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to read response body\n"))
		return
	}

	// Log the response body
	log.Println("Response body:", string(body))

	// Decode the JSON response body to extract the username, role, and expiration
	var responseBody struct {
		Username   string     `json:"username"`
		Role       model.Role `json:"role"`
		Expiration int64      `json:"exp"`
	}
	if err := json.Unmarshal(body, &responseBody); err != nil {
		log.Println("Failed to decode response body:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to decode response body\n"))
		return
	}

	// Check if the user's role is Guide
	if responseBody.Role != role {
		log.Println("Access denied: user's role is not Guide")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Access denied\n"))
		return
	}
	w.WriteHeader(http.StatusOK)

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
