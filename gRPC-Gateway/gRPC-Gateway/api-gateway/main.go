package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example/gateway/config"
	"example/gateway/proto/auth"
	"example/gateway/proto/blogs"
	"example/gateway/proto/tours"
	"example/gateway/proto/users"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.GetConfig()

	// Create a connection to the TourService
	tourConn, err := grpc.DialContext(
		context.Background(),
		cfg.TourServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial TourService:", err)
	}
	defer tourConn.Close()

	// Create a connection to the BlogService
	blogConn, err := grpc.DialContext(
		context.Background(),
		cfg.BlogServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial BlogService:", err)
	}
	defer blogConn.Close()

	authConn, err := grpc.DialContext(
		context.Background(),
		cfg.AuthServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial AuthService:", err)
	}
	defer authConn.Close()

	userConn, err := grpc.DialContext(
		context.Background(),
		cfg.UserServiceAddress,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial AuthService:", err)
	}
	defer userConn.Close()

	gwmux := runtime.NewServeMux()

	// Register TourService handler
	tourClient := tours.NewTourServiceClient(tourConn)
	err = tours.RegisterTourServiceHandlerClient(
		context.Background(),
		gwmux,
		tourClient,
	)
	if err != nil {
		log.Fatalln("Failed to register TourService gateway:", err)
	}

	// Register BlogService handler
	blogClient := blogs.NewBlogServiceClient(blogConn)
	err = blogs.RegisterBlogServiceHandlerClient(
		context.Background(),
		gwmux,
		blogClient,
	)
	if err != nil {
		log.Fatalln("Failed to register BlogService gateway:", err)
	}

	authClient := auth.NewAuthServiceClient(authConn)
	err = auth.RegisterAuthServiceHandlerClient(
		context.Background(),
		gwmux,
		authClient,
	)
	if err != nil {
		log.Fatalln("Failed to register AuthService gateway:", err)
	}

	userClient := users.NewUserServiceClient(userConn)
	err = users.RegisterUserServiceHandlerClient(
		context.Background(),
		gwmux,
		userClient,
	)
	if err != nil {
		log.Fatalln("Failed to register UserService gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    cfg.Address,
		Handler: gwmux,
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM, os.Interrupt)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
