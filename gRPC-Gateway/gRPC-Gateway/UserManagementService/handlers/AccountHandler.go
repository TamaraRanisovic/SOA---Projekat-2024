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
	"userservice.com/proto/auth"
	"userservice.com/proto/users"
	pb "userservice.com/proto/users" // Update this to the correct import path
	"userservice.com/service"
)

type AccountHandler struct {
	users.UnimplementedUserServiceServer
	AccountService    *service.AccountService
	AuthServiceClient auth.AuthServiceClient
}

func NewAccountHandler(accountService *service.AccountService, authServiceClient auth.AuthServiceClient) *AccountHandler {
	return &AccountHandler{
		AccountService:    accountService,
		AuthServiceClient: authServiceClient,
	}
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string     `json:"username"`
	Role     model.Role `json:"role"`
	jwt.StandardClaims
}

func (handler *AccountHandler) AuthenticateGuide(ctx context.Context, req *pb.TokenRequest) (*empty.Empty, error) {
	var role model.Role = 1

	// Call the gRPC service to decode the token
	resp, err := handler.AuthServiceClient.DecodeToken(ctx, &auth.DecodeTokenRequest{Token: req.Token})
	if err != nil {
		log.Printf("Failed to decode token: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to decode token")
	}

	// Check if the response indicates a successful decoding
	if !resp.IsValid {
		log.Println("Failed to decode token")
		return nil, status.Errorf(codes.Unauthenticated, "Failed to decode token")
	}

	// Check if the decoded token has the expected role
	if resp.Role != int32(role) {
		log.Println("Access denied")
		return nil, status.Errorf(codes.PermissionDenied, "Access denied")
	}

	// Return success if the token is valid and has the expected role
	return &empty.Empty{}, nil
}

func (handler *AccountHandler) GetUserByToken(ctx context.Context, req *pb.TokenRequest) (*pb.UserIdResponse, error) {
	// Call the gRPC service to decode the token
	resp, err := handler.AuthServiceClient.DecodeToken(ctx, &auth.DecodeTokenRequest{Token: req.Token})
	if err != nil {
		log.Printf("Failed to decode token: %v", err)
		return nil, status.Errorf(codes.Internal, "Failed to decode token")
	}

	// Check if the response indicates a successful decoding
	if !resp.IsValid {
		log.Println("Failed to decode token")
		return nil, status.Errorf(codes.Unauthenticated, "Failed to decode token")
	}

	// If the token is valid, retrieve the user ID using the decoded username
	userID, err := handler.AccountService.FindUserIDByUsername(resp.Username)
	if err != nil {
		log.Printf("Failed to find user ID for username %s: %v", resp.Username, err)
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	// Return the user ID in the response
	return &pb.UserIdResponse{Id: userID}, nil
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
