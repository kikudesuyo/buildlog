package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/route"
)

func main() {
	portFlag := flag.String("port", "8081", "port to run HTTP server on")
	flag.Parse()

	port := *portFlag
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}

	db, err := library.OpenDatabase(context.Background())
	if err != nil {
		log.Fatalf("database initialization failed: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("database close failed: %v", err)
		}
	}()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: route.NewRouter(db),
	}

	log.Printf("API server is running on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("API server failed to start: %v", err)
	}
}
