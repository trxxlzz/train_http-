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
	"training/internal/repository"
	"training/internal/service"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	dsn := os.Getenv("DSN")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := infra.InitDB()
	if err != nil {
		log.Fatal("Failed to connect to DB", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := api.NewUserHandler(userService)

	r := chi.NewRouter()
	api.RegisterUserRoutes(r, userHandler)

	log.Println("Server running at localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
