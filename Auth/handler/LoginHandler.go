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

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func createToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (loginHandler *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds dto.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode request body\n"))
		return
	}

	credsJSON, err := json.Marshal(creds)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to marshal credentials\n"))
		return
	}

	getByUsernameAndPasswordURL := "http://localhost:8081/accounts/get"
	resp, err := http.Post(getByUsernameAndPasswordURL, "application/json", bytes.NewBuffer(credsJSON))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to make POST request to User Management microservice\n"))
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Failed to authenticate user\n"))
		return
	}

	var account model.Account
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to decode account data\n"))
		return
	}

	tokenString, err := createToken(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to generate token\n"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("You've successfully logged in!\n"))
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}
