package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"blogservice/model"

	handler "blogservice/handlers"
	"blogservice/proto/blogs"
	"blogservice/repo"
	"blogservice/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "root:root@tcp(database:3306)/students?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(connectionStr), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
		return nil
	}

	database.AutoMigrate(&model.Picture{})
	database.AutoMigrate(&model.Blog{})

	// Create sample data
	blogs := []model.Blog{
		{
			Title:       "First Blog",
			Description: "This is the first blog description.",
			DateCreated: time.Now(),
		},
		{
			Title:       "Second Blog",
			Description: "This is the second blog description.",
			DateCreated: time.Now(),
		},
	}

	for _, blog := range blogs {
		database.Create(&blog)
		pictures := []model.Picture{
			{URL: "picture1.jpg", BlogID: blog.ID},
			{URL: "picture2.jpg", BlogID: blog.ID},
		}
		for _, picture := range pictures {
			database.Create(&picture)
		}
	}

	return database
}

func main() {
	database := initDB()
	if database == nil {
		log.Fatal("Failed to connect to the database")
		return
	}

	repo := &repo.BlogRepository{DatabaseConnection: database}
	service := &service.BlogService{BlogRepo: repo}
	handler := &handler.BlogHandler{BlogService: service}

	// Setup gRPC server
	listener, err := net.Listen("tcp", ":8085")
	if err != nil {
		log.Fatalf("Failed to listen on port 8085: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	blogs.RegisterBlogServiceServer(grpcServer, handler)

	go func() {
		log.Println("Starting gRPC server on port 8085")
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	// Handle graceful shutdown
	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	<-stopCh
	log.Println("Received shutdown signal, gracefully stopping gRPC server")
	grpcServer.GracefulStop()
	log.Println("gRPC server stopped")
}
