package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Repo struct {
	DB *sql.DB
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (s *Repo) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("createUserHandler started")

	info := &User{}

	if err := json.NewDecoder(r.Body).Decode(&info); err != nil {
		http.Error(w, "failed to decode data", http.StatusBadRequest)
		return
	}

	if info.Name == "" || info.Email == "" {
		http.Error(w, "name or email are required", http.StatusBadRequest)
		return
	}

	user := &User{
		Name:      info.Name,
		Age:       info.Age,
		Email:     info.Email,
		CreatedAt: time.Now(),
	}

	sql := `INSERT INTO users (name, age, email, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := s.DB.QueryRow(sql, user.Name, user.Age, user.Email, user.CreatedAt).Scan(&user.ID)
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

func (s *Repo) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "id not found", http.StatusBadRequest)
		return
	}

	sql := `SELECT * FROM users WHERE id = $1`
	row := s.DB.QueryRow(sql, id)

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
