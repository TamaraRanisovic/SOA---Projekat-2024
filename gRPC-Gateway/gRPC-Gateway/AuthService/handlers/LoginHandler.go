package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"authservice.com/dto"
	"authservice.com/model"

	"authservice.com/proto/auth"
	"github.com/dgrijalva/jwt-go"
)

type LoginHandler struct {
	auth.UnimplementedAuthServiceServer
}

func NewLoginHandler() *LoginHandler {
	return &LoginHandler{}
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string     `json:"username"`
	Role     model.Role `json:"role"`
	jwt.StandardClaims
}

func createToken(username string, role model.Role) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		Role:     role,
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

func (s *LoginHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	creds := dto.Credentials{
		Username: req.Username,
		Password: req.Password,
	}

	credsJSON, err := json.Marshal(creds)
	if err != nil {
		return &auth.LoginResponse{Success: false, Message: "Failed to marshal credentials"}, nil
	}

	getByUsernameAndPasswordURL := "http://localhost:8085/accounts/get"
	resp, err := http.Post(getByUsernameAndPasswordURL, "application/json", bytes.NewBuffer(credsJSON))
	if err != nil || resp.StatusCode != http.StatusOK {
		return &auth.LoginResponse{Success: false, Message: "Failed to authenticate user"}, nil
	}
	defer resp.Body.Close()

	var account model.Account
	if err := json.NewDecoder(resp.Body).Decode(&account); err != nil {
		return &auth.LoginResponse{Success: false, Message: "Failed to decode account data"}, nil
	}

	tokenString, err := createToken(creds.Username, account.Role)
	if err != nil {
		return &auth.LoginResponse{Success: false, Message: "Failed to generate token"}, nil
	}

	return &auth.LoginResponse{Success: true, Message: "You've successfully logged in!", Token: tokenString}, nil
}

func (s LoginHandler) DecodeToken(ctx context.Context, req *auth.DecodeTokenRequest) (*auth.DecodeTokenResponse, error) {
	token, err := jwt.ParseWithClaims(req.Token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return &auth.DecodeTokenResponse{IsValid: false, Message: "Failed to parse token: " + err.Error()}, nil
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return &auth.DecodeTokenResponse{IsValid: false, Message: "Invalid token"}, nil
	}

	return &auth.DecodeTokenResponse{
		IsValid:  true,
		Username: claims.Username,
		Role:     string(rune(claims.Role)),
		Exp:      time.Unix(claims.ExpiresAt, 0).String(),
	}, nil
}
