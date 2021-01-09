package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/rs/cors"
)

//Start is gracefull shutdown
func Start(handler http.Handler) {
	addr := ":8080"
	if port := os.Getenv("SERVER_PORT"); len(port) > 0 {
		addr = fmt.Sprintf(":%v", port)
	}

	server := &http.Server{
		Addr:    addr,
		Handler: corsHandler(handler),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
		<-signals

		if err := server.Shutdown(context.Background()); err != nil {

			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := server.ListenAndServe(); err != http.ErrServerClosed {

		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func corsHandler(handler http.Handler) http.Handler {

	allowOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
	allowHeaders := strings.Split(os.Getenv("CORS_ALLOWED_HEADERS"), ",")
	allowMethods := strings.Split(os.Getenv("CORS_ALLOWED_METHODS"), ",")

	crs := cors.New(cors.Options{
		AllowedOrigins: allowOrigins,
		AllowedHeaders: allowHeaders,
		AllowedMethods: allowMethods,
	})
	return crs.Handler(handler)
}
