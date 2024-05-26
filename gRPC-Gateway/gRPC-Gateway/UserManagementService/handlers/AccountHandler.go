package handler

import (
	"context"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"userservice.com/model"
	"userservice.com/service"

	pb "userservice.com/proto/users" // Update this to the correct import path
)

type AccountHandler struct {
	pb.UnimplementedUserServiceServer
	AccountService *service.AccountService
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string     `json:"username"`
	Role     model.Role `json:"role"`
	jwt.StandardClaims
}

func (handler *AccountHandler) AuthenticateGuide(ctx context.Context, req *pb.TokenRequest) (*emptypb.Empty, error) {
	var role model.Role = 1

	// Decode token logic (You may need to change this based on your actual implementation)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to parse token: %v", err)
	}

	if !token.Valid || claims.Role != role {
		return nil, status.Errorf(codes.PermissionDenied, "Access denied")
	}

	return &empty.Empty{}, nil
}

func (handler *AccountHandler) GetUserByToken(ctx context.Context, req *pb.TokenRequest) (*pb.UserIdResponse, error) {
	// Decode token logic (You may need to change this based on your actual implementation)
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Failed to parse token: %v", err)
	}

	if !token.Valid {
		return nil, status.Errorf(codes.PermissionDenied, "Invalid token")
	}

	account, err := handler.AccountService.FindAccountByUsername(claims.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}

	return &pb.UserIdResponse{Id: account.ID.String()}, nil
}

func (handler *AccountHandler) BlockAccount(ctx context.Context, req *pb.AccountIdRequest) (*emptypb.Empty, error) {
	err := handler.AccountService.BlockAccount(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}
	return &empty.Empty{}, nil
}

func (handler *AccountHandler) GetByUsernameAndPassword(ctx context.Context, req *pb.Credentials) (*pb.AccountDetailResponse, error) {
	account, err := handler.AccountService.FindAccountByUsernameAndPassword(req.Username, req.Password)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}
	return convertToAccountDetailResponse(account), nil
}

func (handler *AccountHandler) GetByUsername(ctx context.Context, req *pb.UsernameRequest) (*pb.AccountDetailResponse, error) {
	account, err := handler.AccountService.FindAccountByUsername(req.Username)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}
	return convertToAccountDetailResponse(account), nil
}

func (handler *AccountHandler) GetAccount(ctx context.Context, req *pb.UserIdRequest) (*pb.AccountDetailResponse, error) {
	account, err := handler.AccountService.FindAccount(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found")
	}
	return convertToAccountDetailResponse(account), nil
}

func (handler *AccountHandler) GetAllAccounts(ctx context.Context, req *empty.Empty) (*pb.AccountsResponse, error) {
	accounts, err := handler.AccountService.FindAllAccounts()
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "No accounts found")
	}

	var accountDetails []*pb.AccountDetailResponse
	for _, account := range *accounts {
		accountDetails = append(accountDetails, convertToAccountDetailResponse(&account))
	}

	return &pb.AccountsResponse{Accounts: accountDetails}, nil
}

func (handler *AccountHandler) CreateAccount(ctx context.Context, req *pb.AccountDetailResponse) (*emptypb.Empty, error) {
	account := convertToModelAccount(req)
	err := handler.AccountService.Create(account)
	if err != nil {
		return nil, status.Errorf(codes.FailedPrecondition, "Failed to create account")
	}
	return &empty.Empty{}, nil
}

// Helper function to convert AccountDetailResponse to model.Account
func convertToModelAccount(detail *pb.AccountDetailResponse) *model.Account {
	accountID, err := uuid.Parse(detail.Id)
	if err != nil {
		log.Printf("Invalid account ID: %v", err)
		return nil
	}

	userID, err := uuid.Parse(detail.User.Id)
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		return nil
	}

	return &model.Account{
		ID:        accountID,
		Username:  detail.Username,
		Email:     detail.Email,
		Role:      model.Role(detail.Role),
		IsBlocked: detail.Isblocked,
		User: model.User{
			UserID:    userID,
			Name:      detail.User.Name,
			Surname:   detail.User.Surname,
			Picture:   detail.User.Picture,
			Biography: detail.User.Biography,
			Moto:      detail.User.Moto,
		},
	}
}

// Helper function to convert model.Account to AccountDetailResponse
func convertToAccountDetailResponse(account *model.Account) *pb.AccountDetailResponse {

	return &pb.AccountDetailResponse{
		Id:        account.ID.String(),
		Username:  account.Username,
		Email:     account.Email,
		Role:      int32(account.Role),
		Isblocked: account.IsBlocked,
		User: &pb.UserDetail{
			Id:        account.User.UserID.String(),
			Name:      account.User.Name,
			Surname:   account.User.Surname,
			Picture:   account.User.Picture,
			Biography: account.User.Biography,
			Moto:      account.User.Moto,
		},
	}
}
