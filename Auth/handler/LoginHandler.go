package handler

import (
	"bytes"
	"database-example/dto"
	"database-example/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginHandler struct {
}

// Secret key used for JWT token signing
var jwtKey = []byte("my_secret_key")

// Claims for the JWT claims structure
type Claims struct {
	Username string     `json:"username"`
	Role     model.Role `json:"role"`
	jwt.StandardClaims
}

// Create a new JWT token with the given username, role and expiration time
func createToken(username string, role model.Role) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute) // Expiration time in Unix time (seconds)
	claims := &Claims{
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key and get the token string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (loginHandler *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into Credentials struct
	var creds dto.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request body\n"))
		return
	}

	// Marshal the credentials into JSON
	credsJSON, err := json.Marshal(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal credentials\n"))
		return
	}

	// Make a POST request to User Management microservice to authenticate the user
	getByUsernameAndPasswordURL := "http://localhost:8081/accounts/get"
	resp, err := http.Post(getByUsernameAndPasswordURL, "application/json", bytes.NewBuffer(credsJSON))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to make POST request to User Management microservice\n"))
		return
	}

	// If the authentication fails, return Unauthorized status
	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Failed to authenticate user\n"))
		return
	}

	// Decode the account data from the response body
	var account model.Account
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to decode account data\n"))
		return
	}

	// Generate JWT token for the authenticated user
	tokenString, err := createToken(creds.Username, account.Role)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to generate token\n"))
		return
	}

	// Set the token in response headers
	w.Header().Set("Authorization", "Bearer "+tokenString)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Return success message and token
	w.Write([]byte("You've successfully logged in!\n"))
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

	defer resp.Body.Close()

}

// Decodes and verifies JWT token
func (loginHandler *LoginHandler) DecodeToken(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a struct containing token string
	var tokenBody struct {
		Token string `json:"token"`
	}

	err := json.NewDecoder(r.Body).Decode(&tokenBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request body\n"))
		return
	}

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenBody.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Failed to parse token: " + err.Error() + "\n"))
		return
	}

	// Verify if token is valid and extract claims
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Invalid token\n"))
		return
	}

	// Return decoded claims
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(claims)
}
