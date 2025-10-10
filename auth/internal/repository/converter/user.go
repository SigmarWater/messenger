package converter

import (
	"github.com/SigmarWater/messenger/auth/internal/model"
	modelRepo "github.com/SigmarWater/messenger/auth/internal/repository/users/model"
)

func ToUserFromRepo(user modelRepo.UserRepository) *model.UserService {
	return &model.UserService{
		Name:     user.Name,
		Email:    user.Email,
		Role:     user.Role,
		CreateAt: user.CreateAt.Time,
		UpdateAt: user.UpdateAt.Time,
	}
}
