package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"servers/server"
	"syscall"
	"time"
)

func main() {
	// Create a new server handler for http
	r := server.New()
	srv := &http.Server{
		Addr:    ":3050",
		Handler: r,
	}

	// Prepare Graceful shutdown signal
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Run Server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error listen: %s\n", err)
		}
	}()

	// Notify server address in terminal
	fmt.Println("Server started to serve at http://localhost:3050/")

	//  ============= Graceful Shutdown ================
	<-done
	fmt.Println("Detect Signal to shutdown.")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown Failed:%+v\n", err)
	}
	fmt.Println("All service Exited Properly")

}
