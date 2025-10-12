package auth 

import(
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
	"github.com/SigmarWater/messenger/auth/internal/service"
)

type Implementation struct{
	pb.UnimplementedUserAPIServer
	userService service.UsersService
}

func NewImplementation(userService service.UsersService) *Implementation{
	return &Implementation{
		userService: userService,
	}
}