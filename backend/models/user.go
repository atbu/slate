package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	Username     string
	PasswordHash string
	CreatedAt    time.Time
	LastLogin    *time.Time
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(email, username, passwordHash string) (*User, error) {
	user := &User{
		ID:           uuid.New(),
		Email:        email,
		Username:     username,
		PasswordHash: passwordHash,
		CreatedAt:    time.Now(),
	}

	query := `
		INSERT INTO users (id, email, username, password_hash, created_at)
		VALUES ($1, $2, $3, $4, $5);
	`

	_, err := r.db.Exec(query, user.ID, user.Email, user.Username, user.PasswordHash, user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	query := `
		SELECT id, email, username, password_hash, created_at, last_login
		FROM users
		WHERE email = $1;
	`

	var user User
	var lastLogin sql.NullTime

	err := r.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
		&lastLogin,
	)
	if err != nil {
		return nil, err
	}

	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}

func (r *UserRepository) GetUserByID(id uuid.UUID) (*User, error) {
	query := `
		SELECT id, email, username, password_hash, created_at, last_login
		FROM users
		WHERE id = $1;
	`

	var user User
	var lastLogin sql.NullTime

	err := r.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.PasswordHash,
		&user.CreatedAt,
		&lastLogin,
	)
	if err != nil {
		return nil, err
	}

	if lastLogin.Valid {
		user.LastLogin = &lastLogin.Time
	}

	return &user, nil
}
