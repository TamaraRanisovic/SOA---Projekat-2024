package handlers

import (
	"context"
	"log"
	"time"

	"authservice.com/model"

	"authservice.com/proto/auth"
	"authservice.com/proto/users"
	"github.com/dgrijalva/jwt-go"

	"go.opentelemetry.io/otel"
)

type LoginHandler struct {
	auth.UnimplementedAuthServiceServer
	UserServiceClient users.UserServiceClient
}

func NewLoginHandler(userServiceClient users.UserServiceClient) *LoginHandler {
	return &LoginHandler{UserServiceClient: userServiceClient}
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
	tr := otel.Tracer("authservice.handlers.LoginHandler.Login")
	ctx, span := tr.Start(ctx, "LoginHandler.Login")
	defer span.End()

	// Create a request to authenticate the user via gRPC
	user, err := s.UserServiceClient.GetByUsernameAndPassword(ctx, &users.Credentials{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("Failed to authenticate user: %v", err)
		return &auth.LoginResponse{Success: false, Message: "Failed to authenticate user"}, nil
	}

	// Create a JWT token for the authenticated user
	tokenString, err := createToken(req.Username, model.Role(user.Role))
	if err != nil {
		log.Printf("Failed to generate token: %v", err)
		return &auth.LoginResponse{Success: false, Message: "Failed to generate token"}, nil
	}

	// Return the response with the token
	return &auth.LoginResponse{Success: true, Message: "You've successfully logged in!", Token: tokenString}, nil
}

func (s *LoginHandler) DecodeToken(ctx context.Context, req *auth.DecodeTokenRequest) (*auth.DecodeTokenResponse, error) {
	tr := otel.Tracer("authservice.handlers.LoginHandler.DecodeToken")
	ctx, span := tr.Start(ctx, "LoginHandler.DecodeToken")
	defer span.End()

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
		Role:     int32(claims.Role),
		Exp:      time.Unix(claims.ExpiresAt, 0).String(),
	}, nil
}
