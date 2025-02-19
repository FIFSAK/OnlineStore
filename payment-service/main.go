package main

import (
	db "OnlineStore"
	"OnlineStore/payment-service/controllers"
	"OnlineStore/payment-service/repository"
	"OnlineStore/payment-service/routes"
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	database, err := db.InitializeDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Uncomment to run migrations
	// if err := db.MigrateUp(database); err != nil {
	//     log.Fatalf("Error running migrations: %v", err)
	// }

	productModel := repository.NewPaymentRepository(database)
	productController := controllers.NewPaymentController(productModel)

	router := mux.NewRouter()
	routes.Routes(router, productController)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{os.Getenv("BASE_URL")},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "10004"
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: corsHandler,
	}

	go gracefulShutdown(server)

	log.Printf("Server is starting on port %s\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server startup failed: %v\n", err)
	}

	log.Println("Server gracefully stopped")
}

func gracefulShutdown(server *http.Server) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v\n", err)
	}
}
