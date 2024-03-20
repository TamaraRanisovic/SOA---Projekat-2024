package handler

import (
	"database-example/dto"
	"database-example/service"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type LoginHandler struct {
	AccountService *service.AccountService
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
		return
	}

	// Authenticate user using the provided credentials
	account, _ := loginHandler.AccountService.FindAccountByUsernameAndPassword(creds.Username, creds.Password)
	if account == nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Generate JWT token if authentication succeeds
	tokenString, err := createToken(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
