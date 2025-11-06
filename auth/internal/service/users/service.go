package users

import (
	//pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/SigmarWater/messenger/auth/internal/repository"
)

type serv struct {
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) *serv {
	return &serv{
		userRepository: userRepository,
	}
}