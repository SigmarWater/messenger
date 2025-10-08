package converter

import (
	"github.com/SigmarWater/messenger/auth/intenal/model"
	modelRepo "github.com/SigmarWater/messenger/auth/intenal/repository/users/model"
)

func ToUserFromRepo(user modelRepo.UserRepository) *model.UserService {
	return &model.UserService{
		Name: user.Name,
		Email: user.Email,
		EnterPassword: user.Password,
		Role: user.Role,
		CreateAt:  user.CreateAt.Time,
		UpdateAt: user.UpdateAt.Time,
	}
}
