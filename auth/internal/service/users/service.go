package users 

import(
	//pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/SigmarWater/messenger/auth/internal/repository"
	"github.com/SigmarWater/messenger/auth/internal/service"

)

type serv struct{
	userRepository repository.UserRepository
}

func NewService(userRepository repository.UserRepository) service.UsersService{
	return &serv{
		userRepository: userRepository,
	}
}