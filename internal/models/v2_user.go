package models

import "time"

type UserV2 struct {
	ID        int       `json:"id"`
	FullName  string    `json:"full_name"`
	Age       int       `json:"age"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
