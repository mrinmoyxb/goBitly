package main

import (
	"context"
	"goBitly/internal/database"
	"goBitly/internal/handler"
	"goBitly/internal/repository"
	"goBitly/internal/router"
	"goBitly/internal/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, reading from system environment")
	}

	pool, err := database.ConnectPostgres()
	if err != nil {
		log.Fatal("error from postgres db connection: ", err)
	}
	defer pool.Close()
	log.Println("✅ postgres connected successfully")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		log.Fatal("unable to find PORT")
	}

	urlRepository := repository.NewURLRepository(pool)
	urlService := services.NewURLService(urlRepository)
	urlHandler := handler.NewURLHandler(urlService)
	app := router.SetUpRouter(urlHandler)

	srv := &http.Server{
		Addr:         ":" + PORT,
		Handler:      app,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Printf("✅ server is running on PORT: %s", PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server failed to start: %v\n", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	log.Println("shutdown signal received, draining connections...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}else {
		log.Println("✅ server shut down cleanly")
	}
}
