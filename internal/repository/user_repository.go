package repository

import (
	"database/sql"
	"github.com/Masterminds/squirrel"
	"training/internal/models"
)

var psql = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByID(id int) (*models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *models.User) error {
	query, args, err := psql.Insert("users").
		Columns("name", "age", "email", "created_at").
		Values(user.Name, user.Age, user.Email, user.CreatedAt).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return err
	}

	return r.db.QueryRow(query, args...).Scan(&user.ID)
}

func (r *userRepo) GetUserByID(id int) (*models.User, error) {
	query, args, err := psql.Select("*").
		From("users").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(query, args...)
	user := &models.User{}
	err = row.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}
