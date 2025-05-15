package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
	"github.com/nutcase/dao-golang/x/taskbounty/client/rest"
)

func main() {
	log.Println("Starting DAO Task Bounty System...")

	// Initialize Cosmos SDK client context
	clientCtx := client.Context{}

	// Initialize router using gorilla/mux
	router := mux.NewRouter()

	// Register REST routes
	rest.RegisterRoutes(clientCtx, router)

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":1317", // Standard Cosmos SDK REST port
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Printf("Server listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down gracefully...")

	// Shutdown server with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown: %v", err)
	}
}
