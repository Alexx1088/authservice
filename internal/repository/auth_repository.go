package repository

import (
	"context"
	"github.com/Alexx1088/authservice/internal/model"
)

type AuthRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}
