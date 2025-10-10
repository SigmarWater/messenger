package repository

import (
	"context"

	"github.com/SigmarWater/messenger/auth/internal/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) (*model.UserService, error)
	InsertUser(ctx context.Context, user *model.UserService) (int, error)
}
