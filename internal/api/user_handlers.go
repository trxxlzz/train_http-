package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type UserHandler interface {
	CreateUserHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
}

type UserAPI struct {
	DB *sql.DB
}

var _ UserHandler = (*UserAPI)(nil)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *UserAPI) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("createUserHandler started")

	info := User{}

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "failed to decode data", http.StatusBadRequest)
		return
	}

	if info.Name == "" || info.Email == "" || info.Age <= 0 {
		http.Error(w, "name, email and age are required and age must be positive", http.StatusBadRequest)
		return
	}

	user := &User{
		Name:      info.Name,
		Age:       info.Age,
		Email:     info.Email,
		CreatedAt: time.Now(),
	}

	queryBuilder := psql.Insert("users"). // Вот тут важно — `psql`, а не `squirrel`
						Columns("name", "age", "email", "created_at").
						Values(user.Name, user.Age, user.Email, user.CreatedAt).
						Suffix("RETURNING id")

	sql, args, err := queryBuilder.ToSql()
	if err != nil {
		http.Error(w, "failed to build query", http.StatusInternalServerError)
		return
	}

	log.Printf("SQL: %s, ARGS: %+v", sql, args)

	err = s.DB.QueryRow(sql, args...).Scan(&user.ID)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		http.Error(w, fmt.Sprintf("failed to create user: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to write user", http.StatusBadRequest)
	}
}

func (s *UserAPI) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	queryBuilder := psql.Select("*").
		From("users").
		Where(squirrel.Eq{"id": id})

	sql, args, err := queryBuilder.ToSql()
	if err != nil {
		http.Error(w, "failed to build query", http.StatusInternalServerError)
		return
	}

	row := s.DB.QueryRow(sql, args...)

	user := &User{}
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.CreatedAt)
	if err != nil {
		http.Error(w, "failed to fetch user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "failed to write user", http.StatusInternalServerError)
	}
}

func RegisterUserRoutes(r chi.Router, handler UserHandler) {
	r.Post("/users", handler.CreateUserHandler)
	r.Get("/users/{id}", handler.GetUserHandler)
}
