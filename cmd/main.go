package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"training/internal/repository"
)

const (
	baseUrl    = "localhost:8081"
	createUser = "/users"
	getUser    = "/users/{id}"
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
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to open db", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	repo := &repository.Repo{DB: db}

	r := chi.NewRouter()

	r.Post(createUser, repo.CreateUserHandler)
	r.Get(getUser, repo.GetUserHandler)

	log.Printf("Starting server on %s", baseUrl)
	err = http.ListenAndServe(baseUrl, r)
	if err != nil {
		log.Fatal(err)
	}
}
