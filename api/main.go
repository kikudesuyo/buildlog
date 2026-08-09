package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/kikudesuyo/buildlog/api/library"
	"github.com/kikudesuyo/buildlog/api/route"
)

// main はアプリケーションを起動します。
func main() {
	portFlag := flag.String("port", "8081", "port to run HTTP server on")
	flag.Parse()

	port := *portFlag
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	err := library.InitDB()
	if err != nil {
		panic("DB Error")
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: route.NewRouter(),
	}

	log.Printf("API server is running on port %s", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("API server failed to start: %v", err)
	}
}
