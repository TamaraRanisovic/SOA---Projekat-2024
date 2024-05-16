package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"example/gateway/config"
	"example/gateway/proto/blogs"
	"example/gateway/proto/tours"

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
