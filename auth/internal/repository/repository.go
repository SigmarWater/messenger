package repository

import (
	"context"
	"time"

	"github.com/SigmarWater/messenger/auth/internal/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int64) (*model.UserService, error)
	InsertUser(ctx context.Context, user *model.UserService) (int64, error)
}

type UserCache interface {
	Get(ctx context.Context, uuid string) (model.UserService, error)
	Set(ctx context.Context, uuid string, user model.UserService, ttl time.Duration) error
	Del(ctx context.Context, uuid string) error
}
