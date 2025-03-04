package main

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"training/internal/api/v1"
	v2 "training/internal/api/v2"
	"training/internal/infra"
	"training/internal/repository"
	serviceV1 "training/internal/service/v1"
	serviceV2 "training/internal/service/v2"
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

	userServiceV1 := serviceV1.NewUserService(userRepo)
	userHandlerV1 := v1.NewUserHandler(userServiceV1)

	userServiceV2 := serviceV2.NewUserServiceV2(userRepo)
	userHandlerV2 := v2.NewUserHandlerV2(userServiceV2)

	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		v1.RegisterUserRoutesV1(r, userHandlerV1)
	})

	r.Route("/api/v2", func(r chi.Router) {
		v2.RegisterUserRoutesV2(r, userHandlerV2)
	})

	log.Println("Server running at localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
