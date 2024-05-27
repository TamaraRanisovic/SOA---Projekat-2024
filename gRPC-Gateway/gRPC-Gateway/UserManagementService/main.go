package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"userservice.com/model"
	"userservice.com/repo"
	"userservice.com/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	handler "userservice.com/handlers"
	"userservice.com/proto/auth"
	"userservice.com/proto/users"
)

func initDB() *gorm.DB {
	connectionStr := "root:root@tcp(database:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}

	database.AutoMigrate(&model.Account{})
	database.AutoMigrate(&model.User{})

	user := model.User{
		Name:      "Andjela",
		Surname:   "Radojevic",
		Picture:   "slika.png",
		Biography: "Opsi",
		Moto:      "Ide gas",
	}
	account := model.Account{
		Username:  "aya",
		Password:  "123",
		Email:     "aya@email.com",
		Role:      0,
		IsBlocked: false,
		User:      user,
	}
	database.Create(&account)

	user1 := model.User{
		Name:      "Tamara",
		Surname:   "Ranisovic",
		Picture:   "slika.png",
		Biography: "Opis",
		Moto:      "Ide gas",
	}
	account1 := model.Account{
		Username:  "tamara",
		Password:  "123",
		Email:     "tamara@email.com",
		Role:      0,
		IsBlocked: false,
		User:      user1,
	}
	database.Create(&account1)

	return database
}

func main() {
	database := initDB()
	if database == nil {
		log.Fatal("Failed to connect to the database")
		return
	}

	authConn, err := grpc.Dial("auth-service:8084", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to AuthService: %v", err)
	}
	defer authConn.Close()

	// Create a client instance for the user service
	authServiceClient := auth.NewAuthServiceClient(authConn)

	authResp, err := authServiceClient.Login(context.Background(), &auth.LoginRequest{
		Username: "aya",
		Password: "123",
	})
	if err != nil {
		log.Fatalf("Failed to call Login method: %v", err)
	}

	// Handle the response from the Auth service
	if authResp.Success {
		log.Println("Login successful")
		log.Println("Token:", authResp.Token)
	} else {
		log.Println("Login failed:", authResp.Message)
	}

	repo := &repo.AccountRepository{DatabaseConnection: database}
	service := &service.AccountService{AccountRepo: repo}
	handler := &handler.AccountHandler{AccountService: service, AuthServiceClient: authServiceClient}

	listener, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("Failed to listen on port 8089: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	users.RegisterUserServiceServer(grpcServer, handler)

	go func() {
		log.Println("Starting gRPC server on port 8089")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	<-stopCh
	log.Println("Received shutdown signal, gracefully stopping gRPC server")
	grpcServer.GracefulStop()
	log.Println("gRPC server stopped")
}
