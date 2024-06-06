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

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var tp *trace.TracerProvider

func initTracer() (*trace.TracerProvider, error) {
	url := os.Getenv("JAEGER_ENDPOINT")
	if len(url) > 0 {
		return initJaegerTracer(url)
	} else {
		return initFileTracer()
	}
}

func initFileTracer() (*trace.TracerProvider, error) {
	log.Println("Initializing tracing to traces.json")
	f, err := os.Create("traces.json")
	if err != nil {
		return nil, err
	}
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	), nil
}

func initJaegerTracer(url string) (*trace.TracerProvider, error) {
	log.Printf("Initializing tracing to jaeger at %s\n", url)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("api_gateway"),
		)),
	), nil
}

func main() {

	log.SetOutput(os.Stderr)

	// OpenTelemetry
	var err error
	tp, err = initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

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

	// Create a new HTTP handler function that wraps the ServeMux
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		gwmux.ServeHTTP(w, r)
	}

	// Wrap the HTTP handler with the tracer
	httpServer := &http.Server{
		Addr:    cfg.Address,
		Handler: http.HandlerFunc(httpHandler),
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM, os.Interrupt)

	<-stopCh

	if err = httpServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
