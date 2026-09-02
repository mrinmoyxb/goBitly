package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"goBitly/internal/database"
	"log"
	"net/http"
	"os"
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

	app := chi.NewRouter()
	PORT := os.Getenv("PORT")

	if PORT == "" {
		log.Fatal("unable to find PORT")
	}

	address := ":" + PORT
	log.Printf("server is running on PORT: %s", PORT)
	if err := http.ListenAndServe(address, app); err != nil {
		log.Fatalf("server failed to start: %v\n", err)
	}
}
