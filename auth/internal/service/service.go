package service

import (
	"context"
	"github.com/SigmarWater/messenger/auth/internal/model"
)

type UsersService interface {
	Get(ctx context.Context, id int64) (model.UserService, error)
	Create(ctx context.Context, user *model.UserService) (id int64, err error)
}
