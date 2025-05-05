package repository

import (
	"context"
	"database/sql"
	"github.com/Alexx1088/authservice/internal/model"
)

type authRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.ExecContext(ctx,
		`INSERT INTO users (id, email, hashed_password) VALUES ($1, $2, $3)`,
		user.ID, user.Email, user.HashedPassword,
	)
	return err
}

func (r *authRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	row := r.db.QueryRowContext(ctx,
		`SELECT id, email, hashed_password FROM users WHERE email = $1`, email,
	)

	var u model.User
	if err := row.Scan(&u.ID, &u.Email, &u.HashedPassword); err != nil {
		return nil, err
	}
	return &u, nil
}
