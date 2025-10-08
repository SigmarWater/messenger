package users

import (
	"context"

	"github.com/SigmarWater/messenger/auth/intenal/model"
)

type UserRepository interface{
	GetUser(ctx context.Context, id int) (*model.UserService, error)
	InsertUser(ctx context.Context, user *model.UserService) (int, error) 
}