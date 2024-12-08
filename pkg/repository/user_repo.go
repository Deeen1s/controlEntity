package repository

import (
	"context"
	models "controlEntity/pkg/model"
	"database/sql"
	"errors"
)

type UserRepositoryDB struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{db: db}
}

func (r *UserRepositoryDB) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	query := "INSERT INTO users (name, email, created_at, updated_at) VALUES ($1, $2, NOW(), NOW()) RETURNING id"
	var id int
	err := r.db.QueryRow(query, user.Name, user.Email).Scan(&id)
	return models.User{}, err
}

func (r *UserRepositoryDB) GetUserID(ctx context.Context, id int) (models.User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1"
	var user models.User
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err == sql.ErrNoRows {
		return models.User{}, nil
	}
	return user, err
}

func (r *UserRepositoryDB) UpdateUser(ctx context.Context, id int, user models.User) error {
	query := "UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3"
	_, err := r.db.Exec(query, user.Name, user.Email, id)
	return err
}

func (r *UserRepositoryDB) DeleteUser(ctx context.Context, id int) error {
	query := "DELETE FROM users WHERE id = $1"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("Пользователь не найден")
	}
	return nil
}

func (r *UserRepositoryDB) GetListUser(ctx context.Context) ([]models.User, error) {
	query := "SELECT id, name, email, created_at, updated_at FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
