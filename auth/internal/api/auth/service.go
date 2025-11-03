package auth

import (
	"github.com/SigmarWater/messenger/auth/internal/service"
	pb "github.com/SigmarWater/messenger/auth/pkg/api/auth_service"
)

type Implementation struct {
	pb.UnimplementedUserAPIServer
	userService service.UsersService
}

func NewImplementation(userService service.UsersService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}