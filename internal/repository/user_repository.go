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

	CreateUserV2(user *models.UserV2) error
	GetUserByIDV2(id int) (*models.UserV2, error)
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

func (r *userRepo) CreateUserV2(user *models.UserV2) error {
	query, args, err := psql.Insert("users_v2").
		Columns("full_name", "email", "age", "created_at").
		Values(user.FullName, user.Email, user.Age, user.CreatedAt).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(query, args...)
	return err
}

func (r *userRepo) GetUserByIDV2(id int) (*models.UserV2, error) {
	query, args, err := psql.Select("id, full_name, email, age, created_at").
		From("users_v2").
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, err
	}

	row := r.db.QueryRow(query, args...)
	var user models.UserV2
	err = row.Scan(&user.ID, &user.FullName, &user.Email, &user.Age, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
