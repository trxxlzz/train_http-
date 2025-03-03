package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"training/internal/api"
	"training/internal/infra"
)

const (
	baseUrl = "localhost:8081"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	log.Printf("Connecting to DB: %s", dsn)
	db, err := infra.InitDB()
	if err != nil {
		log.Fatal("Failed to open db", err)
	}
	defer db.Close()

	userAPI := &api.UserAPI{DB: db}

	r := chi.NewRouter()

	api.RegisterUserRoutes(r, userAPI)

	log.Printf("Starting server on %s", baseUrl)
	err = http.ListenAndServe(baseUrl, r)
	if err != nil {
		log.Fatal(err)
	}
}
