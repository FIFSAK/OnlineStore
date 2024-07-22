package main

import (
	"OnlineStore/api-gateway/routes"
	_ "OnlineStore/docs"
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Online Store Service API
// version 1.0
// @description This is online store service API
// @host onlinestore-bq6f.onrender.com
// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	router := mux.NewRouter()
	routes.Routes(router)

	port := "10000"
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
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
