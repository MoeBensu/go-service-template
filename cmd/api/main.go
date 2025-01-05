package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"yourproject/config"
	"yourproject/internal/app"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create app context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize application
	app, err := app.New(ctx, cfg)
	if err != nil {
		log.Fatalf("Failed to initialize app: %v", err)
	}

	// Handle graceful shutdown
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-shutdown
		app.Shutdown(ctx)
		cancel()
	}()

	// Start the application
	if err := app.Run(); err != nil {
		log.Fatalf("Error running app: %v", err)
	}
}
