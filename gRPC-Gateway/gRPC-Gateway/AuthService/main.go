package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	handler "authservice.com/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"authservice.com/proto/auth"
)

func main() {
	loginHandler := &handler.LoginHandler{}

	listener, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("Failed to listen on port 8084: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	auth.RegisterAuthServiceServer(grpcServer, loginHandler)

	go func() {
		log.Println("Starting gRPC server on port 8084")
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
