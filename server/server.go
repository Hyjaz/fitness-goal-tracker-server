package server

import (
	"context"
	"log"
	"net/http"
	"time"
)

var server *http.Server

// Init starts web server
func Init() {
	server = &http.Server{
		Addr:    ":8080",
		Handler: initRoute()}
	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
}

// Shutdown gracefully shuts down server
func Shutdown() {
	log.Println("Shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
