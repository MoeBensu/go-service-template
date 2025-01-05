package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"yourproject/config"
	"yourproject/internal/app"
	"yourproject/pkg/version"
)

func main() {
	// Define flags
	versionFlag := flag.Bool("version", false, "Print version information")
	flag.Parse()

	// Handle version flag
	if *versionFlag {
		info := version.Get()
		fmt.Printf("Version:\t%s\n", info.Version)
		fmt.Printf("Commit:\t\t%s\n", info.CommitSHA)
		fmt.Printf("Build Time:\t%s\n", info.BuildTime)
		os.Exit(0)
	}

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
